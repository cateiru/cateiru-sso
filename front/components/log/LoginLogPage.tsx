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
  Tooltip,
  Stack,
  Spacer,
} from '@chakra-ui/react';
import type {ResponsiveValue} from '@chakra-ui/react';
import {Property} from 'csstype';
import Link from 'next/link';
import {useRouter} from 'next/router';
import React from 'react';
import {
  IoPhonePortraitOutline,
  IoDesktopOutline,
  IoArrowBackOutline,
  IoTabletPortraitOutline,
} from 'react-icons/io5';
import {useSetRecoilState} from 'recoil';
import useLoginLog from '../../hooks/useLoginLog';
import {LoginLogResponse} from '../../utils/api/log';
import {formatDate, hawManyDaysAgo} from '../../utils/date';
import {LoadState} from '../../utils/state/atom';
import UserAgent, {Device} from '../../utils/ua';

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

    let date: Date;
    if (v.is_logout) {
      date = new Date(v.date);
    } else {
      date = new Date(v.last_login_date);
    }

    return (
      <Tr key={v.access_id}>
        <Td>
          <Flex justifyContent="left" alignItems="center">
            <Box
              width="10px"
              height="10px"
              borderRadius="256px"
              backgroundColor={
                v.this_device
                  ? 'green.500'
                  : v.is_logout
                  ? 'red.500'
                  : 'yellow.500'
              }
              boxShadow={
                v.this_device
                  ? '0 0px 10px 0 #38A169'
                  : v.is_logout
                  ? '0 0px 10px 0 #E53E3E'
                  : '0 0px 10px 0 #DD6B20'
              }
              mr=".5rem"
            />
            <Tooltip
              label={formatDate(date)}
              placement="top"
              borderRadius="5px"
              hasArrow
            >
              {hawManyDaysAgo(date)}
            </Tooltip>
          </Flex>
        </Td>
        <Td textAlign="center">{v.ip_address}</Td>
        <Td>
          <Flex justifyContent="start">
            <Box>
              {(() => {
                switch (userAgent.device()) {
                  case Device.Mobile:
                    return (
                      <Tooltip
                        label="スマートフォン"
                        placement="top"
                        borderRadius="5px"
                        hasArrow
                      >
                        <Box>
                          <IoPhonePortraitOutline size="25px" />
                        </Box>
                      </Tooltip>
                    );
                  case Device.Desktop:
                    return (
                      <Tooltip
                        label="デスクトップ"
                        placement="top"
                        borderRadius="5px"
                        hasArrow
                      >
                        <Box>
                          <IoDesktopOutline size="25px" />
                        </Box>
                      </Tooltip>
                    );
                  case Device.Tablet:
                    return (
                      <Tooltip
                        label="タブレット"
                        placement="top"
                        borderRadius="5px"
                        hasArrow
                      >
                        <Box>
                          <IoTabletPortraitOutline size="25px" />
                        </Box>
                      </Tooltip>
                    );
                }
              })()}
            </Box>
            <Text ml=".5rem">{userAgent.uniqName()}</Text>
          </Flex>
        </Td>
      </Tr>
    );
  };

  const Header: React.FC<{align?: ResponsiveValue<Property.TextAlign>}> = ({
    children,
    align = 'center',
  }) => {
    return (
      <Th
        textAlign={align}
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
      <Box width={{base: '100%', lg: '1000px'}} mt="2rem">
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
        <Flex
          justifyContent={{base: 'left', md: 'right'}}
          mx="1rem"
          mt="1.5rem"
        >
          <Stack direction={{base: 'column', md: 'row'}} spacing="1rem">
            <Flex alignItems="center" justifyContent="left" mr="1rem">
              <Box
                width="10px"
                height="10px"
                borderRadius="256px"
                backgroundColor="green.500"
                boxShadow="0 0px 10px 0 #38A169"
                mr=".5rem"
              />
              <Text>このデバイス</Text>
            </Flex>
            <Flex alignItems="center" justifyContent="left" mr="1rem">
              <Box
                width="10px"
                height="10px"
                borderRadius="256px"
                backgroundColor="yellow.500"
                boxShadow="0 0px 10px 0 #E53E3E"
                mr=".5rem"
              />
              <Text>ログイン済み</Text>
            </Flex>
            <Flex alignItems="center" justifyContent="left" mr="1rem">
              <Box
                width="10px"
                height="10px"
                borderRadius="256px"
                backgroundColor="red.500"
                boxShadow="0 0px 10px 0 #DD6B20"
                mr=".5rem"
              />
              <Text>ログアウト済み</Text>
            </Flex>
          </Stack>
        </Flex>
        {/* TODO: overflow: auto属性が親~先祖についていると position: stickyが適用されない */}
        {/*        ref. https://github.com/w3c/csswg-drafts/issues/865 */}
        <Box mx=".5rem" overflowX={{base: 'auto', lg: 'visible'}} mt="1rem">
          <Table
            variant="striped"
            minWidth="calc(1000px - 1rem)"
            size="lg"
            alignItems="center"
          >
            <Thead>
              <Tr>
                <Header align="left">ログイン日時</Header>
                <Header>IPアドレス</Header>
                <Header>端末</Header>
              </Tr>
            </Thead>
            <Tbody>{log.map(v => element(v))}</Tbody>
          </Table>
        </Box>
      </Box>
    </Center>
  );
};

export default LoginLogPage;
