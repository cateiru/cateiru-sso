'use client';

import {
  Badge,
  Link,
  Table,
  TableContainer,
  Tbody,
  Td,
  Text,
  Tr,
} from '@chakra-ui/react';
import React from 'react';
import {config} from '../../utils/config';
import {useRequest} from '../Common/useRequest';

export const DeployData = () => {
  const [apiConnectOk, setSpiConnectOk] = React.useState(false);
  const [apiMode, setApiMode] = React.useState('');

  const {request} = useRequest('/debug');

  React.useEffect(() => {
    const f = async () => {
      const res = await request();

      if (res) {
        const mode = (await res.json()).mode;
        setApiMode(mode);
        setSpiConnectOk(true);
      }
    };
    f();
  }, []);

  return (
    <TableContainer>
      <Table variant="simple">
        <Tbody>
          <Tr>
            <Td fontWeight="bold">モード</Td>
            <Td>{config.mode}</Td>
          </Tr>
          <Tr>
            <Td fontWeight="bold">APIモード</Td>
            <Td>{apiMode}</Td>
          </Tr>
          <Tr>
            <Td fontWeight="bold">リビジョン</Td>
            <Td>{config.revision}</Td>
          </Tr>
          <Tr>
            <Td fontWeight="bold">ブランチ名</Td>
            <Td>{config.branchName ?? ''}</Td>
          </Tr>
          <Tr>
            <Td fontWeight="bold">APIホスト</Td>
            <Td>
              <Link isExternal href={config.apiHost ?? config.apiPathPrefix}>
                {config.apiHost ?? config.apiPathPrefix}
              </Link>
              <Text ml="1rem" as="span">
                {apiConnectOk ? (
                  <Badge colorScheme="green" verticalAlign="top">
                    Connected
                  </Badge>
                ) : (
                  <Badge colorScheme="red" verticalAlign="top">
                    Not Connect
                  </Badge>
                )}
              </Text>
            </Td>
          </Tr>
          <Tr>
            <Td fontWeight="bold">CORS設定</Td>
            <Td>{config.apiCors ? '有効' : '無効'}</Td>
          </Tr>
          <Tr>
            <Td fontWeight="bold">タイトル</Td>
            <Td>{config.title}</Td>
          </Tr>
        </Tbody>
      </Table>
    </TableContainer>
  );
};
