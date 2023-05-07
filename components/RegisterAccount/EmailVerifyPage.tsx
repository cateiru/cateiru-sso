import {Box, useToast} from '@chakra-ui/react';
import React from 'react';
import {RegisterVerifyEmailResponseSchema} from '../../utils/types/createAccount';
import {useRequest} from '../Common/useRequest';
import {EmailVerifyForm} from './EmailVerifyForm';
import {DefaultPageProps} from './RegisterAccount';

interface Props extends DefaultPageProps {
  registerToken: string;
}

export const EmailInputPage: React.FC<Props> = props => {
  const toast = useToast();
  const {request} = useRequest('/v2/register/email/verify', {
    errorCallback: () => {
      props.setStatus('error');
    },
  });

  const onSubmit = async (data: EmailVerifyForm) => {
    props.setStatus('loading');

    const form = new FormData();
    form.append('code', data.code);

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
      const result = RegisterVerifyEmailResponseSchema.safeParse(
        await res.json()
      );

      if (!result.success) {
        return;
      }

      if (!result.data.verified) {
        // 試行上限を上回った場合は、最初からやり直す
        if (result.data.remaining_count <= 0) {
          toast({
            title: 'もう一度最初からお試しください',
            description: '試行上限を上回りました',
            status: 'error',
          });
          props.reset();
          return;
        }

        toast({
          title: 'メールアドレスの確認に失敗しました',
          description: `あと、${result.data.remaining_count}回試行することができます`,
          status: 'error',
        });
        throw new Error();
      }

      props.setStatus(undefined);
      props.nextStep();
    }
  };

  return (
    <Box w={{base: '95%', md: '600px'}} m="auto">
      <EmailVerifyForm onSubmit={onSubmit} />
    </Box>
  );
};
