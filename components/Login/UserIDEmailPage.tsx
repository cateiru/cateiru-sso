import {useToast} from '@chakra-ui/react';
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

  const {isConditionSupported, onClickWebAuthn} = useWebAuthn(user => {
    // ログインする
    setUser({
      user: user,
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
          });
          props.setStep(LoginStep.CompleteLogin);
        } else if (data.data.otp) {
          // OTPの認証が必要な場合はOTPのページに遷移
          props.setOTPToken(data.data.otp.token);
          props.setLoginUser(data.data.otp.login_user);
          props.setStep(LoginStep.OTP);
        }
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
    </Margin>
  );
};
