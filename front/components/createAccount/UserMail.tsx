import {
  Button,
  Input,
  FormControl,
  FormLabel,
  FormErrorMessage,
  Box,
} from '@chakra-ui/react';
import React from 'react';
import {useGoogleReCaptcha} from 'react-google-recaptcha-v3';
import {useForm} from 'react-hook-form';
import type {FieldValues} from 'react-hook-form';

const UserPassword: React.FC<{
  submit: (values: FieldValues, recaptcha: string) => void;
}> = ({submit}) => {
  const {
    handleSubmit,
    register,
    formState: {errors, isSubmitting},
  } = useForm();

  const [recaptcha, setRecaptcha] = React.useState('');

  const {executeRecaptcha} = useGoogleReCaptcha();
  const handleReCaptchaVerify = React.useCallback(async () => {
    if (!executeRecaptcha) {
      return;
    }
    const token = await executeRecaptcha();

    setRecaptcha(token);
  }, [executeRecaptcha, setRecaptcha]);

  // reCAPTCHAのトークンを取得する
  React.useEffect(() => {
    let unmounted = false;
    if (!unmounted) {
      handleReCaptchaVerify();
    }
    return () => {
      unmounted = true;
    };
  }, [executeRecaptcha]);

  const submitHandler = (values: FieldValues) => {
    submit(values, recaptcha);

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
        >
          メールアドレスを認証する
        </Button>
      </form>
    </Box>
  );
};

export default UserPassword;
