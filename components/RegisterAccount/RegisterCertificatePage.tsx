import {Box, Button, Center, Heading} from '@chakra-ui/react';
import type {RegistrationPublicKeyCredential} from '@github/webauthn-json/dist/types/browser-ponyfill';
import {useSetAtom} from 'jotai';
import React from 'react';
import {TbFingerprint, TbPassword} from 'react-icons/tb';
import {UserState} from '../../utils/state/atom';
import {UserMeSchema} from '../../utils/types/user';
import {RegisterPasswordForm} from '../Common/Form/RegisterPasswordForm';
import {RegisterPasswordFormContextData} from '../Common/Form/RegisterPasswordFormContext';
import {useSecondaryColor} from '../Common/useColor';
import {useRequest} from '../Common/useRequest';
import type {DefaultPageProps} from './RegisterAccount';
import {RegisterPasskeyForm} from './RegisterPasskeyForm';

interface Props extends DefaultPageProps {
  registerToken: string;
}

export const RegisterCertificatePage: React.FC<Props> = props => {
  const setUser = useSetAtom(UserState);
  const buttonColor = useSecondaryColor();

  const [certType, setCertType] = React.useState<'passkey' | 'password'>(
    'passkey'
  );
  const {request: requestPassword} = useRequest('/register/password', {
    errorCallback: () => {
      props.setStatus('error');
    },
  });
  const {request: requestPasskey} = useRequest('/register/webauthn', {
    errorCallback: () => {
      props.setStatus('error');
    },
  });

  const onSubmitPassword = async (data: RegisterPasswordFormContextData) => {
    const form = new FormData();
    form.append('password', data.new_password);

    const res = await requestPassword({
      method: 'POST',
      body: form,
      headers: {
        'X-Register-Token': props.registerToken,
      },
    });

    if (res) {
      await afterRegister(res);
      return;
    }

    props.setStatus('error');
  };

  const onSubmitPasskey = async (data: RegistrationPublicKeyCredential) => {
    const res = await requestPasskey({
      method: 'POST',
      body: JSON.stringify(data),
      headers: {
        'X-Register-Token': props.registerToken,
        'Content-Type': 'application/json',
      },
    });

    if (res) {
      await afterRegister(res);
      return;
    }

    props.setStatus('error');
  };

  const afterRegister = async (res: Response) => {
    const data = UserMeSchema.safeParse(await res.json());
    if (!data.success) {
      console.error(data.error);
      props.setStatus('error');
      return;
    }

    setUser(data.data);

    props.setStatus(undefined);
    props.nextStep();
  };

  return (
    <Box w={{base: '95%', md: '600px'}} m="auto" mt="2rem">
      <Box>
        <Heading fontSize="1.5rem" textAlign="center">
          {certType === 'passkey' ? '生体認証' : 'パスワード'}
          を使用して認証情報を追加します
        </Heading>
        <Center my="1rem">
          <Button
            leftIcon={
              certType !== 'passkey' ? (
                <TbFingerprint size="20px" />
              ) : (
                <TbPassword size="20px" />
              )
            }
            variant="link"
            color={buttonColor}
            onClick={() =>
              setCertType(certType === 'passkey' ? 'password' : 'passkey')
            }
          >
            {certType !== 'passkey'
              ? '生体認証で認証する'
              : 'パスワードで認証する'}
          </Button>
        </Center>

        {certType === 'passkey' ? (
          <RegisterPasskeyForm
            onSubmit={onSubmitPasskey}
            registerToken={props.registerToken}
          />
        ) : (
          <RegisterPasswordForm
            onSubmit={onSubmitPassword}
            buttonText="パスワードを設定"
          />
        )}
      </Box>
    </Box>
  );
};
