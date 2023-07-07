import {
  Button,
  Flex,
  FormControl,
  FormErrorMessage,
  Input,
  Select,
} from '@chakra-ui/react';
import React from 'react';
import {useForm} from 'react-hook-form';
import {useRequest} from '../../Common/useRequest';

interface Props {
  orgId: string;

  handleSuccess: () => void;
}

interface JoinFormData {
  user_name_or_email: string;
  role: string;
}

export const OrgJoinUser: React.FC<Props> = props => {
  const {
    handleSubmit,
    register,
    reset,
    formState: {isSubmitting, errors},
  } = useForm<JoinFormData>({
    defaultValues: {
      role: 'guest',
    },
  });
  const {request} = useRequest('/v2/admin/org/member');

  const onSubmit = async (data: JoinFormData) => {
    const form = new FormData();

    form.append('org_id', props.orgId);
    form.append('user_name_or_email', data.user_name_or_email);
    form.append('role', data.role);

    const res = await request({
      method: 'POST',
      body: form,
      mode: 'cors',
      credentials: 'include',
    });

    if (res) {
      reset();
      props.handleSuccess();
    }
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <Flex>
        <FormControl isInvalid={!!errors.user_name_or_email}>
          <Input
            type="text"
            placeholder="ユーザー名またはメールアドレス"
            {...register('user_name_or_email', {
              required: 'ユーザーは必須です',
            })}
          />
          <FormErrorMessage>
            {errors.user_name_or_email && errors.user_name_or_email.message}
          </FormErrorMessage>
        </FormControl>
        <FormControl isInvalid={!!errors.role} ml=".5rem">
          <Select
            {...register('role', {
              required: 'ロールは必須です',
            })}
          >
            <option value="owner">管理者</option>
            <option value="member">メンバー</option>
            <option value="guest">ゲスト</option>
          </Select>

          <FormErrorMessage>
            {errors.role && errors.role.message}
          </FormErrorMessage>
        </FormControl>
      </Flex>

      <Button
        mt=".5rem"
        w="100%"
        colorScheme="cateiru"
        type="submit"
        isLoading={isSubmitting}
      >
        ユーザーを追加
      </Button>
    </form>
  );
};
