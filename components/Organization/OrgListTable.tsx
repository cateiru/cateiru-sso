import {
  Badge,
  Button,
  Center,
  Link,
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
import NextLink from 'next/link';
import {TbExternalLink} from 'react-icons/tb';
import useSWR from 'swr';
import {hawManyDaysAgo} from '../../utils/date';
import {orgListFeather} from '../../utils/swr/featcher';
import {colorTheme} from '../../utils/theme';
import {ErrorType} from '../../utils/types/error';
import {PublicOrganizationList} from '../../utils/types/organization';
import {Avatar} from '../Common/Chakra/Avatar';
import {Tooltip} from '../Common/Chakra/Tooltip';
import {Error} from '../Common/Error/Error';

export const OrgListTable = () => {
  const {data, error} = useSWR<PublicOrganizationList, ErrorType>(
    '/v2/org/list',
    orgListFeather
  );

  const tableHeadBgColor = useColorModeValue(
    colorTheme.lightBackground,
    colorTheme.darkBackground
  );

  if (error) {
    <Error {...error} />;
  }

  return (
    <TableContainer mt="2rem">
      <Table variant="simple">
        <Thead>
          <Tr
            position={['sticky', '-webkit-sticky']}
            zIndex="0"
            top="0"
            bgColor={tableHeadBgColor}
          >
            <Th></Th>
            <Th textAlign="center">組織名</Th>
            <Th>加入日</Th>
            <Th textAlign="center">ロール</Th>
            <Th textAlign="center">組織詳細</Th>
          </Tr>
        </Thead>
        <Tbody>
          {data
            ? data.map(v => {
                const joinDate = new Date(v.join_date);

                const badgeColor = () => {
                  if (v.role === 'owner') {
                    return 'red';
                  }

                  if (v.role === 'member') {
                    return 'blue';
                  }

                  return 'gray';
                };

                return (
                  <Tr key={v.id}>
                    <Td>
                      <Center>
                        <Avatar src={v.image ?? ''} size="sm" />
                      </Center>
                    </Td>
                    <Td>
                      {v.link ? (
                        <Link href={v.link} isExternal>
                          {v.name}{' '}
                          <TbExternalLink
                            size="1rem"
                            style={{
                              display: 'inline-block',
                              verticalAlign: 'middle',
                              marginLeft: '0.1rem',
                            }}
                          />
                        </Link>
                      ) : (
                        v.name
                      )}
                    </Td>
                    <Td>
                      <Tooltip
                        placement="top"
                        label={joinDate.toLocaleString()}
                      >
                        {hawManyDaysAgo(joinDate)}
                      </Tooltip>
                    </Td>
                    <Td textAlign="center">
                      <Badge colorScheme={badgeColor()}>{v.role}</Badge>
                    </Td>
                    <Td>
                      <Center>
                        <Button
                          size="sm"
                          colorScheme="cateiru"
                          as={NextLink}
                          href={`/org/${v.id}`}
                        >
                          詳細
                        </Button>
                      </Center>
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
                        <Skeleton h="32px" w="32px" borderRadius="50%" />
                      </Td>
                      <Td>
                        <Skeleton height="1rem" w="7rem" />
                      </Td>
                      <Td textAlign="center">
                        <Skeleton height="1rem" w="5rem" />
                      </Td>
                      <Td textAlign="center">
                        <Skeleton height="1rem" w="5rem" />
                      </Td>
                      <Td textAlign="center">
                        <Skeleton height="1rem" w="3rem" />
                      </Td>
                      <Td></Td>
                    </Tr>
                  );
                })}
        </Tbody>
      </Table>
    </TableContainer>
  );
};
