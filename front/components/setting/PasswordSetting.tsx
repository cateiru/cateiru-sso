import {
  Box,
  Center,
  Heading,
  Divider,
  Input,
  FormControl,
  FormErrorMessage,
  Button,
  useToast,
  Text,
  InputRightElement,
  InputGroup,
  IconButton,
  FormLabel,
} from '@chakra-ui/react';
import dynamic from 'next/dynamic';
import React from 'react';
import {useForm} from 'react-hook-form';
import type {FieldValues} from 'react-hook-form';
import {IoEyeOutline, IoEyeOffOutline} from 'react-icons/io5';
import PasswordStrengthBar from 'react-password-strength-bar';
import {changePassword} from '../../utils/api/change';

const PasswordChecklist = dynamic(() => import('react-password-checklist'), {
  ssr: false,
});

const PasswordSetting = () => {
  return (
    <Center height={{base: 'auto', md: '50vh'}} mx=".5rem">
      <Box
        width={{base: '100%', sm: '90%', md: '600px'}}
        mt={{base: '3rem', md: '0'}}
      >
        <Box>
          <Heading fontSize="1.5rem" mb="1.5rem" textAlign="center">
            ワンタイムパスワードを設定、変更する
          </Heading>
        </Box>
        <Divider />
        <Box mt="1.7rem">
          <Heading fontSize="1.5rem" mb="1.5rem" textAlign="center">
            パスワードを変更する
          </Heading>
          <ChangePassword />
        </Box>
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
  const [pass, setPass] = React.useState('');
  const [pwOk, setPWOK] = React.useState(false);

  const [showOld, setShowOld] = React.useState(false);
  const [showNew, setShowNew] = React.useState(false);

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
      <FormControl isInvalid={errors.newPassword} mt="1rem">
        <FormLabel htmlFor="newPassword" marginTop="1rem">
          新しいパスワード（8文字以上128文字以下）
        </FormLabel>
        <InputGroup>
          <Input
            id="newPassword"
            type={showNew ? 'text' : 'password'}
            placeholder="パスワード"
            {...register('newPassword', {
              required: true,
              onChange: e => setPass(e.target.value || ''),
            })}
          />
          <InputRightElement>
            <IconButton
              variant="ghost"
              aria-label="show password"
              icon={
                showNew ? (
                  <IoEyeOutline size="25px" />
                ) : (
                  <IoEyeOffOutline size="25px" />
                )
              }
              size="sm"
              onClick={() => setShowNew(!showNew)}
            />
          </InputRightElement>
        </InputGroup>
        <PasswordStrengthBar
          password={pass}
          scoreWords={[
            '弱すぎかな',
            '弱いパスワードだと思う',
            '少し弱いパスワードかなと思う',
            'もう少し長くしてみない？',
            '最強!すごく良いよ!',
          ]}
          shortScoreWord="8文字以上にしてほしいな"
          minLength={8}
        />
        <FormErrorMessage>{errors.password}</FormErrorMessage>
      </FormControl>
      <Box marginTop=".5rem">
        <PasswordChecklist
          rules={['minLength', 'specialChar', 'number', 'capital']}
          minLength={8}
          value={pass}
          messages={{
            minLength: 'パスワードは8文字以上',
            specialChar: 'パスワードに記号が含まれている',
            number: 'パスワードに数字が含まれている',
            capital: 'パスワードに大文字が含まれている',
          }}
          onChange={isValid => {
            let unmounted = false;
            if (pwOk !== isValid && !unmounted) {
              setPWOK(isValid);
            }

            return () => {
              unmounted = true;
            };
          }}
        />
      </Box>
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
  );
};

export default PasswordSetting;
