import {Button, useToast} from '@chakra-ui/react';
import React from 'react';
import {useGoogleReCaptcha} from 'react-google-recaptcha-v3';
import {useRequest} from '../Common/useRequest';

interface Props {
  registerToken: string;
}

export const EmailResend: React.FC<Props> = props => {
  const toast = useToast();
  const {executeRecaptcha} = useGoogleReCaptcha();
  const [isDisabled, setIsDisabled] = React.useState(false);
  const {request} = useRequest('/v2/register/email/resend', {
    errorCallback: () => {},
  });

  const onClickHandler = async () => {
    if (!executeRecaptcha) {
      toast({
        title: 'reCAPTCHAの読み込みに失敗しました',
        status: 'error',
      });
      return;
    }

    const form = new FormData();
    try {
      form.append('recaptcha', await executeRecaptcha());
    } catch {
      toast({
        title: 'reCAPTCHAの読み込みに失敗しました',
        status: 'error',
      });
      return;
    }

    const res = await request({
      method: 'POST',
      credentials: 'include',
      mode: 'cors',
      body: form,
      headers: {
        'X-Register-Token': props.registerToken,
      },
    });

    if (res) {
      toast({
        title: 'メールを再送しました',
        description: 'メールを確認してください',
        status: 'success',
      });
      setIsDisabled(true);
      setTimeout(() => {
        setIsDisabled(false);
      }, 5000);
      return;
    }
  };

  return (
    <Button onClick={onClickHandler} variant="link" isDisabled={isDisabled}>
      メールを再送する
    </Button>
  );
};
