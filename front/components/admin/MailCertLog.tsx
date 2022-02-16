import {
  useToast,
  Table,
  Thead,
  Tbody,
  Tr,
  Th,
  Td,
  Center,
  Box,
  Heading,
  Tooltip,
} from '@chakra-ui/react';
import React from 'react';
import {useSetRecoilState} from 'recoil';
import {mailCertLog, MailCertLog} from '../../utils/api/admin';
import {formatDate, hawManyDaysAgo} from '../../utils/date';
import {LoadState} from '../../utils/state/atom';

const MailCertLog = () => {
  const toast = useToast();
  const [logs, setLogs] = React.useState<MailCertLog[]>([]);
  const setLoad = useSetRecoilState(LoadState);

  React.useEffect(() => {
    const f = async () => {
      setLoad(true);
      try {
        const logs = await mailCertLog();
        setLogs(logs);
      } catch (error) {
        if (error instanceof Error) {
          toast({
            title: error.message,
            status: 'error',
            isClosable: true,
            duration: 9000,
          });
        }
      }
      setLoad(false);
    };

    f();
  }, []);

  const LogElement = (log: MailCertLog) => {
    const date = new Date(log.try_date);

    return (
      <Tr key={log.log_id}>
        <Td textAlign="center">
          <Tooltip
            label={formatDate(date)}
            placement="top"
            borderRadius="5px"
            hasArrow
          >
            {hawManyDaysAgo(date)}
          </Tooltip>
        </Td>
        <Td textAlign="center">{log.target_mail}</Td>
        <Td textAlign="center">{log.ip}</Td>
      </Tr>
    );
  };

  return (
    <Center>
      <Box width={{base: '100%', lg: '1000px'}} mt="2rem">
        <Heading textAlign="center">メール認証ログ</Heading>
        <Box mx=".5rem" overflowX={{base: 'auto', lg: 'visible'}} mt="1rem">
          <Table
            variant="striped"
            minWidth="calc(1000px - 1rem)"
            size="lg"
            alignItems="center"
          >
            <Thead>
              <Tr>
                <Th textAlign="center">日時</Th>
                <Th textAlign="center">メールアドレス</Th>
                <Th textAlign="center">IP</Th>
              </Tr>
            </Thead>
            <Tbody>{logs.map(LogElement)}</Tbody>
          </Table>
        </Box>
      </Box>
    </Center>
  );
};

export default MailCertLog;
