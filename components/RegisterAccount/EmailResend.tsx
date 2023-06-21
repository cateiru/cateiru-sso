import {Button, useToast} from '@chakra-ui/react';
import React from 'react';
import {useRecaptcha} from '../Common/useRecaptcha';
import {useRequest} from '../Common/useRequest';

interface Props {
  registerToken: string;
}

export const EmailResend: React.FC<Props> = props => {
  const toast = useToast();
  const {getRecaptchaToken} = useRecaptcha();
  const [isDisabled, setIsDisabled] = React.useState(false);
  const {request} = useRequest('/v2/register/email/resend', {
    errorCallback: () => {},
  });

  const form = new FormData();

  const onClickHandler = async () => {
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
