import {useToast} from '@chakra-ui/react';
import React from 'react';
import {useGoogleReCaptcha} from 'react-google-recaptcha-v3';
import {CreateAccountRegisterEmailResponseSchema} from '../../utils/types/createAccount';
import {Margin} from '../Common/Margin';
import {useRequest} from '../Common/useRequest';
import {type EmailForm, EmailInputForm} from './EmailInputForm';
import {DefaultPageProps} from './RegisterAccount';

interface Props extends DefaultPageProps {
  setRegisterToken: (token: string) => void;
}

export const EmailInputPage: React.FC<Props> = props => {
  const toast = useToast();
  const {executeRecaptcha} = useGoogleReCaptcha();
  const {request} = useRequest('/v2/register/email/send', {
    errorCallback: () => {
      props.setStatus('error');
    },
  });

  const onSubmit = async (data: EmailForm) => {
    if (!executeRecaptcha) {
      toast({
        title: 'reCAPTCHAの読み込みに失敗しました',
        status: 'error',
      });
      return;
    }

    props.setStatus('loading');

    const form = new FormData();
    form.append('email', data.email);

    try {
      form.append('recaptcha', await executeRecaptcha());
    } catch {
      toast({
        title: 'reCAPTCHAの読み込みに失敗しました',
        status: 'error',
      });
      props.setStatus('error');
      return;
    }

    const res = await request({
      method: 'POST',
      credentials: 'include',
      mode: 'cors',
      body: form,
    });

    if (res) {
      const data = CreateAccountRegisterEmailResponseSchema.safeParse(
        await res.json()
      );
      if (data.success) {
        props.setRegisterToken(data.data.register_token);
        props.setStatus(undefined);
        props.nextStep();
        return;
      }
      console.error(data.error);
    }
    props.setStatus('error');
  };

  return (
    <Margin>
      <EmailInputForm onSubmit={onSubmit} />
    </Margin>
  );
};
