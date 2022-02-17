import {
  Box,
  FormControl,
  FormLabel,
  Input,
  FormErrorMessage,
  Button,
  Center,
  Heading,
  InputGroup,
  InputRightElement,
  IconButton,
  useToast,
  Link,
  Divider,
} from '@chakra-ui/react';
import NextLink from 'next/link';
import {useRouter} from 'next/router';
import React from 'react';
import {useGoogleReCaptcha} from 'react-google-recaptcha-v3';
import {useForm} from 'react-hook-form';
import type {FieldValues} from 'react-hook-form';
import {IoEyeOutline, IoEyeOffOutline} from 'react-icons/io5';
import {useSetRecoilState, useResetRecoilState} from 'recoil';
import {login} from '../../utils/api/login';
import {LoadState, UserState} from '../../utils/state/atom';

const LoginForm = () => {
  const {
    handleSubmit,
    register,
    formState: {errors, isSubmitting},
  } = useForm();
  const [recaptcha, setRecaptcha] = React.useState('');
  const [show, setShow] = React.useState(false);
  const [load, setLoad] = React.useState(false);
  const [redirect, setRedirect] = React.useState('');
  const router = useRouter();
  const toast = useToast();
  const resetUser = useResetRecoilState(UserState);
  const setRLoad = useSetRecoilState(LoadState);

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

  React.useEffect(() => {
    if (!router.isReady) return;
    const query = router.query;

    if (typeof query['redirect'] === 'string') {
      setRedirect(query['redirect']);
    }
  }, [router.isReady, router.query]);

  const submit = (values: FieldValues) => {
    if (recaptcha.length === 0) {
      toast({
        title: 'reCAPTCHAのトークンを取得できませんでした',
        status: 'error',
        isClosable: true,
        duration: 9000,
      });
      router.reload();
      return;
    }

    const f = async () => {
      setLoad(true);
      let token: string | undefined;
      try {
        token = await login(values.email, values.password, recaptcha);
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

      if (token) {
        if (redirect !== '') {
          router.push(
            `/login/otp?t=${token}&redirect=${encodeURIComponent(redirect)}`
          );
        } else {
          router.push(`/login/otp?t=${token}`);
        }
      } else {
        // me情報を取得するためにuserを初期化する
        resetUser();

        // redirectが定義されている場合はそれに飛ぶ
        if (redirect !== '') {
          router.push(redirect);
        } else {
          setRLoad(true);
          router.push('/hello');
        }
      }
    };

    f();

    return () => {};
  };

  return (
    <Box width={{base: '100%', sm: '90%', md: '600px'}}>
      <Center mb="2rem">
        <Heading>ログイン</Heading>
      </Center>
      <form onSubmit={handleSubmit(submit)}>
        <FormControl isInvalid={errors.email}>
          <FormLabel htmlFor="email">メールアドレス</FormLabel>
          <Input
            id="email"
            type="email"
            placeholder="メールアドレス"
            {...register('email', {
              required: 'メールアドレスの入力が必要です',
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
        <FormControl isInvalid={errors.password} mt="1rem">
          <FormLabel htmlFor="password">パスワード</FormLabel>
          <InputGroup>
            <Input
              id="password"
              type={show ? 'text' : 'password'}
              placeholder="パスワード"
              {...register('password', {
                required: 'パスワードの入力が必要です',
              })}
            />
            <InputRightElement>
              <IconButton
                variant="ghost"
                aria-label="show password"
                icon={
                  show ? (
                    <IoEyeOutline size="25px" />
                  ) : (
                    <IoEyeOffOutline size="25px" />
                  )
                }
                size="sm"
                onClick={() => setShow(!show)}
              />
            </InputRightElement>
          </InputGroup>
          <FormErrorMessage>
            {errors.password && errors.password.message}
          </FormErrorMessage>
        </FormControl>
        <Box mt="1rem">
          <NextLink href="/forget" passHref>
            <Link>パスワードを忘れましたか？</Link>
          </NextLink>
        </Box>
        <Button
          marginTop="1rem"
          colorScheme="blue"
          isLoading={isSubmitting || load}
          type="submit"
          width={{base: '100%', md: 'auto'}}
        >
          ログインする
        </Button>
      </form>
      <Divider my="1rem" />
      <Center>
        <NextLink href="/create" passHref>
          <Link>アカウントを作成する</Link>
        </NextLink>
      </Center>
    </Box>
  );
};

export default LoginForm;
