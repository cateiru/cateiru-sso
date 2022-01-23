import {
  InputGroup,
  Button,
  InputRightElement,
  Input,
  IconButton,
  FormControl,
  FormLabel,
  FormErrorMessage,
  Box,
} from '@chakra-ui/react';
import React from 'react';
import {useForm} from 'react-hook-form';
import type {FieldValues} from 'react-hook-form';
import {IoEyeOutline, IoEyeOffOutline} from 'react-icons/io5';
import PasswordChecklist from 'react-password-checklist';
import PasswordStrengthBar from 'react-password-strength-bar';

const UserPassword = () => {
  const {
    handleSubmit,
    register,
    formState: {errors, isSubmitting},
  } = useForm();

  const [show, setShow] = React.useState(false);
  const [pass, setPass] = React.useState('');
  const [pwOK, setPwOK] = React.useState(false);

  const submit = (values: FieldValues) => {
    return new Promise(() => {
      setTimeout(() => {
        alert(JSON.stringify(values, null, 2));
      }, 3000);
    });
  };

  console.log(errors.email);

  return (
    <Box width={{base: '90%', md: '600px'}}>
      <form onSubmit={handleSubmit(submit)}>
        <FormControl isInvalid={errors.email}>
          <FormLabel htmlFor="email">メールアドレス</FormLabel>
          <Input
            id="email"
            type="email"
            placeholder="メールアドレス"
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
        <FormControl isInvalid={errors.password}>
          <FormLabel htmlFor="password" marginTop="1rem">
            パスワード（12文字以上128文字以下）
          </FormLabel>
          <InputGroup>
            <Input
              id="password"
              type={show ? 'text' : 'password'}
              placeholder="パスワード"
              {...register('password', {
                required: true,
                onChange: e => setPass(e.target.value || ''),
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
          <PasswordStrengthBar
            password={pass}
            scoreWords={[
              '短すぎかな',
              '弱いパスワードだと思う',
              '少し弱いパスワードかなと思う',
              'もう少し長くしてみない？',
              '最強!すごく良いよ!',
            ]}
            shortScoreWord="8文字以上にしてみよう"
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
              if (pwOK !== isValid) {
                setPwOK(isValid);
              }
            }}
          />
        </Box>
        <Button
          marginTop="1rem"
          colorScheme="blue"
          isLoading={isSubmitting}
          type="submit"
          disabled={!pwOK}
        >
          アカウントを作る
        </Button>
      </form>
    </Box>
  );
};

export default UserPassword;