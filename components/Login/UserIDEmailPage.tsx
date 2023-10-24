import {Button, Center, Divider, useToast} from '@chakra-ui/react';
import {useSetAtom} from 'jotai';
import React from 'react';
import {UserState} from '../../utils/state/atom';
import {ErrorUniqueMessage} from '../../utils/types/error';
import {LoginResponseSchema, LoginUser} from '../../utils/types/login';
import {Margin} from '../Common/Margin';
import {Link} from '../Common/Next/Link';
import {useSecondaryColor} from '../Common/useColor';
import {useRecaptcha} from '../Common/useRecaptcha';
import {useRequest} from '../Common/useRequest';
import {type DefaultPageProps, LoginStep} from './Login';
import {UserIDEmailForm} from './UserIDEmailForm';
import {useGetOauthLoginSession} from './useGetOauthLoginSession';
import {useWebAuthn} from './useWebAuthn';

interface Props extends DefaultPageProps {
  setLoginUser: (user: LoginUser) => void;
  setOTPToken: (token: string) => void;
}

export const UserIDEmailPage: React.FC<Props> = props => {
  const buttonColor = useSecondaryColor();

  const setUser = useSetAtom(UserState);
  const {getRecaptchaToken} = useRecaptcha();
  const toast = useToast();
  const getOauthLoginSession = useGetOauthLoginSession();

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

  const {onClickWebAuthn} = useWebAuthn(user => {
    // ログインする
    setUser(user);
    props.setStep(LoginStep.CompleteLogin);
  });

  const onSubmit = async (data: UserIDEmailForm) => {
    const form = new FormData();
    form.append('username_or_email', data.user_id_email);
    form.append('password', data.password);

    const recaptchaToken = await getRecaptchaToken();
    if (typeof recaptchaToken === 'undefined') {
      return;
    }
    form.append('recaptcha', recaptchaToken);

    const res = await request({
      method: 'POST',
      body: form,
      credentials: 'include',
      mode: 'cors',
      headers: {
        ...getOauthLoginSession(),
      },
    });

    if (res) {
      const data = LoginResponseSchema.safeParse(await res.json());

      if (data.success) {
        if (data.data.user) {
          // ログインする
          setUser(data.data.user);
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
      <UserIDEmailForm onSubmit={onSubmit} onClickWebAuthn={onClickWebAuthn} />
      <Divider my="1rem" />
      <Center>
        <Button variant="link" as={Link} href="/register" color={buttonColor}>
          アカウントを作成
        </Button>
      </Center>
    </Margin>
  );
};
