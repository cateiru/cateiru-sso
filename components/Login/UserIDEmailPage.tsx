import {
  Button,
  Center,
  Divider,
  useColorModeValue,
  useToast,
} from '@chakra-ui/react';
import Link from 'next/link';
import React from 'react';
import {useGoogleReCaptcha} from 'react-google-recaptcha-v3';
import {useSetRecoilState} from 'recoil';
import {UserState} from '../../utils/state/atom';
import {ErrorUniqueMessage} from '../../utils/types/error';
import {LoginResponseSchema, LoginUser} from '../../utils/types/login';
import {Margin} from '../Common/Margin';
import {useRequest} from '../Common/useRequest';
import {type DefaultPageProps, LoginStep} from './Login';
import {UserIDEmailForm} from './UserIDEmailForm';
import {useWebAuthn} from './useWebAuthn';

interface Props extends DefaultPageProps {
  setLoginUser: (user: LoginUser) => void;
  setOTPToken: (token: string) => void;
}

export const UserIDEmailPage: React.FC<Props> = props => {
  const buttonColor = useColorModeValue('gray.500', 'gray.400');

  const setUser = useSetRecoilState(UserState);
  const {executeRecaptcha} = useGoogleReCaptcha();
  const toast = useToast();

  const {request} = useRequest('/v2/login/password', {
    customError: e => {
      const message = e.unique_code
        ? ErrorUniqueMessage[e.unique_code] ?? e.message
        : e.message;

      toast({
        title: message,
        status: 'error',
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

  const {isConditionSupported, onClickWebAuthn} = useWebAuthn(user => {
    // ログインする
    setUser({
      user: user,
      is_staff: false,
    });
    props.setStep(LoginStep.CompleteLogin);
  });

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
            is_staff: false,
          });
          props.setStep(LoginStep.CompleteLogin);
        } else if (data.data.otp) {
          // OTPの認証が必要な場合はOTPのページに遷移
          props.setOTPToken(data.data.otp.token);
          props.setLoginUser(data.data.otp.login_user);
          props.setStep(LoginStep.OTP);
        }
      } else {
        console.error(data.error);
      }
    }
  };

  return (
    <Margin>
      <UserIDEmailForm
        onSubmit={onSubmit}
        isConditionSupported={isConditionSupported}
        onClickWebAuthn={onClickWebAuthn}
      />
      <Divider my="1rem" />
      <Center>
        <Button variant="link" as={Link} href="/register" color={buttonColor}>
          アカウントを作成
        </Button>
      </Center>
    </Margin>
  );
};
