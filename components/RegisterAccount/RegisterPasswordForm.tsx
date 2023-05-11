import {
  Box,
  Button,
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
import {useForm} from 'react-hook-form';
import {TbEye, TbEyeOff} from 'react-icons/tb';
import PasswordStrengthBar from 'react-password-strength-bar';

const PasswordChecklist = dynamic(() => import('react-password-checklist'), {
  ssr: false,
});

export interface PasswordForm {
  password: string;
}

interface Props {
  onSubmit: (data: PasswordForm) => Promise<void>;
}

export const RegisterPasswordForm: React.FC<Props> = props => {
  const {
    handleSubmit,
    register,
    setError,
    clearErrors,
    formState: {errors, isSubmitting},
  } = useForm<PasswordForm>();
  const [show, setShow] = React.useState(false);
  const [pass, setPass] = React.useState('');
  const [pwOk, setPWOK] = React.useState(false);

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

  const onSubmit = async (data: PasswordForm) => {
    if (!pwOk) {
      return;
    } else {
      clearErrors('password');
    }

    await props.onSubmit(data);
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <FormControl isInvalid={!!errors.password}>
        <FormLabel htmlFor="password">パスワード</FormLabel>
        <InputGroup>
          <Input
            id="password"
            autoComplete="new-password"
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
          scoreWords={[]}
          shortScoreWord=""
          minLength={8}
        />
        <Skeleton
          marginTop=".5rem"
          minH="150px"
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
              if (isValid !== pwOk && !unmounted) {
                setPWOK(isValid);
              }

              return () => {
                unmounted = true;
              };
            }}
          />
        </Skeleton>
      </FormControl>
      <Button
        mt="1rem"
        isLoading={isSubmitting}
        colorScheme="cateiru"
        type="submit"
        w="100%"
      >
        パスワードを登録する
      </Button>
    </form>
  );
};
