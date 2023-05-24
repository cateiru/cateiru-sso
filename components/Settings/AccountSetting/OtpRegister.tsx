import {
  Box,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalHeader,
  ModalOverlay,
  useToast,
} from '@chakra-ui/react';
import React from 'react';
import {useSWRConfig} from 'swr';
import {z} from 'zod';
import {
  AccountOTPPublic,
  AccountOTPPublicSchema,
} from '../../../utils/types/account';
import {ErrorUniqueMessage} from '../../../utils/types/error';
import {useRequest} from '../../Common/useRequest';
import {OtpBackups} from './OtpBackups';
import {OtpRegisterForm, OtpRegisterFormData} from './OtpRegisterForm';
import {OtpRegisterReadQR} from './OtpRegisterReadQR';
import {OtpRegisterStart} from './OtpRegisterStart';

interface Props {
  isOpen: boolean;
  onClose: () => void;
}

export const OtpRegister = React.memo<Props>(props => {
  const toast = useToast();
  const {mutate} = useSWRConfig();

  const [step, setStep] = React.useState(0);
  const [token, setToken] = React.useState<AccountOTPPublic | null>(null);
  const [backups, setBackups] = React.useState<string[]>([]);

  const {request: otpRequest} = useRequest('/v2/account/otp', {
    customError: e => {
      const message = e.unique_code
        ? ErrorUniqueMessage[e.unique_code] ?? e.message
        : e.message;

      toast({
        title: message,
        status: 'error',
      });

      // POST時のみ動作する
      if (e.unique_code !== 8 && step === 1) {
        // OTPが間違っている場合以外はリセット
        setToken(null);
        setStep(0);
      }
    },
  });

  const C = React.useCallback(() => {
    switch (step) {
      case 0:
        return <OtpRegisterStart onRegisterStart={getBeginToken} />;
      case 1:
        if (!token) throw new Error('token is null');
        return (
          <Box>
            <OtpRegisterReadQR token={token.public_key} />
            <Box mt="1rem">
              <OtpRegisterForm onSubmit={submitOtp} />
            </Box>
          </Box>
        );
      case 2:
        return <OtpBackups backups={backups} title="バックアップコード" />;
      default:
        return null;
    }
  }, [step]);

  // OTPのPublic keyを取得するメソッド
  const getBeginToken = async () => {
    const res = await otpRequest({
      method: 'GET',
      mode: 'cors',
      credentials: 'include',
    });

    if (res) {
      const data = AccountOTPPublicSchema.safeParse(await res.json());
      if (data.success) {
        setToken(data.data);
        setStep(1);
      } else {
        alert(data.error);
      }
    }
  };

  const submitOtp = async (data: OtpRegisterFormData) => {
    if (!token) throw new Error('token is null');

    const form = new FormData();
    form.append('otp_session', token.otp_session);
    form.append('code', data.code);

    const res = await otpRequest({
      method: 'POST',
      mode: 'cors',
      credentials: 'include',
      body: form,
    });

    if (res) {
      const data = z.array(z.string()).safeParse(await res.json());
      if (data.success) {
        setBackups(data.data);
        setStep(2);
      } else {
        console.error(data.error);
      }
    }
  };

  const modalClose = () => {
    if (step === 2) {
      // パージする
      mutate(
        key =>
          typeof key === 'string' && key.startsWith('/v2/account/certificates'),
        undefined,
        {revalidate: true}
      );
    }

    setStep(0);
    setToken(null);
    setBackups([]);

    props.onClose();
  };

  return (
    <Modal isOpen={props.isOpen} onClose={modalClose} isCentered size="lg">
      <ModalOverlay />
      <ModalContent>
        <ModalHeader>二段階認証を設定します</ModalHeader>
        <ModalCloseButton size="lg" />
        <ModalBody mb="1rem">
          <C />
        </ModalBody>
      </ModalContent>
    </Modal>
  );
});

OtpRegister.displayName = 'OtpRegister';
