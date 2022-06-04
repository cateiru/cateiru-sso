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
  Modal,
  ModalOverlay,
  ModalContent,
  ModalHeader,
  ModalCloseButton,
  ModalBody,
  ModalFooter,
  useDisclosure,
  Text,
} from '@chakra-ui/react';
import NextLink from 'next/link';
import {useRouter} from 'next/router';
import React from 'react';
import {useGoogleReCaptcha} from 'react-google-recaptcha-v3';
import {useForm} from 'react-hook-form';
import type {FieldValues} from 'react-hook-form';
import {TbEye, TbEyeOff} from 'react-icons/tb';
import {TbExternalLink} from 'react-icons/tb';
import {useSetRecoilState, useResetRecoilState} from 'recoil';
import {login} from '../../utils/api/login';
import cookieValue from '../../utils/cookie';
import {UserState, NoLoginState} from '../../utils/state/atom';

const LoginForm = () => {
  const {
    handleSubmit,
    register,
    formState: {errors, isSubmitting},
  } = useForm();
  const [show, setShow] = React.useState(false);
  const [load, setLoad] = React.useState(false);
  const [redirect, setRedirect] = React.useState('');
  const router = useRouter();
  const toast = useToast();
  const resetUser = useResetRecoilState(UserState);
  const setNoLogin = useSetRecoilState(NoLoginState);
  const {isOpen, onOpen, onClose} = useDisclosure();

  const {executeRecaptcha} = useGoogleReCaptcha();

  React.useEffect(() => {
    if (!router.isReady) return;
    const query = router.query;

    if (typeof query['redirect'] === 'string') {
      setRedirect(query['redirect']);
    }
  }, [router.isReady, router.query]);

  const submit = (values: FieldValues) => {
    if (!executeRecaptcha) {
      return;
    }

    const f = async () => {
      setLoad(true);

      const recaptcha = await executeRecaptcha();

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
        setNoLogin(true);

        // redirectが定義されている場合はそれに飛ぶ
        if (redirect !== '') {
          router.push(redirect);
        } else {
          router.push('/hello');
        }
      }
    };

    f();

    return () => {};
  };

  const loginHandler = () => {
    const refresh = cookieValue('refresh-token');

    if (refresh) {
      onClose();
      resetUser();
      setNoLogin(true);

      // このhandlerはredirectが存在しているときに使用するものである
      router.push(redirect);
    } else {
      toast({
        title: 'ログインができませんでした',
        description:
          'もう一度試すか、モーダルを閉じてパスワードからログインしてください',
        status: 'warning',
        isClosable: true,
        duration: 9000,
      });
    }
  };

  return (
    <>
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
                  icon={show ? <TbEye size="25px" /> : <TbEyeOff size="25px" />}
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
          {redirect ? (
            <Link isExternal onClick={onOpen}>
              アカウントを作成する
            </Link>
          ) : (
            <NextLink href="/create" passHref>
              <Link>アカウントを作成する</Link>
            </NextLink>
          )}
        </Center>
      </Box>
      <Modal isOpen={isOpen} onClose={onClose} isCentered>
        <ModalOverlay />
        <ModalContent>
          <ModalHeader>アカウントを作成してください</ModalHeader>
          <ModalCloseButton size="lg" />
          <ModalBody>
            <Text>
              <Link href="/create" isExternal fontWeight="bold">
                アカウント作成ページ
              </Link>
              でアカウントを作成してください。
            </Text>
            <Text>
              このブラウザで作成が完了した場合は「ログインしました」を押してください。
            </Text>
            <Text color="red.500" mt=".5rem">
              * ブラウザのリロードをするとログイン出来ない場合があります
            </Text>
          </ModalBody>

          <ModalFooter>
            <Center width="100%">
              <Button
                colorScheme="blue"
                mr=".5rem"
                as={Link}
                href="/create"
                isExternal
                variant="solid"
                rightIcon={<TbExternalLink size="20px" />}
              >
                アカウントを作成
              </Button>
              <Button colorScheme="green" mr=".5rem" onClick={loginHandler}>
                ログインしました
              </Button>
            </Center>
          </ModalFooter>
        </ModalContent>
      </Modal>
    </>
  );
};

export default LoginForm;
