import {
  Box,
  Heading,
  Flex,
  FormControl,
  Button,
  Input,
  FormErrorMessage,
  useToast,
} from '@chakra-ui/react';
import {useForm} from 'react-hook-form';
import type {FieldValues} from 'react-hook-form';
import {sendForget} from '../../utils/api/forget';

const Forget = () => {
  const {
    handleSubmit,
    register,
    formState: {errors, isSubmitting},
  } = useForm();
  const toast = useToast();

  const submit = (values: FieldValues) => {
    const f = async () => {
      try {
        await sendForget(values.email);

        toast({
          title: '再登録用のメールを送信しました',
          status: 'info',
          isClosable: true,
          duration: 9000,
        });
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
          >
            再登録メールを送信する
          </Button>
        </form>
      </Box>
    </Flex>
  );
};

export default Forget;
