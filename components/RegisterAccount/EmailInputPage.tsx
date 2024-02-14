import React from 'react';
import {CreateAccountRegisterEmailResponseSchema} from '../../utils/types/createAccount';
import {Margin} from '../Common/Margin';
import {useRecaptcha} from '../Common/useRecaptcha';
import {useRequest} from '../Common/useRequest';
import {type EmailForm, EmailInputForm} from './EmailInputForm';
import {DefaultPageProps} from './RegisterAccount';

interface Props extends DefaultPageProps {
  setRegisterToken: (token: string) => void;
  setEmail: (email: string) => void;
}

export const EmailInputPage: React.FC<Props> = props => {
  const {getRecaptchaToken} = useRecaptcha();
  const {request} = useRequest('/register/email/send', {
    errorCallback: () => {
      props.setStatus('error');
    },
  });

  const onSubmit = async (data: EmailForm) => {
    props.setStatus('loading');

    const email = data.email;

    const form = new FormData();
    form.append('email', email);

    const recaptchaToken = await getRecaptchaToken();
    if (typeof recaptchaToken === 'undefined') {
      return;
    }
    form.append('recaptcha', recaptchaToken);

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
        props.setEmail(email);
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
