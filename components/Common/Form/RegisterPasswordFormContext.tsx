import {
  FormControl,
  FormLabel,
  IconButton,
  Input,
  InputGroup,
  InputRightElement,
  Skeleton,
} from '@chakra-ui/react';
import dynamic from 'next/dynamic';
import React from 'react';
import {useFormContext} from 'react-hook-form';
import {TbEye, TbEyeOff} from 'react-icons/tb';
import PasswordStrengthBar from 'react-password-strength-bar';

const PasswordChecklist = dynamic(() => import('react-password-checklist'), {
  ssr: false,
});

export interface RegisterPasswordFormContextData {
  new_password: string;
}

interface Props {
  setOk: (ok: boolean) => void;
  ok: boolean;
  label?: string;
  mt?: string;
}

export const RegisterPasswordFormContext: React.FC<Props> = props => {
  const {
    register,
    setError,
    clearErrors,
    formState: {errors},
  } = useFormContext<RegisterPasswordFormContextData>();
  const [show, setShow] = React.useState(false);
  const [pass, setPass] = React.useState('');

  React.useEffect(() => {
    if (pass.length !== 0) {
      if (!props.ok) {
        setError('new_password', {});
      } else {
        clearErrors('new_password');
      }
    }
  }, [props.ok, pass]);

  return (
    <FormControl isInvalid={!!errors.new_password} mt={props.mt}>
      <FormLabel htmlFor="new_password">
        {props.label ?? 'パスワード'}
      </FormLabel>
      <InputGroup>
        <Input
          id="new_password"
          autoComplete="new-password"
          type={show ? 'text' : 'password'}
          {...register('new_password', {
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
        scoreWords={[]}
        shortScoreWord=""
        minLength={8}
      />
      <Skeleton
        marginTop=".5rem"
        minH="130px"
        isLoaded={!!PasswordChecklist}
        verticalAlign="middle"
      >
        <PasswordChecklist
          rules={['minLength', 'specialChar', 'number', 'capital']}
          minLength={13}
          value={pass}
          messages={{
            minLength: 'パスワードは13文字以上',
            specialChar: 'パスワードに記号が含まれている',
            number: 'パスワードに数字が含まれている',
            capital: 'パスワードに大文字が含まれている',
          }}
          onChange={isValid => {
            let unmounted = false;
            if (isValid !== props.ok && !unmounted) {
              props.setOk(isValid);
            }

            return () => {
              unmounted = true;
            };
          }}
        />
      </Skeleton>
    </FormControl>
  );
};
