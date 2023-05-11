import {Button, FormControl, FormErrorMessage, Input} from '@chakra-ui/react';
import React from 'react';
import {useForm} from 'react-hook-form';

const emailRegex = /[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}/;
const userIDRegex = /[A-Z0-9_]{3,15}/;
const userIdEmailRegex = new RegExp(
  `^(${userIDRegex.source})|(${emailRegex.source})$`,
  'i'
);

export interface UserIDEmailForm {
  user_id_email: string;
}

interface Props {
  onSubmit: (data: UserIDEmailForm) => Promise<void>;
  onClickWebAuthn: () => Promise<void>;
  isConditionSupported: boolean;
}

export const UserIDEmailForm: React.FC<Props> = ({
  onSubmit,
  isConditionSupported,
  onClickWebAuthn,
}) => {
  const {
    handleSubmit,
    register,
    formState: {errors, isSubmitting},
  } = useForm<UserIDEmailForm>();

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <FormControl isInvalid={!!errors.user_id_email}>
        <Input
          id="user_id_email"
          type="email text"
          autoComplete="username webauthn"
          {...register('user_id_email', {
            required: 'この値は必須です',
            pattern: {
              value: userIdEmailRegex,
              message: '正しい形式で入力してください',
            },
          })}
        />
        <FormErrorMessage>
          {errors.user_id_email && errors.user_id_email.message}
        </FormErrorMessage>
      </FormControl>
      <Button
        mt="1rem"
        isLoading={isSubmitting}
        colorScheme="cateiru"
        type="submit"
        w="100%"
      >
        ログイン
      </Button>
      {isConditionSupported || (
        <Button w="100%" mt="1rem" onClick={onClickWebAuthn}>
          生体認証でログイン
        </Button>
      )}
    </form>
  );
};
