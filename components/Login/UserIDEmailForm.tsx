import {
  Button,
  Center,
  FormControl,
  FormErrorMessage,
  FormLabel,
  Input,
  Link,
  useColorModeValue,
} from '@chakra-ui/react';
import NextLink from 'next/link';
import React from 'react';
import {FormProvider, useForm} from 'react-hook-form';
import {type LoginUser, LoginUserSchema} from '../../utils/types/login';
import {Avatar} from '../Common/Avatar';
import {PasswordForm, type PasswordFormData} from '../Common/Form/PasswordForm';
import {useRequest} from '../Common/useRequest';

const emailRegex = /[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}/;
const userIDRegex = /[A-Z0-9_]{3,15}/;
const userIdEmailRegex = new RegExp(
  `^(${userIDRegex.source})|(${emailRegex.source})$`,
  'i'
);

export interface UserIDEmailForm extends PasswordFormData {
  user_id_email: string;
}

interface Props {
  onSubmit: (data: UserIDEmailForm) => Promise<void>;
  onClickWebAuthn: () => Promise<void>;
  isConditionSupported: boolean;
}

export const UserIDEmailForm: React.FC<Props> = ({
  onSubmit,
  isConditionSupported,
  onClickWebAuthn,
}) => {
  const buttonColor = useColorModeValue('gray.500', 'gray.400');

  const methods = useForm<UserIDEmailForm>();
  const {
    handleSubmit,
    register,
    formState: {errors, isSubmitting},
  } = methods;

  const {request} = useRequest('/v2/login/user');
  const [user, setUser] = React.useState<LoginUser | null>(null);
  const [userName, setUserName] = React.useState<string>('');

  const handleRequestUserInfo = React.useCallback(() => {
    const f = async () => {
      if (userName.length === 0) return;

      const form = new FormData();
      form.append('username_or_email', userName);
      const res = await request({
        method: 'POST',
        body: form,
        mode: 'cors',
        credentials: 'include',
      });

      if (!res) return;

      const data = LoginUserSchema.safeParse(await res.json());
      if (data.success) {
        setUser(data.data);
      } else {
        console.error(data.error);
      }
    };
    f();
  }, [userName]);

  return (
    <>
      <Center>
        <Avatar src={user?.avatar ?? ''} size="xl" mb="1rem" />
      </Center>
      <FormProvider {...methods}>
        <form onSubmit={handleSubmit(onSubmit)}>
          <FormControl isInvalid={!!errors.user_id_email}>
            <FormLabel htmlFor="user_id_email">
              ユーザーIDまたはメールアドレス
            </FormLabel>
            <Input
              id="user_id_email"
              type="email text"
              autoComplete="username webauthn"
              {...register('user_id_email', {
                required: 'この値は必須です',
                pattern: {
                  value: userIdEmailRegex,
                  message: '正しい形式で入力してください',
                },
                onChange: e => setUserName(e.target.value),
                onBlur: handleRequestUserInfo,
              })}
            />
            <FormErrorMessage>
              {errors.user_id_email && errors.user_id_email.message}
            </FormErrorMessage>
          </FormControl>
          <PasswordForm enableWebauthn />
          <Link as={NextLink} href="/forget_password" color={buttonColor}>
            パスワードを忘れましたか？
          </Link>
          <Button
            mt="1rem"
            isLoading={isSubmitting}
            colorScheme="cateiru"
            type="submit"
            w="100%"
          >
            ログイン
          </Button>
          {isConditionSupported || (
            <Button w="100%" mt="1rem" onClick={onClickWebAuthn}>
              生体認証でログイン
            </Button>
          )}
        </form>
      </FormProvider>
    </>
  );
};
