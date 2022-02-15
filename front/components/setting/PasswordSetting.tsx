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
} from '@chakra-ui/react';
import {useRouter} from 'next/router';
import React from 'react';
import {useForm} from 'react-hook-form';
import type {FieldValues} from 'react-hook-form';
import {IoEyeOutline, IoEyeOffOutline} from 'react-icons/io5';
import {changePassword} from '../../utils/api/change';
import Password from '../common/form/Password';

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
  const {
    handleSubmit,
    register,
    formState: {errors, isSubmitting},
  } = useForm();
  const [pwOk, setPWOK] = React.useState(false);
  const router = useRouter();

  const [showOld, setShowOld] = React.useState(false);

  const toast = useToast();

  const submitHandler = (values: FieldValues) => {
    const f = async () => {
      try {
        await changePassword(values.oldPassword, values.newPassword);
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
    <form onSubmit={handleSubmit(submitHandler)}>
      <FormControl isInvalid={errors.oldPassword} mt="1rem">
        <FormLabel htmlFor="oldPassword" marginTop="1rem">
          今のパスワード
        </FormLabel>
        <InputGroup>
          <Input
            id="oldPassword"
            type={showOld ? 'text' : 'password'}
            placeholder="パスワード"
            {...register('oldPassword', {
              required: true,
            })}
          />
          <InputRightElement>
            <IconButton
              variant="ghost"
              aria-label="show password"
              icon={
                showOld ? (
                  <IoEyeOutline size="25px" />
                ) : (
                  <IoEyeOffOutline size="25px" />
                )
              }
              size="sm"
              onClick={() => setShowOld(!showOld)}
            />
          </InputRightElement>
        </InputGroup>
      </FormControl>
      <Password
        errors={errors}
        register={register}
        onChange={status => setPWOK(status)}
        label="newPassword"
      >
        新しいパスワード（8文字以上128文字以下）
      </Password>
      <Button
        marginTop="1rem"
        colorScheme="blue"
        isLoading={isSubmitting}
        type="submit"
        width={{base: '100%', md: 'auto'}}
        disabled={!pwOk}
      >
        変更する
      </Button>
    </form>
  );
};

export default PasswordSetting;
