import {Box, Center, Text} from '@chakra-ui/react';
import {useSteps} from 'chakra-ui-steps';
import {useRouter, useSearchParams} from 'next/navigation';
import React from 'react';
import {CreateAccountRegisterEmailResponseSchema} from '../../utils/types/createAccount';
import {Spinner} from '../Common/Icons/Spinner';
import {useSecondaryColor} from '../Common/useColor';
import {useRequest} from '../Common/useRequest';
import {CompleteRegisterPage} from './CompleteRegisterPage';
import {EmailInputPage} from './EmailInputPage';
import {EmailVerifyPage} from './EmailVerifyPage';
import {RegisterCertificatePage} from './RegisterCertificatePage';
import {type StepStatus, Steps} from './Steps';

export enum RegisterAccountStep {
  EmailInput = 0,
  EmailVerify = 1,
  RegisterCertificate = 2,
  CompleteRegister = 3,
}

export interface DefaultPageProps {
  nextStep: () => void;
  prevStep: () => void;
  setStatus: (status: StepStatus) => void;
  reset: () => void;
}

export const RegisterAccount = () => {
  const prams = useSearchParams();
  const router = useRouter();
  const textColor = useSecondaryColor();

  const {request} = useRequest('/v2/register/invite_generate_session');

  const [registerToken, setRegisterToken] = React.useState<string | null>(null);
  const {
    activeStep,
    nextStep,
    prevStep,
    setStep,
    reset: resetStep,
  } = useSteps({
    initialStep: RegisterAccountStep.EmailInput,
  });
  const [status, setStatus] = React.useState<StepStatus>(undefined);
  const [isInvite, setIsInvite] = React.useState<boolean>(false);
  const [email, setEmail] = React.useState<string | null>(null);

  // invite_token が存在する場合、それを使用して登録用のセッションを取得する
  React.useEffect(() => {
    const token = prams.get('invite_token');
    const email = prams.get('email');

    const f = async () => {
      if (typeof token !== 'string') return;
      if (typeof email !== 'string') return;

      // 重複して送信しないようにする
      if (isInvite) return;
      // 成功、失敗に関わらずにフラグを立てる
      setIsInvite(true);

      const form = new FormData();

      form.append('invite_token', token);
      form.append('email', email);

      const res = await request({
        method: 'POST',
        mode: 'cors',
        credentials: 'include',
        body: form,
      });

      if (res) {
        const data = CreateAccountRegisterEmailResponseSchema.safeParse(
          await res.json()
        );
        if (data.success) {
          setEmail(email);
          setRegisterToken(data.data.register_token);
          setStep(RegisterAccountStep.RegisterCertificate);
          return;
        }
        console.error(data.error);
      }

      router.replace('/');
    };
    f();
  }, [prams.get('invite_token'), prams.get('email')]);

  const reset = React.useCallback(() => {
    setRegisterToken(null);
    setStatus(undefined);
    setEmail(null);

    resetStep();
  }, []);

  const C = React.useCallback(() => {
    switch (activeStep) {
      case RegisterAccountStep.EmailInput:
        // invite_tokenを使用した場合はスピナーを表示させる
        if (isInvite) {
          return (
            <Center>
              <Spinner />
            </Center>
          );
        }

        return (
          <EmailInputPage
            nextStep={nextStep}
            prevStep={prevStep}
            setStatus={setStatus}
            reset={reset}
            setRegisterToken={setRegisterToken}
            setEmail={setEmail}
          />
        );
      case RegisterAccountStep.EmailVerify:
        if (!registerToken) {
          reset();
          return null;
        }
        return (
          <EmailVerifyPage
            nextStep={nextStep}
            prevStep={prevStep}
            setStatus={setStatus}
            reset={reset}
            registerToken={registerToken}
          />
        );
      case RegisterAccountStep.RegisterCertificate:
        if (!registerToken) {
          reset();
          return null;
        }
        return (
          <RegisterCertificatePage
            nextStep={nextStep}
            prevStep={prevStep}
            setStatus={setStatus}
            reset={reset}
            registerToken={registerToken}
          />
        );
      case RegisterAccountStep.CompleteRegister:
        if (!registerToken) {
          reset();
          return null;
        }
        return <CompleteRegisterPage />;
    }
    return <></>;
  }, [activeStep]);

  return (
    <Box>
      {email && (
        <Text textAlign="center" color={textColor}>
          {email}
        </Text>
      )}
      <Steps activeStep={activeStep} state={status} />
      <Center mt="1rem" h="50%">
        <C />
      </Center>
    </Box>
  );
};
