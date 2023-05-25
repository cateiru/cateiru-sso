import {Box, Text, useColorModeValue, useToast} from '@chakra-ui/react';
import {useRequest} from '../../Common/useRequest';
import {
  UpdatePasswordForm,
  type UpdatePasswordFormData,
} from './UpdatePasswordForm';

export const UpdatePassword = () => {
  const textColor = useColorModeValue('gray.500', 'gray.400');
  const toast = useToast();

  const {request} = useRequest('/v2/account/password/update');

  const onSubmit = async (data: UpdatePasswordFormData) => {
    const form = new FormData();
    form.append('new_password', data.new_password);
    form.append('old_password', data.password);

    const res = await request({
      method: 'POST',
      body: form,
      mode: 'cors',
      credentials: 'include',
    });

    if (res) {
      toast({
        title: 'パスワードを更新しました',
        status: 'success',
      });
    }
  };

  return (
    <Box mt="2rem">
      <Text color={textColor} mb="1rem">
        パスワードを更新します。
      </Text>
      <UpdatePasswordForm onSubmit={onSubmit} />
    </Box>
  );
};
