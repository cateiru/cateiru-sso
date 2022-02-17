import {Box, Flex, Heading, useColorMode} from '@chakra-ui/react';
import React from 'react';
import JSONPretty from 'react-json-pretty';
import {useSetRecoilState, useRecoilValue} from 'recoil';
import {useGetUserInfo} from '../../hooks/useGetUserInfo';
import {UserState, LoadState, NoLoginState} from '../../utils/state/atom';
import Spinner from '../common/Spinner';
import LoginButtons from '../createAccount/LoginButtons';

const ShowUser = () => {
  const user = useRecoilValue(UserState);
  const {colorMode} = useColorMode();
  const setLoad = useSetRecoilState(LoadState);
  const noLoad = useRecoilValue(NoLoginState);
  const get = useGetUserInfo();

  React.useEffect(() => {
    // ロードしながらこのページに飛んだときにロードを削除する
    if (noLoad) {
      setLoad(false);

      if (typeof user === 'undefined') {
        get();
      }
    }
  }, []);

  return (
    <Flex
      justifyContent="center"
      alignItems="center"
      flexDirection="column"
      width="100%"
      height="80vh"
      px={{base: '1rem', md: '5rem'}}
    >
      {user === undefined ? (
        <Spinner />
      ) : user === null ? (
        <>
          <Heading mb="1.4rem" textAlign="center">
            あれ、もしかしてログインしていませんか？
          </Heading>
          <LoginButtons />
        </>
      ) : (
        <>
          <Heading>こんにちは、{user?.user_name}</Heading>
          <Box
            fontSize={{base: '3vw', md: '1.3rem'}}
            fontWeight="600"
            mt="2rem"
            width={{base: '100%', lg: '800px'}}
          >
            <Box
              overflow="scroll"
              css={{
                '&::-webkit-scrollbar': {display: 'none'},
                scrollbarWidth: 'none',
              }}
            >
              <JSONPretty
                data={user}
                mainStyle={
                  colorMode === 'dark'
                    ? "line-height:1.3;color:#A0AEC0;font-family: 'Source Code Pro', monospace;"
                    : "line-height:1.3;color:#4A5568;font-family: 'Source Code Pro', monospace;"
                }
                errorStyle={
                  colorMode === 'dark'
                    ? 'line-height:1.3;color:#C53030;'
                    : 'line-height:1.3;color:#C53030;'
                }
                keyStyle={
                  colorMode === 'dark' ? 'color:#63B3ED;' : 'color:#4299E1;'
                }
                stringStyle={
                  colorMode === 'dark' ? 'color:#F6AD55;' : 'color:#DD6B20;'
                }
                valueStyle={
                  colorMode === 'dark' ? 'color:#68D391;' : 'color:#22543D;'
                }
                booleanStyle={
                  colorMode === 'dark' ? 'color:#68D391' : 'color:#22543D'
                }
              />
            </Box>
          </Box>
        </>
      )}
    </Flex>
  );
};

export default ShowUser;
