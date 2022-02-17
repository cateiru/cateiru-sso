import {
  Box,
  FormControl,
  Input,
  FormErrorMessage,
  Button,
  Center,
  Heading,
  useToast,
} from '@chakra-ui/react';
import {useRouter} from 'next/router';
import React from 'react';
import {useForm} from 'react-hook-form';
import type {FieldValues} from 'react-hook-form';
import {useSetRecoilState, useResetRecoilState} from 'recoil';
import {loginOTP} from '../../utils/api/login';
import {LoadState, UserState} from '../../utils/state/atom';

const OTPLoginForm: React.FC<{token: string; redirect: string}> = ({
  token,
  redirect,
}) => {
  const {
    handleSubmit,
    register,
    formState: {errors, isSubmitting},
  } = useForm();
  const [load, setLoad] = React.useState(false);
  const setRLoad = useSetRecoilState(LoadState);
  const resetUser = useResetRecoilState(UserState);
  const router = useRouter();
  const toast = useToast();

  const submit = (values: FieldValues) => {
    validate(values.otp);

    return () => {};
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    // OTPは数字6文字であるためtrueの場合は強制的に送信してしまう
    // backup codeは手動
    const passcode = e.target.value;
    if (/[0-9]{6}/g.test(passcode)) {
      validate(passcode);
    }
  };

  const validate = async (passcode: string) => {
    setLoad(true);

    try {
      await loginOTP(passcode, token);
    } catch (error) {
      if (error instanceof Error) {
        toast({
          title: error.message,
          status: 'error',
          isClosable: true,
          duration: 9000,
        });
      }
      setLoad(false);
      return;
    }

    // me情報を取得するためにuserを初期化する
    resetUser();

    if (redirect !== '') {
      router.push(redirect);
    } else {
      setRLoad(true);
      router.push('/hello');
    }
  };

  return (
    <Box width={{base: '100%', sm: '90%', md: '600px'}}>
      <Center mb="2.3rem">
        <Heading>ワンタイムパスワードを入力</Heading>
      </Center>
      <Center>
        <Box width={{base: '100%', sm: '400px', md: '460px'}}>
          <form onSubmit={handleSubmit(submit)}>
            <FormControl isInvalid={errors.otp}>
              <Input
                id="otp"
                type="text"
                autoComplete="one-time-code"
                placeholder="ワンタイムパスワード"
                {...register('otp', {
                  onChange: handleChange,
                  required: 'ワンタイムパスワードを入力する必要があります',
                })}
              />
              <FormErrorMessage>
                {errors.otp && errors.otp.otp}
              </FormErrorMessage>
            </FormControl>
            <Button
              marginTop="1rem"
              colorScheme="blue"
              isLoading={isSubmitting || load}
              type="submit"
              width={{base: '100%', md: 'auto'}}
            >
              検証してログインする
            </Button>
          </form>
        </Box>
      </Center>
    </Box>
  );
};

export default OTPLoginForm;
