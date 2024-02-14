import {
  Button,
  Center,
  FormControl,
  FormErrorMessage,
  FormLabel,
  Input,
  Link,
} from '@chakra-ui/react';
import React from 'react';
import {FormProvider, useForm} from 'react-hook-form';
import {userIdEmailRegex} from '../../utils/regex';
import {type LoginUser, LoginUserSchema} from '../../utils/types/login';
import {Avatar} from '../Common/Chakra/Avatar';
import {PasswordForm, type PasswordFormData} from '../Common/Form/PasswordForm';
import {Link as NextLink} from '../Common/Next/Link';
import {useSecondaryColor} from '../Common/useColor';
import {useRequest} from '../Common/useRequest';

export interface UserIDEmailForm extends PasswordFormData {
  user_id_email: string;
}

interface Props {
  onSubmit: (data: UserIDEmailForm) => Promise<void>;
  onClickWebAuthn: () => Promise<void>;
}

export const UserIDEmailForm: React.FC<Props> = ({
  onSubmit,
  onClickWebAuthn,
}) => {
  const buttonColor = useSecondaryColor();

  const methods = useForm<UserIDEmailForm>();
  const {
    handleSubmit,
    register,
    setError,
    clearErrors,
    formState: {errors, isSubmitting},
  } = methods;

  const [user, setUser] = React.useState<LoginUser | null>(null);
  const [userName, setUserName] = React.useState<string>('');

  const {request} = useRequest('/login/user', {
    errorCallback: () => {
      setUser(null);
    },

    customError: e => {
      if (e.unique_code === 10) {
        setError('user_id_email', {
          message: 'このユーザーは存在しません',
        });
        setUser(null);
      }
    },
  });

  const handleRequestUserInfo = React.useCallback(() => {
    const f = async () => {
      if (userName.length === 0) {
        setUser(null);
        return;
      }

      const form = new FormData();
      form.append('username_or_email', userName);
      const res = await request({
        method: 'POST',
        body: form,
        mode: 'cors',
        credentials: 'include',
      });

      if (res) {
        const data = LoginUserSchema.safeParse(await res.json());
        if (data.success) {
          setUser(data.data);
          clearErrors('user_id_email');
        } else {
          console.error(data.error);
        }
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
          <PasswordForm enableWebauthn={false} />
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
          <Button w="100%" mt="1rem" onClick={onClickWebAuthn}>
            生体認証でログイン
          </Button>
        </form>
      </FormProvider>
    </>
  );
};
