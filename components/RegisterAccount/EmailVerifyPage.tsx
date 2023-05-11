import {Box, Text, VStack, useToast} from '@chakra-ui/react';
import React from 'react';
import {RegisterVerifyEmailResponseSchema} from '../../utils/types/createAccount';
import {Margin} from '../Common/Margin';
import {useRequest} from '../Common/useRequest';
import {EmailResend} from './EmailResend';
import {EmailVerifyForm} from './EmailVerifyForm';
import {DefaultPageProps} from './RegisterAccount';

interface Props extends DefaultPageProps {
  registerToken: string;
}

export const EmailVerifyPage: React.FC<Props> = props => {
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
        props.setStatus('error');
        throw new Error();
      }

      props.setStatus(undefined);
      props.nextStep();
    }
    throw new Error();
  };

  return (
    <Margin>
      <VStack>
        <Text mb="1rem" textAlign="center">
          メールアドレスに送信された6桁のコードを入力してください
        </Text>
        <Box>
          <EmailVerifyForm onSubmit={onSubmit} />
        </Box>
        <EmailResend registerToken={props.registerToken} />
      </VStack>
    </Margin>
  );
};
