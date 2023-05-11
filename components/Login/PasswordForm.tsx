import {
  Button,
  FormControl,
  FormErrorMessage,
  IconButton,
  Input,
  InputGroup,
  InputRightElement,
} from '@chakra-ui/react';
import React from 'react';
import {useForm} from 'react-hook-form';
import {TbEye, TbEyeOff} from 'react-icons/tb';

export interface PasswordFormData {
  password: string;
}

interface Props {
  onSubmit: (data: PasswordFormData) => Promise<void>;
}

export const PasswordForm: React.FC<Props> = ({onSubmit}) => {
  const {
    handleSubmit,
    register,
    formState: {errors, isSubmitting},
  } = useForm<PasswordFormData>();
  const [show, setShow] = React.useState(false);

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <FormControl isInvalid={!!errors.password}>
        <InputGroup>
          <Input
            id="password"
            autoComplete="current-password"
            type={show ? 'text' : 'password'}
            {...register('password', {
              required: 'パスワードを入力してください',
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
      <Button
        mt="1rem"
        isLoading={isSubmitting}
        colorScheme="cateiru"
        type="submit"
        w="100%"
      >
        パスワードを認証する
      </Button>
    </form>
  );
};
