import {
  Box,
  Center,
  Heading,
  Input,
  FormControl,
  Button,
  useToast,
  InputRightElement,
  InputGroup,
  IconButton,
  FormLabel,
  FormErrorMessage,
} from '@chakra-ui/react';
import {useRouter} from 'next/router';
import React from 'react';
import {SubmitHandler, useForm, FormProvider} from 'react-hook-form';
import {TbEye, TbEyeOff} from 'react-icons/tb';
import {changePassword} from '../../utils/api/change';
import Password, {PasswordForm} from '../common/form/Password';

interface Form extends PasswordForm {
  oldPassword: string;
}

const PasswordSetting = () => {
  return (
    <Center mx=".5rem">
      <Box
        width={{base: '100%', sm: '90%', md: '600px'}}
        mt={{base: '3rem', md: '5rem'}}
      >
        <Heading fontSize="1.5rem" mb="1.5rem" textAlign="center">
          パスワードを変更する
        </Heading>
        <ChangePassword />
      </Box>
    </Center>
  );
};

const ChangePassword = () => {
  const methods = useForm<Form>();
  const {
    handleSubmit,
    register,
    setError,
    clearErrors,
    formState: {errors, isSubmitting},
  } = methods;

  const [pwOk, setPWOK] = React.useState(false);
  const router = useRouter();

  const [showOld, setShowOld] = React.useState(false);

  const toast = useToast();

  const submitHandler: SubmitHandler<Form> = values => {
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
        await changePassword(values.oldPassword, values.password);
        toast({
          title: 'パスワードを変更しました',
          status: 'info',
          isClosable: true,
          duration: 9000,
        });
        router.reload();
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

    f();

    return () => {};
  };

  return (
    <FormProvider {...methods}>
      <form onSubmit={handleSubmit(submitHandler)}>
        <FormControl isInvalid={Boolean(errors.oldPassword)} mt="1rem">
          <FormLabel htmlFor="oldPassword" marginTop="1rem">
            今のパスワード
          </FormLabel>
          <InputGroup>
            <Input
              id="oldPassword"
              type={showOld ? 'text' : 'password'}
              placeholder="パスワード"
              {...register('oldPassword', {
                required: 'パスワードは必須です',
              })}
            />
            <InputRightElement>
              <IconButton
                variant="ghost"
                aria-label="show password"
                icon={
                  showOld ? <TbEye size="25px" /> : <TbEyeOff size="25px" />
                }
                size="sm"
                onClick={() => setShowOld(!showOld)}
              />
            </InputRightElement>
          </InputGroup>
          <FormErrorMessage>
            {errors.oldPassword && errors.oldPassword.message}
          </FormErrorMessage>
        </FormControl>
        <Password setOk={setPWOK}>
          新しいパスワード（8文字以上128文字以下）
        </Password>
        <Button
          marginTop="1rem"
          colorScheme="blue"
          isLoading={isSubmitting}
          type="submit"
          width={{base: '100%', md: 'auto'}}
        >
          変更する
        </Button>
      </form>
    </FormProvider>
  );
};

export default PasswordSetting;
