import {
  Button,
  FormControl,
  FormErrorMessage,
  FormLabel,
  Input,
} from '@chakra-ui/react';
import React from 'react';
import {useForm} from 'react-hook-form';
import {emailRegex} from '../../utils/regex';

export interface ForgetFormData {
  email: string;
}

interface Props {
  onSubmit: (data: ForgetFormData) => Promise<void>;
}

export const ForgetForm: React.FC<Props> = ({onSubmit}) => {
  const {
    handleSubmit,
    register,
    formState: {errors, isSubmitting},
  } = useForm<ForgetFormData>();

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <FormControl isInvalid={!!errors.email}>
        <FormLabel htmlFor="email">メールアドレス</FormLabel>
        <Input
          id="email"
          type="email"
          autoComplete="email"
          {...register('email', {
            required: 'メールアドレスは必須です',
            pattern: {
              value: emailRegex,
              message: '正しいメールアドレスを入力してください',
            },
          })}
        />
        <FormErrorMessage>
          {errors.email && errors.email.message}
        </FormErrorMessage>
      </FormControl>
      <Button
        mt="1rem"
        isLoading={isSubmitting}
        colorScheme="cateiru"
        type="submit"
        w="100%"
      >
        パスワード再登録メールを送信する
      </Button>
    </form>
  );
};
