import {Box, useToast} from '@chakra-ui/react';
import React from 'react';
import {useGoogleReCaptcha} from 'react-google-recaptcha-v3';
import {api} from '../../utils/api';
import {CreateAccountRegisterEmailResponseSchema} from '../../utils/types/createAccount';
import {ErrorSchema, ErrorUniqueMessage} from '../../utils/types/error';
import {type EmailForm, EmailInputForm} from './EmailInputForm';
import {DefaultPageProps} from './RegisterAccount';

interface Props extends DefaultPageProps {
  setRegisterToken: (token: string) => void;
}

export const EmailInputPage: React.FC<Props> = props => {
  const toast = useToast();
  const {executeRecaptcha} = useGoogleReCaptcha();

  const onSubmit = async (data: EmailForm) => {
    if (!executeRecaptcha) {
      console.log('Execute recaptcha not yet available');
      return;
    }

    props.setStatus('loading');

    const form = new FormData();
    form.append('email', data.email);
    form.append('recaptcha', await executeRecaptcha());

    try {
      const res = await fetch(api('/v2/register/email/send'), {
        method: 'POST',
        credentials: 'include',
        mode: 'cors',
        body: form,
      });

      if (!res.ok) {
        const error = ErrorSchema.parse(await res.json());
        toast({
          title: ErrorUniqueMessage[error.unique_code] ?? error.message,
          status: 'error',
          duration: 5000,
          isClosable: true,
        });
        return;
      }

      const data = CreateAccountRegisterEmailResponseSchema.parse(
        await res.json()
      );

      props.setRegisterToken(data.register_token);
      props.setStatus(undefined);
      props.nextStep();
    } catch (e) {
      if (e instanceof Error) {
        toast({
          title: 'エラー',
          description: e.message,
          status: 'error',
          duration: 5000,
          isClosable: true,
        });
      }
      props.setStatus('error');
    }
  };

  return (
    <Box w={{base: '95%', md: '600px'}} m="auto">
      <EmailInputForm onSubmit={onSubmit} />
    </Box>
  );
};
