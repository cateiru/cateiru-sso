import {
  Button,
  Input,
  FormControl,
  FormLabel,
  FormErrorMessage,
  Box,
} from '@chakra-ui/react';
import React from 'react';
import {useForm} from 'react-hook-form';
import type {FieldValues} from 'react-hook-form';

const UserPassword: React.FC<{
  submit: (values: FieldValues) => void;
}> = ({submit}) => {
  const {
    handleSubmit,
    register,
    formState: {errors, isSubmitting},
  } = useForm();

  const submitHandler = (values: FieldValues) => {
    submit(values);

    return () => {};
  };

  return (
    <Box width={{base: '100%', sm: '90%', md: '600px'}}>
      <form onSubmit={handleSubmit(submitHandler)}>
        <FormControl isInvalid={errors.email}>
          <FormLabel htmlFor="email">メールアドレス</FormLabel>
          <Input
            id="email"
            type="email"
            placeholder="メールアドレス"
            {...register('email', {
              required: 'メールアドレスが必要です',
              pattern: {
                value:
                  /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/,
                message: 'メールアドレスの形式が違うようです',
              },
            })}
          />
          <FormErrorMessage>
            {errors.email && errors.email.message}
          </FormErrorMessage>
        </FormControl>
        <Button
          marginTop="1rem"
          colorScheme="blue"
          isLoading={isSubmitting}
          type="submit"
          width={{base: '100%', md: 'auto'}}
        >
          メールアドレスを認証する
        </Button>
      </form>
    </Box>
  );
};

export default UserPassword;
