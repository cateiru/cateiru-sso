import {
  Table,
  Thead,
  Tbody,
  Tr,
  Th,
  Td,
  Center,
  Box,
  Text,
  Button,
  Heading,
  Flex,
  useColorModeValue,
} from '@chakra-ui/react';
import Link from 'next/link';
import {useRouter} from 'next/router';
import React from 'react';
import {
  IoPhonePortraitOutline,
  IoDesktopOutline,
  IoArrowBackOutline,
} from 'react-icons/io5';
import {useSetRecoilState} from 'recoil';
import useLoginLog from '../../hooks/useLoginLog';
import {LoginLogResponse} from '../../utils/api/log';
import {formatDate} from '../../utils/date';
import {LoadState} from '../../utils/state/atom';
import UserAgent from '../../utils/ua';

const LoginLogPage = () => {
  const [log, load, getLog] = useLoginLog();
  const router = useRouter();
  const setLoad = useSetRecoilState(LoadState);
  const tableHeadBG = useColorModeValue('white', 'gray.800');

  React.useEffect(() => {
    if (!router.isReady) return;
    const query = router.query;

    if (typeof query['limit'] === 'string') {
      getLog(parseInt(query['limit']));
    } else {
      getLog(undefined);
    }
  }, [router.isReady, router.query]);

  React.useEffect(() => {
    setLoad(load);
  }, [load]);

  const element = (v: LoginLogResponse) => {
    const userAgent = new UserAgent(v.user_agent);

    return (
      <Tr key={v.access_id}>
        <Td textAlign="center">{formatDate(new Date(v.date))}</Td>
        <Td textAlign="center">{v.ip_address}</Td>
        <Td>
          <Flex justifyContent="start">
            <Box>
              {userAgent.isMobile() ? (
                <IoPhonePortraitOutline size="25px" />
              ) : (
                <IoDesktopOutline size="25px" />
              )}
            </Box>
            <Text ml=".5rem">{userAgent.uniqName()}</Text>
          </Flex>
        </Td>
      </Tr>
    );
  };

  const Header: React.FC = ({children}) => {
    return (
      <Th
        textAlign="center"
        position={['sticky', '-webkit-sticky']}
        zIndex="1"
        top="0"
        backgroundColor={tableHeadBG}
      >
        {children}
      </Th>
    );
  };

  return (
    <Center>
      <Box width={{base: '100%', lg: '1000px'}} mt="2rem" minHeight="50vh">
        <Box mx=".5rem">
          <Link href="/setting/account" passHref>
            <Button
              pl=".5rem"
              variant="ghost"
              leftIcon={<IoArrowBackOutline size="25px" />}
            >
              戻る
            </Button>
          </Link>
        </Box>
        <Heading textAlign="center">ログイン履歴</Heading>
        <Box overflow="auto" mx=".5rem" height="20vh" mt="2rem">
          <Table variant="striped" minWidth="800px" size="lg">
            <Thead>
              <Tr>
                <Header>ログイン日時</Header>
                <Header>IPアドレス</Header>
                <Header>端末</Header>
              </Tr>
            </Thead>
            <Tbody overflowY="scroll">{log.map(v => element(v))}</Tbody>
          </Table>
        </Box>
      </Box>
    </Center>
  );
};

export default LoginLogPage;
