import {Flex} from '@chakra-ui/react';
import {useRouter} from 'next/router';
import React from 'react';
import {GoogleReCaptchaProvider} from 'react-google-recaptcha-v3';
import SelectCreate, {CreateType} from './SelectCreate';

const reCAPTCHA = process.env.NEXT_PUBLIC_RE_CAPTCHA;

const CreateAccountPage: React.FC = () => {
  const [mailToken, setMailToken] = React.useState('');
  const [selectType, setSelectType] = React.useState(CreateType.Initialize);

  const router = useRouter();
  React.useEffect(() => {
    if (!router.isReady) return;
    const query = router.query;

    if (typeof query['token'] === 'string') {
      setMailToken(query['token']);
      setSelectType(CreateType.ValidateMail);
    }
  }, [router.isReady, router.query]);

  return (
    <Flex
      justifyContent="center"
      alignItems="center"
      flexDirection="column"
      width="100%"
      height="80vh"
      px={{base: '1rem', md: '5rem'}}
    >
      <GoogleReCaptchaProvider
        reCaptchaKey={reCAPTCHA}
        language="ja"
        scriptProps={{
          async: false,
          defer: false,
          appendTo: 'head',
          nonce: undefined,
        }}
      >
        <SelectCreate
          selectType={selectType}
          mailToken={mailToken}
          setSelectType={setSelectType}
        />
      </GoogleReCaptchaProvider>
    </Flex>
  );
};

export default CreateAccountPage;
