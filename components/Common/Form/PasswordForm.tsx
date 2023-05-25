import {
  FormControl,
  FormErrorMessage,
  FormLabel,
  IconButton,
  Input,
  InputGroup,
  InputRightElement,
} from '@chakra-ui/react';
import React from 'react';
import {useFormContext} from 'react-hook-form';
import {TbEye, TbEyeOff} from 'react-icons/tb';

export interface PasswordFormData {
  password: string;
}

interface Props {
  enableWebauthn: boolean;
  label?: string;
}

export const PasswordForm: React.FC<Props> = props => {
  const {
    register,
    formState: {errors},
  } = useFormContext<PasswordFormData>();
  const [show, setShow] = React.useState(false);

  return (
    <FormControl isInvalid={!!errors.password} mt=".5rem">
      <FormLabel htmlFor="password">{props.label ?? 'パスワード'}</FormLabel>
      <InputGroup>
        <Input
          id="password"
          autoComplete={
            props.enableWebauthn
              ? 'current-password webauthn'
              : 'current-password'
          }
          type={show ? 'text' : 'password'}
          {...register('password', {
            required: `${props.label ?? 'パスワード'}を入力してください`,
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
      <FormErrorMessage>
        {errors.password && errors.password.message}
      </FormErrorMessage>
    </FormControl>
  );
};
