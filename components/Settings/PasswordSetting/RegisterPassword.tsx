import {Box, Text, useToast} from '@chakra-ui/react';
import {useSWRConfig} from 'swr';
import {RegisterPasswordForm} from '../../Common/Form/RegisterPasswordForm';
import {RegisterPasswordFormContextData} from '../../Common/Form/RegisterPasswordFormContext';
import {useSecondaryColor} from '../../Common/useColor';
import {useRequest} from '../../Common/useRequest';

export const RegisterPassword = () => {
  const textColor = useSecondaryColor();
  const toast = useToast();
  const {mutate} = useSWRConfig();

  const {request} = useRequest('/account/password');

  const onSubmitPassword = async (data: RegisterPasswordFormContextData) => {
    const form = new FormData();
    form.append('new_password', data.new_password);

    const res = await request({
      method: 'POST',
      body: form,
    });

    if (res) {
      toast({
        title: 'パスワードを追加しました',
        status: 'success',
      });

      // キャッシュ飛ばす
      mutate(
        key =>
          typeof key === 'string' && key.startsWith('/account/certificates'),
        undefined,
        {revalidate: true}
      );
    }
  };

  return (
    <Box mt="2rem">
      <Text color={textColor}>
        現在、このアカウントにはパスワードは設定されていません。
      </Text>
      <Text mb="1rem" color={textColor}>
        新規にパスワードを設定する場合は、以下フォームから追加してください。
      </Text>
      <RegisterPasswordForm
        onSubmit={onSubmitPassword}
        buttonText="パスワードを設定"
      />
    </Box>
  );
};
