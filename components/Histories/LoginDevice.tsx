import {
  Center,
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
import {TbCheck} from 'react-icons/tb';
import useSWR from 'swr';
import {loginDeviceFeather} from '../../utils/swr/history';
import {ErrorType} from '../../utils/types/error';
import {LoginDeviceList} from '../../utils/types/history';
import {Tooltip} from '../Common/Chakra/Tooltip';
import {Error} from '../Common/Error/Error';
import {AgoTime} from '../Common/Time';
import {Device} from './Device';
import {LogoutDevice} from './LogoutDevice';

export const LoginDevice = () => {
  const checkMarkColor = useColorModeValue('#68D391', '#38A169');
  const {data, error} = useSWR<LoginDeviceList, ErrorType>(
    '/history/login_devices',
    loginDeviceFeather
  );

  if (error) {
    return <Error {...error} />;
  }

  return (
    <TableContainer>
      <Table variant="simple">
        <Thead>
          <Tr>
            <Th></Th>
            <Th>ログイン日時</Th>
            <Th textAlign="center">端末</Th>
            <Th textAlign="center">IPアドレス</Th>
          </Tr>
        </Thead>
        <Tbody>
          {data
            ? data.map(v => {
                return (
                  <Tr key={v.id}>
                    <Td p="0">
                      {v.is_current ? (
                        <Tooltip label="このデバイス" placement="top">
                          <Center>
                            <TbCheck
                              size="25px"
                              color={checkMarkColor}
                              strokeWidth="3px"
                            />
                          </Center>
                        </Tooltip>
                      ) : (
                        <LogoutDevice loginHistoryId={v.id} />
                      )}
                    </Td>
                    <Td>
                      <AgoTime time={v.created_at} />
                    </Td>
                    <Td>
                      <Device
                        device={v.device}
                        os={v.os}
                        browser={v.browser}
                        isMobile={v.is_mobile}
                      />
                    </Td>
                    <Td textAlign="center">{v.ip}</Td>
                  </Tr>
                );
              })
            : Array(2)
                .fill(0)
                .map((_, i) => {
                  return (
                    <Tr key={`load-history-item-${i}`}>
                      <Td></Td>
                      <Td>
                        <Skeleton height="1rem" w="10rem" />
                      </Td>
                      <Td>
                        <Skeleton height="1rem" w="10rem" />
                      </Td>
                      <Td textAlign="center">
                        <Skeleton height="1rem" w="10rem" mx="auto" />
                      </Td>
                    </Tr>
                  );
                })}
        </Tbody>
      </Table>
    </TableContainer>
  );
};
