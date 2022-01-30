import {
  Box,
  Input,
  FormControl,
  FormErrorMessage,
  Button,
  Center,
  Heading,
  useToast,
  Text,
} from '@chakra-ui/react';
import {useRouter} from 'next/router';
import React from 'react';
import {useForm} from 'react-hook-form';
import type {FieldValues} from 'react-hook-form';
import {useRecoilValue, useRecoilState} from 'recoil';
import {changeMail, changeMailVerify} from '../../utils/api/change';
import {UserState} from '../../utils/state/atom';

const MailSetting = () => {
  const [token, setToken] = React.useState('');
  const router = useRouter();
  const [user, setUser] = useRecoilState(UserState);
  const toast = useToast();

  React.useEffect(() => {
    if (!router.isReady) return;
    const query = router.query;

    if (typeof query['t'] === 'string') {
      setToken(query['t']);
    }
  }, [router.isReady, router.query]);

  React.useEffect(() => {
    const f = async () => {
      try {
        const newMail = await changeMailVerify(token);

        setUser(v => {
          if (v) {
            return {...v, mail: newMail};
          } else {
            return undefined;
          }
        });

        toast({
          title: 'メールアドレスを変更しました',
          status: 'info',
          isClosable: true,
          duration: 9000,
        });
      } catch (error) {
        if (error instanceof Error) {
          toast({
            title: error.message,
            status: 'error',
            isClosable: true,
            duration: 9000,
          });
        }
      }

      router.push('/setting/mail');
    };

    if (token) {
      f();
    }
  }, [token]);

  return (
    <>
      <Center height={{base: 'auto', md: '50vh'}} mx=".5rem">
        <Box
          width={{base: '100%', sm: '90%', md: '600px'}}
          mt={{base: '3rem', md: '0'}}
        >
          <Heading fontSize="1.5rem" mb="1.5rem" textAlign="center">
            メールアドレスを変更する
          </Heading>
          <Text>{user?.mail}</Text>
          <MailInput />
        </Box>
      </Center>
    </>
  );
};

const MailInput = () => {
  const {
    handleSubmit,
    register,
    formState: {errors, isSubmitting},
  } = useForm();
  const user = useRecoilValue(UserState);
  const toast = useToast();
  const [isSubmit, setSubmit] = React.useState(false);

  const submitHandler = (values: FieldValues) => {
    const f = async () => {
      try {
        await changeMail(values.email);
        toast({
          title: '新しいメールアドレスに確認メールを送信しました',
          status: 'info',
          isClosable: true,
          duration: 9000,
        });
        setSubmit(true);
      } catch (error) {
        if (error instanceof Error) {
          toast({
            title: error.message,
            status: 'error',
            isClosable: true,
            duration: 9000,
          });
        }
      }
    };

    if (user?.mail !== values.email) {
      f();
    }

    return () => {};
  };

  return (
    <form onSubmit={handleSubmit(submitHandler)}>
      <FormControl isInvalid={errors.email}>
        <Input
          id="email"
          type="email"
          placeholder="新しいメールアドレス"
          disabled={isSubmit}
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
  );
};

export default MailSetting;
