import {Flex} from '@chakra-ui/react';
import {GoogleReCaptchaProvider} from 'react-google-recaptcha-v3';
import LoginForm from './LoginForm';

const reCAPTCHA = process.env.NEXT_PUBLIC_RE_CAPTCHA;

const LoginPage = () => {
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
        <LoginForm />
      </GoogleReCaptchaProvider>
    </Flex>
  );
};

export default LoginPage;
