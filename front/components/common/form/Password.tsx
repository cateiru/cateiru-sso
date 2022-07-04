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
import {useFormContext} from 'react-hook-form';
import {TbEye, TbEyeOff} from 'react-icons/tb';
import PasswordStrengthBar from 'react-password-strength-bar';

const PasswordChecklist = dynamic(() => import('react-password-checklist'), {
  ssr: false,
});

export interface PasswordForm {
  password: string;
}

const Password: React.FC<{
  children: React.ReactNode;
  setOk: (ok: boolean) => void;
}> = ({children, setOk}) => {
  const {
    register,
    clearErrors,
    setError,
    formState: {errors},
  } = useFormContext<PasswordForm>();

  const [show, setShow] = React.useState(false);
  const [pwOk, setPWOK] = React.useState(false);
  const [pass, setPass] = React.useState('');

  React.useEffect(() => {
    if (pass.length !== 0) {
      if (!pwOk) {
        setError('password', {
          type: 'custom',
          message: 'custom message',
        });
      } else {
        clearErrors('password');
      }
    }
  }, [pwOk, pass]);

  return (
    <>
      <FormControl isInvalid={Boolean(errors.password)} mt="1rem">
        <FormLabel htmlFor="password" marginTop="1rem">
          {children}
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
              icon={show ? <TbEye size="25px" /> : <TbEyeOff size="25px" />}
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
            if (isValid !== pwOk && !unmounted) {
              setPWOK(isValid);
              setOk(isValid);
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
