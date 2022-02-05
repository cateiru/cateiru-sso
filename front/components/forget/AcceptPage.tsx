import {Flex, Heading, Box, Button, useToast} from '@chakra-ui/react';
import {useRouter} from 'next/router';
import React from 'react';
import {useForm} from 'react-hook-form';
import type {FieldValues} from 'react-hook-form';
import {acceptPassword} from '../../utils/api/forget';
import Password from '../common/form/Password';

const AcceptPage: React.FC<{token: string}> = ({token}) => {
  const {
    handleSubmit,
    register,
    formState: {errors, isSubmitting},
  } = useForm();
  const toast = useToast();
  const [pwOk, setPWOK] = React.useState(false);
  const router = useRouter();

  const submit = (values: FieldValues) => {
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
      <Heading>パスワードを再設定します</Heading>
      <Box width={{base: '100%', lg: '800px'}} mt="2rem">
        <form onSubmit={handleSubmit(submit)}>
          <Password
            errors={errors}
            register={register}
            onChange={status => setPWOK(status)}
          >
            再設定するパスワード（8文字以上128文字以下）
          </Password>
          <Button
            marginTop="1rem"
            colorScheme="blue"
            isLoading={isSubmitting}
            type="submit"
            width={{base: '100%', md: 'auto'}}
            disabled={!pwOk}
          >
            パスワードを変更する
          </Button>
        </form>
      </Box>
    </Flex>
  );
};

export default AcceptPage;
