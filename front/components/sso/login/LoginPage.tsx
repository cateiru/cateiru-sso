import {
  Box,
  Center,
  Heading,
  Text,
  Button,
  Flex,
  Textarea,
  useClipboard,
  Link,
} from '@chakra-ui/react';
import {useRouter} from 'next/router';
import React from 'react';
import {IoCheckmarkOutline} from 'react-icons/io5';
import {useRecoilValue} from 'recoil';
import {login, preview, ServicePreview} from '../../../utils/api/loginSSO';
import {OIDCRequestQuery} from '../../../utils/sso/login';
import {UserState} from '../../../utils/state/atom';
import Avatar from '../../common/Avatar';
import Spinner from '../../common/Spinner';

enum LoginState {
  Loading,
  Success,
  Failed,
}

const LoginPage: React.FC<{
  oidc: OIDCRequestQuery;
  require: boolean;
}> = ({oidc, require}) => {
  const [service, setService] = React.useState<ServicePreview>();
  const [token, setToken] = React.useState<string | undefined>();
  const [status, setStatus] = React.useState(LoginState.Loading);
  const [load, setLoad] = React.useState(false);
  const [cancelLoad, setCancelLoad] = React.useState(false);
  const [errorMessage, setErrorMessage] = React.useState('値が不正です');

  const router = useRouter();
  const user = useRecoilValue(UserState);
  const {hasCopied, onCopy} = useClipboard(token || '');

  React.useEffect(() => {
    if (!require) {
      setStatus(LoginState.Failed);
      return;
    }

    const f = async () => {
      try {
        let from = document.referrer;
        if (from === '') {
          from = 'direct';
        }
        const service = await preview(oidc, from);
        setService(service);
        setStatus(LoginState.Success);
      } catch (error) {
        if (error instanceof Error) {
          setStatus(LoginState.Failed);
          setErrorMessage(error.message);
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
    setCancelLoad(true);
    errorResp('consent_required');
  };

  const errorResp = (code: string) => {
    if (oidc.redirectURL !== 'direct') {
      router.push(`${oidc.redirectURL}?error=${code}`);
    } else {
      router.push('/');
    }
  };

  switch (status) {
    case LoginState.Loading:
      return (
        <Center>
          <Box
            width={{base: '95%', sm: '400px'}}
            height="600px"
            mt={{base: '0', sm: '3rem'}}
            borderRadius="20px"
            borderWidth={{base: '0', sm: '2px'}}
          >
            <Center height="100%">
              <Spinner />
            </Center>
          </Box>
        </Center>
      );
    case LoginState.Success:
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
                      <Avatar src={user?.avatar_url} size="xl" isShadow />
                      <Text fontSize="1.5rem" fontWeight="bold" mx="1rem">
                        …
                      </Text>
                      <Avatar
                        name={service?.name}
                        src={service?.service_icon}
                        size="xl"
                        isShadow
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
                        isLoading={load}
                        disabled={cancelLoad}
                      >
                        許可する
                      </Button>
                    </Center>
                    <Center mt=".5rem">
                      <Button
                        w="95%"
                        size="md"
                        onClick={cancel}
                        disabled={load || cancelLoad}
                      >
                        キャンセルする
                      </Button>
                    </Center>
                    <Center mt="1rem">
                      {oidc.redirectURL === 'direct' ? (
                        <Text fontSize=".8rem">
                          リダイレクトせずトークンが表示されます
                        </Text>
                      ) : (
                        <Text fontSize=".8rem" textAlign="center">
                          <Link href={oidc.redirectURL} isExternal>
                            {oidc.redirectURL}
                          </Link>
                          <Text as="span" whiteSpace="nowrap">
                            へリダイレクトします
                          </Text>
                        </Text>
                      )}
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
                      <Button
                        w="95%"
                        colorScheme="blue"
                        size="md"
                        onClick={onCopy}
                      >
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
    case LoginState.Failed:
      return (
        <Center>
          <Box
            width={{base: '95%', sm: '400px'}}
            height="600px"
            mt={{base: '0', sm: '3rem'}}
            borderRadius="20px"
            borderWidth={{base: '0', sm: '2px'}}
          >
            <Flex height="100%" alignItems="center" color="red.600">
              <Box width="100%">
                <Heading textAlign="center">エラー</Heading>
                <Text textAlign="center">{errorMessage}</Text>
              </Box>
            </Flex>
          </Box>
        </Center>
      );
  }
};

export default LoginPage;
