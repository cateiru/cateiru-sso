import {
  Button,
  FormControl,
  FormErrorMessage,
  FormLabel,
  Input,
} from '@chakra-ui/react';
import React from 'react';
import {useForm} from 'react-hook-form';

export interface EmailForm {
  email: string;
}

interface Props {
  onSubmit: (data: EmailForm) => Promise<void>;
}

export const EmailInputForm: React.FC<Props> = ({onSubmit}) => {
  const {
    handleSubmit,
    register,
    formState: {errors, isSubmitting},
  } = useForm<EmailForm>();

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <FormControl isInvalid={!!errors.email}>
        <FormLabel htmlFor="email">メールアドレス</FormLabel>
        <Input
          id="email"
          type="email"
          {...register('email', {
            required: 'Emailは必須項目です',
            pattern: {
              value: /^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}$/i,
              message: '正しいEmailを入力してください',
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
        認証メールを送信する
      </Button>
    </form>
  );
};
