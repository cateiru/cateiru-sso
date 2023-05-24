'use client';

import {
  Skeleton,
  Table,
  TableContainer,
  Tbody,
  Td,
  Th,
  Thead,
  Tr,
  useColorModeValue,
} from '@chakra-ui/react';
import useSWR from 'swr';
import {hawManyDaysAgo} from '../../utils/date';
import {loginTryHistoryFeather} from '../../utils/swr/featcher';
import {colorTheme} from '../../utils/theme';
import {ErrorType} from '../../utils/types/error';
import {
  LOGIN_TRY_IDENTIFIER,
  LoginTryHistoryList,
} from '../../utils/types/history';
import {Tooltip} from '../Common/Chakra/Tooltip';
import {Error} from '../Common/Error/Error';
import {Device} from './Device';

export const LoginTryTable = () => {
  const tableHeadBgColor = useColorModeValue(
    colorTheme.lightBackground,
    colorTheme.darkBackground
  );

  const {data, error} = useSWR<LoginTryHistoryList, ErrorType>(
    '/v2/history/try_login',
    loginTryHistoryFeather
  );

  if (error) {
    return <Error {...error} />;
  }

  return (
    <TableContainer>
      <Table variant="simple">
        <Thead>
          <Tr
            position={['sticky', '-webkit-sticky']}
            zIndex="0"
            top="0"
            bgColor={tableHeadBgColor}
          >
            <Th>ログイン日時</Th>
            <Th>種類</Th>
            <Th textAlign="center">IPアドレス</Th>
            <Th textAlign="center">端末</Th>
          </Tr>
        </Thead>
        <Tbody>
          {data
            ? data.map(v => {
                const created = new Date(v.created);
                return (
                  <Tr key={v.id}>
                    <Td>
                      <Tooltip placement="top" label={created.toLocaleString()}>
                        {hawManyDaysAgo(created)}
                      </Tooltip>
                    </Td>
                    <Td>{LOGIN_TRY_IDENTIFIER[v.identifier] ?? '不明'}</Td>
                    <Td textAlign="center">{v.ip}</Td>
                    <Td>
                      <Device
                        device={v.device}
                        os={v.os}
                        browser={v.browser}
                        isMobile={v.is_mobile}
                      />
                    </Td>
                  </Tr>
                );
              })
            : Array(5)
                .fill(0)
                .map((_, i) => {
                  return (
                    <Tr key={`load-history-item-${i}`}>
                      <Td>
                        <Skeleton height="1rem" w="10rem" />
                      </Td>
                      <Td>
                        <Skeleton height="1rem" w="10rem" />
                      </Td>
                      <Td textAlign="center">
                        <Skeleton height="1rem" w="10rem" mx="auto" />
                      </Td>
                      <Td>
                        <Skeleton height="1rem" w="10rem" />
                      </Td>
                    </Tr>
                  );
                })}
        </Tbody>
      </Table>
    </TableContainer>
  );
};
