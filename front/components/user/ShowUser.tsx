import {Box, Flex, Heading, Spinner, useColorMode} from '@chakra-ui/react';
import React from 'react';
import JSONPretty from 'react-json-pretty';
import {useGetUserInfo} from '../../hooks/useGetUserInfo';

const ShowUser = () => {
  const [get, user, err] = useGetUserInfo();
  const [load, setLoad] = React.useState(true);
  const {colorMode} = useColorMode();

  React.useEffect(() => {
    get();
    setLoad(false);
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
      {load ? (
        <Spinner
          mt="2rem"
          thickness="4px"
          speed="0.65s"
          color="blue.500"
          size="xl"
        />
      ) : err ? (
        <>
          <Heading>あれ、もしかしてログインしていませんか？</Heading>
        </>
      ) : (
        <>
          <Heading>こんにちは、{user?.user_name}</Heading>
          <Box
            fontSize={{base: '3vw', md: '1.3rem'}}
            fontWeight="600"
            mt="2rem"
          >
            <JSONPretty
              data={user}
              mainStyle={
                colorMode === 'dark'
                  ? "line-height:1.3;color:#A0AEC0;font-family: 'Source Code Pro', monospace;overflow:auto;"
                  : "line-height:1.3;color:#4A5568;font-family: 'Source Code Pro', monospace;overflow:auto;"
              }
              errorStyle={
                colorMode === 'dark'
                  ? 'line-height:1.3;color:#C53030;'
                  : 'line-height:1.3;color:#C53030;'
              }
              keyStyle={
                colorMode === 'dark' ? 'color:#63B3ED;' : 'color:#086F83;'
              }
              stringStyle={
                colorMode === 'dark' ? 'color:#F6AD55;' : 'color:#975A16;'
              }
              valueStyle={
                colorMode === 'dark' ? 'color:#68D391;' : 'color:#22543D;'
              }
              booleanStyle={
                colorMode === 'dark' ? 'color:#68D391' : 'color:#22543D'
              }
            />
          </Box>
        </>
      )}
    </Flex>
  );
};

export default ShowUser;
