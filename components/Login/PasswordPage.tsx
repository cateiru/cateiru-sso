import {
  Avatar,
  Center,
  Text,
  useColorModeValue,
  useToast,
} from '@chakra-ui/react';
import React from 'react';
import {useGoogleReCaptcha} from 'react-google-recaptcha-v3';
import {useSetRecoilState} from 'recoil';
import {UserState} from '../../utils/state/atom';
import {ErrorUniqueMessage} from '../../utils/types/error';
import {LoginResponseSchema, LoginUser} from '../../utils/types/login';
import {Margin} from '../Common/Margin';
import {useRequest} from '../Common/useRequest';
import {DefaultPageProps, LoginStep} from './Login';
import {PasswordForm, type PasswordFormData} from './PasswordForm';

interface Props extends DefaultPageProps {
  loginUser: LoginUser;
  setOTPToken: (token: string) => void;
}

export const PasswordPage: React.FC<Props> = props => {
  const toast = useToast();
  const accentColor = useColorModeValue('my.primary', 'my.secondary');
  const setUser = useSetRecoilState(UserState);
  const {executeRecaptcha} = useGoogleReCaptcha();
  const {request} = useRequest('/v2/login/password', {
    customError: e => {
      const message = e.unique_code
        ? ErrorUniqueMessage[e.unique_code] ?? e.message
        : e.message;

      toast({
        title: message,
        status: 'error',
        duration: 5000,
        isClosable: true,
      });

      if (e.unique_code !== 8) {
        // パスワードが間違っている場合以外はリセット
        props.reset();
      }
    },
    errorCallback: () => {
      props.reset();
    },
  });

  const onSubmit = async (data: PasswordFormData) => {
    if (!executeRecaptcha) {
      toast({
        title: 'reCAPTCHAの読み込みに失敗しました',
        status: 'error',
      });
      return;
    }

    const form = new FormData();
    form.append('username_or_email', props.loginUser.user_name);
    form.append('password', data.password);

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
      const data = LoginResponseSchema.safeParse(await res.json());

      if (data.success) {
        if (data.data.user) {
          // ログインする
          setUser({
            user: data.data.user,
          });
        } else if (data.data.otp) {
          // OTPの認証が必要な場合はOTPのページに遷移
          props.setOTPToken(data.data.otp);
          props.setStep(LoginStep.OTP);
        }
      }
    }
  };

  return (
    <Margin>
      <Center mt="1rem">
        <Avatar src={props.loginUser.avatar ?? ''} size="lg" />
      </Center>
      <Text fontSize="1.5rem" fontWeight="bold" mb="1rem" textAlign="center">
        <Text as="span" color={accentColor}>
          パスワード
        </Text>
        を入力
      </Text>
      <PasswordForm onSubmit={onSubmit} />
    </Margin>
  );
};
