import {Box, Button, Center, Heading} from '@chakra-ui/react';
import type {RegistrationPublicKeyCredential} from '@github/webauthn-json/dist/types/browser-ponyfill';
import React from 'react';
import {TbFingerprint, TbPassword} from 'react-icons/tb';
import {useRequest} from '../Common/useRequest';
import type {DefaultPageProps} from './RegisterAccount';
import {RegisterPasskeyForm} from './RegisterPasskeyForm';
import {PasswordForm, RegisterPasswordForm} from './RegisterPasswordForm';

interface Props extends DefaultPageProps {
  registerToken: string;
}

export const RegisterCertificatePage: React.FC<Props> = props => {
  const [certType, setCertType] = React.useState<'passkey' | 'password'>(
    'passkey'
  );
  const {request: requestPassword} = useRequest('/v2/register/password', {
    errorCallback: () => {
      props.setStatus('error');
    },
  });
  const {request: requestPasskey} = useRequest('/v2/register/webauthn', {
    errorCallback: () => {
      props.setStatus('error');
    },
  });

  const onSubmitPassword = async (data: PasswordForm) => {
    const form = new FormData();
    form.append('password', data.password);

    const res = await requestPassword({
      method: 'POST',
      credentials: 'include',
      mode: 'cors',
      body: form,
      headers: {
        'X-Register-Token': props.registerToken,
      },
    });

    if (res) {
      props.setStatus(undefined);
      props.nextStep();
      return;
    }

    props.setStatus('error');
  };

  const onSubmitPasskey = async (data: RegistrationPublicKeyCredential) => {
    const res = await requestPasskey({
      method: 'POST',
      credentials: 'include',
      mode: 'cors',
      body: JSON.stringify(data),
      headers: {
        'X-Register-Token': props.registerToken,
        'Content-Type': 'application/json',
      },
    });

    if (res) {
      props.setStatus(undefined);
      props.nextStep();
      return;
    }

    props.setStatus('error');
  };

  return (
    <Box w={{base: '95%', md: '600px'}} m="auto">
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
          <RegisterPasswordForm onSubmit={onSubmitPassword} />
        )}
      </Box>
    </Box>
  );
};
