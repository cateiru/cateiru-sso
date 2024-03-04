import {Box, Text, useToast} from '@chakra-ui/react';
import {useSecondaryColor} from '../../Common/useColor';
import {useRequest} from '../../Common/useRequest';
import {
  UpdatePasswordForm,
  type UpdatePasswordFormData,
} from './UpdatePasswordForm';

export const UpdatePassword = () => {
  const textColor = useSecondaryColor();
  const toast = useToast();

  const {request} = useRequest('/account/password/update');

  const onSubmit = async (data: UpdatePasswordFormData) => {
    const form = new FormData();
    form.append('new_password', data.new_password);
    form.append('old_password', data.password);

    const res = await request({
      method: 'PUT',
      body: form,
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
