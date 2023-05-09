import {Button, useToast} from '@chakra-ui/react';
import {
  type RegistrationPublicKeyCredential,
  create,
  parseCreationOptionsFromJSON,
} from '@github/webauthn-json/browser-ponyfill';
import React from 'react';
import {useRequest} from '../Common/useRequest';

interface Props {
  registerToken: string;
  onSubmit: (data: RegistrationPublicKeyCredential) => Promise<void>;
}

export const RegisterPasskeyForm: React.FC<Props> = props => {
  const toast = useToast();
  const {request} = useRequest('/v2/register/begin_webauthn');
  const [credential, setCredential] = React.useState<
    CredentialCreationOptions | undefined
  >();

  React.useEffect(() => {
    const f = async () => {
      const res = await request({
        method: 'POST',
        credentials: 'include',
        mode: 'cors',
        headers: {
          'X-Register-Token': props.registerToken,
        },
      });

      if (!res) {
        return;
      }

      const data = parseCreationOptionsFromJSON(await res.json());
      setCredential(data);
    };
    f();
  }, []);

  const onClickHandler = async () => {
    if (!credential) {
      toast({
        title: 'パスキーは現在使用できません',
        status: 'error',
      });
      return;
    }

    let c: RegistrationPublicKeyCredential;
    try {
      c = await create(credential);
    } catch (e) {
      console.log(e);
      toast({
        title: '生体認証の登録に失敗しました',
        status: 'error',
      });
      return;
    }

    await props.onSubmit(c);
  };

  return (
    <Button colorScheme="cateiru" w="100%" onClick={onClickHandler}>
      生体認証を追加する
    </Button>
  );
};