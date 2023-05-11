import {Text, useColorModeValue, useToast} from '@chakra-ui/react';
import React from 'react';
import {useGoogleReCaptcha} from 'react-google-recaptcha-v3';
import {LoginUser, LoginUserSchema} from '../../utils/types/login';
import {Margin} from '../Common/Margin';
import {useRequest} from '../Common/useRequest';
import {type DefaultPageProps, LoginStep} from './Login';
import {UserIDEmailForm} from './UserIDEmailForm';
import {useWebAuthn} from './useWebAuthn';

interface Props extends DefaultPageProps {
  setLoginUser: (user: LoginUser) => void;
}

export const UserIDEmailPage: React.FC<Props> = props => {
  const accentColor = useColorModeValue('my.primary', 'my.secondary');
  const {executeRecaptcha} = useGoogleReCaptcha();
  const {request} = useRequest('/v2/login/user');
  const toast = useToast();

  const {isConditionSupported, onClickWebAuthn} = useWebAuthn();

  const onSubmit = async (data: UserIDEmailForm) => {
    if (!executeRecaptcha) {
      toast({
        title: 'reCAPTCHAの読み込みに失敗しました',
        status: 'error',
      });
      return;
    }

    const form = new FormData();
    form.append('username_or_email', data.user_id_email);

    try {
      form.append('recaptcha', await executeRecaptcha());
    } catch {
      toast({
        title: 'reCAPTCHAの読み込みに失敗しました',
        status: 'error',
      });
      return;
    }

    const res = await request({
      method: 'POST',
      body: form,
      credentials: 'include',
      mode: 'cors',
    });

    if (res) {
      const data = LoginUserSchema.safeParse(await res.json());
      if (data.success) {
        if (data.data.available_password) {
          props.setLoginUser(data.data);
          props.setStep(LoginStep.Password);
          return;
        }

        toast({
          title: 'パスワードが設定されていません',
          description:
            '生体認証でログインする場合は、入力候補から選択してください',
          status: 'error',
        });
        return;
      }
    }

    toast({
      title: 'ログインに失敗しました',
      status: 'error',
    });
  };

  return (
    <Margin>
      <Text fontSize="1.5rem" fontWeight="bold" mb="1rem" textAlign="center">
        <Text as="span" color={accentColor}>
          Email
        </Text>
        または、
        <Text as="span" color={accentColor}>
          ユーザーID
        </Text>
        を入力
      </Text>
      <UserIDEmailForm
        onSubmit={onSubmit}
        isConditionSupported={isConditionSupported}
        onClickWebAuthn={onClickWebAuthn}
      />
    </Margin>
  );
};
