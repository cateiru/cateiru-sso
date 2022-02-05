import {
  FormControl,
  FormLabel,
  Input,
  InputGroup,
  InputRightElement,
  Box,
  IconButton,
  FormErrorMessage,
} from '@chakra-ui/react';
import dynamic from 'next/dynamic';
import React from 'react';
import {useForm} from 'react-hook-form';
import type {FieldValues, UseFormRegister} from 'react-hook-form';
import {IoEyeOutline, IoEyeOffOutline} from 'react-icons/io5';
import PasswordStrengthBar from 'react-password-strength-bar';

const PasswordChecklist = dynamic(() => import('react-password-checklist'), {
  ssr: false,
});

const Password: React.FC<{
  errors: {[x: string]: any};
  register: UseFormRegister<FieldValues>;
  onChange: (status: boolean) => void,
  label?: string,
}> = ({errors, register, onChange, label = "password", children}) => {
  const [show, setShow] = React.useState(false);
  const [pwOk, setPWOK] = React.useState(false);
  const [pass, setPass] = React.useState('');

  React.useEffect(() => {
    onChange(pwOk)
  }, [pwOk])

  return (
    <>
      <FormControl isInvalid={errors[label]} mt="1rem">
        <FormLabel htmlFor={label} marginTop="1rem">
          {children}
        </FormLabel>
        <InputGroup>
          <Input
            id={label}
            type={show ? 'text' : 'password'}
            placeholder="パスワード"
            {...register(label, {
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
    </>
  );
};

export default Password;
