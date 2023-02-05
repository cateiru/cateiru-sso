import {Flex, Heading, Box, Button, useToast} from '@chakra-ui/react';
import {useRouter} from 'next/router';
import React from 'react';
import {useForm, FormProvider, SubmitHandler} from 'react-hook-form';
import {acceptPassword} from '../../utils/api/forget';
import Password, {PasswordForm} from '../common/form/Password';

const AcceptPage: React.FC<{token: string}> = ({token}) => {
  const methods = useForm<PasswordForm>();
  const {
    handleSubmit,
    setError,
    clearErrors,
    formState: {isSubmitting},
  } = methods;
  const toast = useToast();
  const [pwOk, setPWOK] = React.useState(false);
  const router = useRouter();

  const submit: SubmitHandler<PasswordForm> = values => {
    if (!pwOk) {
      setError('password', {
        type: 'custom',
        message: 'custom message',
      });
      return;
    } else {
      clearErrors('password');
    }

    const f = async () => {
      try {
        await acceptPassword(token, values.password);
        toast({
          title: 'パスワードを変更しました',
          status: 'info',
          isClosable: true,
          duration: 9000,
        });
        router.replace('/login');
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

  return (
    <Flex
      justifyContent="center"
      alignItems="center"
      flexDirection="column"
      width="100%"
      height="80vh"
      px={{base: '1rem', md: '5rem'}}
    >
      <Heading fontSize={{base: '1.5rem', md: '2rem'}}>
        パスワードを再設定します
      </Heading>
      <Box width={{base: '100%', lg: '800px'}} mt="2rem">
        <FormProvider {...methods}>
          <form onSubmit={handleSubmit(submit)}>
            <Password setOk={setPWOK}>
              再設定するパスワード（8文字以上128文字以下）
            </Password>
            <Button
              marginTop="1rem"
              colorScheme="blue"
              isLoading={isSubmitting}
              type="submit"
              width={{base: '100%', md: 'auto'}}
            >
              パスワードを変更する
            </Button>
          </form>
        </FormProvider>
      </Box>
    </Flex>
  );
};

export default AcceptPage;
