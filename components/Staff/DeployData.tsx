'use client';

import {
  Badge,
  Link,
  Skeleton,
  Table,
  TableContainer,
  Tbody,
  Td,
  Text,
  Tr,
} from '@chakra-ui/react';
import React from 'react';
import {config} from '../../utils/config';
import {
  DeployData as DeployDataType,
  DeployDataSchema,
} from '../../utils/types/staff';
import {useRequest} from '../Common/useRequest';

export const DeployData = () => {
  const [apiConnectOk, setSpiConnectOk] = React.useState(false);
  const [apiData, setApiData] = React.useState<
    DeployDataType | undefined | null
  >();

  const {request} = useRequest('/debug');

  React.useEffect(() => {
    const f = async () => {
      const res = await request();

      if (res) {
        const data = DeployDataSchema.safeParse(await res.json());
        if (data.success) {
          setApiData(data.data);
          setSpiConnectOk(true);
          return;
        }
      }

      setApiData(null);
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
            <Td>
              {apiData === null ? (
                <>No Connected</>
              ) : apiData ? (
                apiData.mode
              ) : (
                <Skeleton h="1.2rem" w="3rem" />
              )}
            </Td>
          </Tr>
          <Tr>
            <Td fontWeight="bold">クライアントIPアドレス</Td>
            <Td>
              {apiData === null ? (
                <>No Connected</>
              ) : apiData ? (
                apiData.ip_address
              ) : (
                <Skeleton h="1.2rem" w="3rem" />
              )}
            </Td>
          </Tr>
          <Tr>
            <Td fontWeight="bold">X-Forwarded-For</Td>
            <Td>
              {apiData === null ? (
                <>No Connected</>
              ) : apiData ? (
                apiData.xff
              ) : (
                <Skeleton h="1.2rem" w="3rem" />
              )}
            </Td>
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
                    No Connect
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
