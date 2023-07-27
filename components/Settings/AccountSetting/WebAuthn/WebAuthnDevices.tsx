import {
  Skeleton,
  Table,
  TableContainer,
  Tbody,
  Td,
  Th,
  Thead,
  Tr,
} from '@chakra-ui/react';
import useSWR from 'swr';
import {webAuthnDevicesFeather} from '../../../../utils/swr/account';
import {AccountWebAuthnDevices} from '../../../../utils/types/account';
import {ErrorType} from '../../../../utils/types/error';
import {Error} from '../../../Common/Error/Error';
import {AgoTime} from '../../../Common/Time';
import {Device} from '../../../Histories/Device';
import {DeleteWebAuthn} from './DeleteWebauthn';

export const WebAuthnDevices = () => {
  const {data, error} = useSWR<AccountWebAuthnDevices, ErrorType>(
    '/v2/account/webauthn',
    webAuthnDevicesFeather
  );

  if (error) {
    return <Error {...error} />;
  }

  // 何も登録していないときにはなにも表示しない
  if (data && data.length === 0) {
    return null;
  }

  const Body = () => {
    if (!data) {
      return (
        <Tr>
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
    }

    return (
      <>
        {data.map(v => {
          return (
            <Tr key={`webauthn-${v.id}`}>
              <Td>
                <DeleteWebAuthn id={v.id} />
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
            </Tr>
          );
        })}
      </>
    );
  };

  return (
    <TableContainer mt="1rem">
      <Table variant="simple">
        <Thead>
          <Tr>
            <Th></Th>
            <Th>作成日時</Th>
            <Th textAlign="center">端末</Th>
          </Tr>
        </Thead>
        <Tbody>
          <Body />
        </Tbody>
      </Table>
    </TableContainer>
  );
};
