import {useToast} from '@chakra-ui/react';
import {
  Box,
  Center,
  Avatar,
  Heading,
  Text,
  Button,
  Flex,
  Textarea,
  useClipboard,
} from '@chakra-ui/react';
import {useRouter} from 'next/router';
import React from 'react';
import {IoCheckmarkOutline} from 'react-icons/io5';
import {useRecoilValue, useSetRecoilState} from 'recoil';
import {login, preview, ServicePreview} from '../../../utils/api/loginSSO';
import {OIDCRequestQuery} from '../../../utils/sso/login';
import {UserState, LoadState} from '../../../utils/state/atom';

const LoginPage: React.FC<{
  oidc: OIDCRequestQuery;
  require: boolean;
}> = ({oidc, require}) => {
  const [service, setService] = React.useState<ServicePreview>();
  const [token, setToken] = React.useState<string | undefined>();

  const router = useRouter();
  const toast = useToast();
  const user = useRecoilValue(UserState);
  const setLoad = useSetRecoilState(LoadState);
  const {hasCopied, onCopy} = useClipboard(token || '');

  React.useEffect(() => {
    if (!require) {
      router.replace('/hello');
      return;
    }

    const f = async () => {
      try {
        setLoad(true);
        let from = document.referrer;
        if (from === '') {
          from = 'direct';
        }
        const service = await preview(oidc, from);
        setService(service);
        setLoad(false);
      } catch (error) {
        if (error instanceof Error) {
          toast({
            title: 'URLが正しくありません',
            description: error.message,
            status: 'error',
            isClosable: true,
            duration: 9000,
          });

          router.replace('/hello');
        }
      }
    };

    f();
  }, []);

  const submit = () => {
    const f = async () => {
      try {
        setLoad(true);
        let from = document.referrer;
        if (from === '') {
          from = 'direct';
        }
        const resp = await login(oidc, from);

        setLoad(false);
        if (oidc.redirectURL !== 'direct') {
          let url = `${oidc.redirectURL}?code=${resp.access_token}`;
          if (oidc.state !== '') {
            url += `&state=${oidc.state}`;
          }
          router.push(url);
        } else {
          setToken(resp.access_token);
        }
      } catch (error) {
        errorResp('interaction_required');
      }
    };

    f();
  };

  const cancel = () => {
    errorResp('consent_required');
  };

  const errorResp = (code: string) => {
    if (oidc.redirectURL !== 'direct') {
      router.push(`${oidc.redirectURL}?error=${code}`);
    } else {
      router.push('/');
    }
  };

  return (
    <Center>
      <Box
        width={{base: '95%', sm: '400px'}}
        height="600px"
        mt={{base: '0', sm: '3rem'}}
        borderRadius="20px"
        borderWidth={{base: '0', sm: '2px'}}
      >
        <Flex height="100%" alignItems="center">
          <Box width="100%">
            {!token ? (
              <>
                <Center mt="2rem" mb="1rem">
                  <Avatar src={user?.avatar_url} size="xl" />
                  <Text fontSize="1.5rem" fontWeight="bold" mx="1rem">
                    …
                  </Text>
                  <Avatar
                    name={service?.name}
                    src={service?.service_icon}
                    size="xl"
                  />
                </Center>
                <Heading textAlign="center">{service?.name}</Heading>
                <Text textAlign="center" mt=".5rem">
                  が、ログインを要求しています。
                </Text>
                <Center mt="2rem">
                  <Button
                    colorScheme="green"
                    w="95%"
                    size="md"
                    onClick={submit}
                  >
                    ログインする
                  </Button>
                </Center>
                <Center mt=".5rem" mb={{base: '3rem', sm: '5rem'}}>
                  <Button w="95%" size="md" onClick={cancel}>
                    キャンセルする
                  </Button>
                </Center>
              </>
            ) : (
              <>
                <Heading textAlign="center" fontSize="1.5rem">
                  アクセストークンをコピーしてログインしてください。
                </Heading>
                <Center mt="2rem">
                  <Textarea
                    placeholder="アクセストークン"
                    value={token}
                    width="95%"
                  />
                </Center>
                <Center mt=".5rem">
                  <Button w="95%" colorScheme="blue" size="md" onClick={onCopy}>
                    {hasCopied ? (
                      <IoCheckmarkOutline size="30px" />
                    ) : (
                      'コピーする'
                    )}
                  </Button>
                </Center>
              </>
            )}
          </Box>
        </Flex>
      </Box>
    </Center>
  );
};

export default LoginPage;
