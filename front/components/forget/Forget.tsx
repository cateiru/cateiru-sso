import {
  Box,
  Heading,
  Flex,
  FormControl,
  Button,
  Input,
  FormErrorMessage,
  useToast,
  Text,
} from '@chakra-ui/react';
import {useRouter} from 'next/router';
import React from 'react';
import {useForm} from 'react-hook-form';
import type {FieldValues} from 'react-hook-form';
import {TbMailForward} from 'react-icons/tb';
import {sendForget} from '../../utils/api/forget';

const Forget = () => {
  const {
    handleSubmit,
    register,
    formState: {errors, isSubmitting},
  } = useForm();
  const toast = useToast();
  const router = useRouter();
  const [mail, setMail] = React.useState('');

  const submit = (values: FieldValues) => {
    const f = async () => {
      try {
        await sendForget(values.email);

        setMail(values.email);
      } catch (error) {
        if (error instanceof Error) {
          toast({
            title: 'このメールを送信できませんでした',
            description: error.message,
            status: 'error',
            isClosable: true,
            duration: 9000,
          });
        }
      }
    };

    f();

    return () => {};
  };

  const SetMail = () => {
    return (
      <>
        <Heading fontSize={{base: '1.5rem', md: '2rem'}}>
          パスワードを再設定します
        </Heading>
        <Box width={{base: '100%', lg: '800px'}} mt="2rem">
          <form onSubmit={handleSubmit(submit)}>
            <FormControl isInvalid={errors.email}>
              <Input
                id="email"
                type="email"
                placeholder="登録したメールアドレス"
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
              rightIcon={<TbMailForward size="20px" />}
            >
              再登録メールを送信する
            </Button>
          </form>
        </Box>
      </>
    );
  };

  const AlreadySendMail = () => {
    return (
      <>
        <Heading fontSize={{base: '1.5rem', md: '2rem'}} mb="1rem">
          メールを送信しました。
        </Heading>
        <Text textAlign="center">
          <Text as="span" fontWeight="bold">
            {mail}
          </Text>
          にパスワード再登録メールを送信しました。
        </Text>
      </>
    );
  };

  return (
    <Flex
      justifyContent="center"
      alignItems="center"
      flexDirection="column"
      width="100%"
      height="80vh"
      px={{base: '1rem', md: '5rem'}}
    >
      {mail ? <AlreadySendMail /> : <SetMail />}
    </Flex>
  );
};

export default Forget;
