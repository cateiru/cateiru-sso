import {
  Button,
  FormControl,
  FormErrorMessage,
  FormLabel,
  Input,
} from '@chakra-ui/react';
import React from 'react';
import {useForm} from 'react-hook-form';

export interface OtpRegisterFormData {
  code: string;
}

interface Props {
  onSubmit: (data: OtpRegisterFormData) => Promise<void>;
}

export const OtpRegisterForm: React.FC<Props> = props => {
  const {
    handleSubmit,
    register,
    formState: {errors, isSubmitting},
  } = useForm<OtpRegisterFormData>();

  return (
    <form onSubmit={handleSubmit(props.onSubmit)}>
      <FormControl isInvalid={!!errors.code}>
        <FormLabel htmlFor="code">アプリで生成したコード</FormLabel>
        <Input
          id="code"
          type="text"
          autoComplete="one-time-code"
          {...register('code', {
            required: 'アプリで生成したコードを入力してください',
            pattern: {
              value: /^[0-9]{6}$/,
              message: '6桁の数字を入力してください',
            },
          })}
        />
        <FormErrorMessage>
          {errors.code && errors.code.message}
        </FormErrorMessage>
      </FormControl>
      <Button
        mt="1rem"
        isLoading={isSubmitting}
        colorScheme="cateiru"
        type="submit"
        w="100%"
      >
        二段階認証を有効化します
      </Button>
    </form>
  );
};
