import {Center, Text, useColorModeValue, useToast} from '@chakra-ui/react';
import React from 'react';
import {useSetRecoilState} from 'recoil';
import {UserState} from '../../utils/state/atom';
import {ErrorUniqueMessage} from '../../utils/types/error';
import {LoginResponseSchema, LoginUser} from '../../utils/types/login';
import {Avatar} from '../Common/Chakra/Avatar';
import {Margin} from '../Common/Margin';
import {useRecaptcha} from '../Common/useRecaptcha';
import {useRequest} from '../Common/useRequest';
import {LoginStep, type DefaultPageProps} from './Login';
import {OtpForm, OtpFormData} from './OtpForm';

interface Props extends DefaultPageProps {
  loginUser: LoginUser | null;
  otpToken: string;
}

export const OtpPage: React.FC<Props> = props => {
  const descriptionTextColor = useColorModeValue('gray.500', 'gray.400');
  const accentColor = useColorModeValue('my.primary', 'my.secondary');

  const setUser = useSetRecoilState(UserState);
  const {getRecaptchaToken} = useRecaptcha();
  const toast = useToast();

  const {request} = useRequest('/v2/login/otp', {
    customError: e => {
      const message = e.unique_code
        ? ErrorUniqueMessage[e.unique_code] ?? e.message
        : e.message;

      toast({
        title: message,
        status: 'error',
      });

      if (e.unique_code !== 8) {
        // OTPが間違っている場合以外はリセット
        props.reset();
      }
    },
    errorCallback: () => {
      props.reset();
    },
  });

  const onSubmit = async (data: OtpFormData, reset: () => void) => {
    const form = new FormData();
    form.append('otp_session', props.otpToken);
    form.append('code', data.otp);

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
    });

    if (res) {
      const data = LoginResponseSchema.safeParse(await res.json());

      if (data.success) {
        if (data.data.user) {
          // ログインする
          setUser({
            user: data.data.user,
            is_staff: false, // TODO: サーバーから取得できるようにする
            joined_organization: false,
          });
          props.setStep(LoginStep.CompleteLogin);
          return;
        }
      } else {
        console.error(data.error);
      }

      toast({
        title: 'ログインに失敗しました',
        status: 'error',
      });
    }

    reset();
  };

  return (
    <Margin>
      <Center mt="1rem">
        <Avatar src={props.loginUser?.avatar ?? ''} size="xl" />
      </Center>
      <Text fontSize="1.5rem" fontWeight="bold" textAlign="center" mt="1rem">
        <Text as="span" color={accentColor}>
          ワンタイムパスワード
        </Text>{' '}
        を入力
      </Text>
      <Text mb="1rem" textAlign="center" color={descriptionTextColor}>
        Authenticatorアプリで表示された6桁の数字を入力してください。
      </Text>
      <OtpForm onSubmit={onSubmit} />
    </Margin>
  );
};
