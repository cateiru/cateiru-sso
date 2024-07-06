import {
  Box,
  Button,
  Center,
  Flex,
  Skeleton,
  Table,
  TableContainer,
  Tbody,
  Td,
  Text,
  Th,
  Thead,
  Tr,
  useColorModeValue,
} from '@chakra-ui/react';
import React from 'react';
import type {ClientList as ClientListType} from '../../utils/types/client';
import {Avatar} from '../Common/Chakra/Avatar';
import {Link} from '../Common/Next/Link';
import {AgoTime} from '../Common/Time';
import {useSecondaryColor} from '../Common/useColor';

interface Props {
  clients?: ClientListType;
}

export const ClientListTable: React.FC<Props> = ({clients}) => {
  const shadow = useColorModeValue(
    '0px 0px 5px -2px #242838',
    '0px 0px 10px -2px #000000'
  );
  const textColor = useSecondaryColor();

  if (typeof clients === 'undefined') {
    return <></>;
  }

  return (
    <Box>
      {clients.map(v => {
        return (
          <Link
            href={`/client/${v.client_id}`}
            key={`client-list-${v.client_id}`}
          >
            <Box
              w="100%"
              minH="100px"
              boxShadow={shadow}
              borderRadius="10px"
              p="1rem"
              my="1rem"
            >
              <Flex>
                <Avatar src={v.image ?? ''} size="sm" />
                <Text ml=".5rem">
                  <Text fontSize="1.2rem" fontWeight="bold" as="span">
                    {v.name}
                  </Text>
                  <Text as="span" pl=".2rem" color={textColor}>
                    - {new Date(v.created_at).toLocaleString()}
                  </Text>
                </Text>
              </Flex>
              {v.description && (
                <Text as="pre" whiteSpace="pre-wrap">
                  {v.description}
                </Text>
              )}
            </Box>
          </Link>
        );
      })}
    </Box>
  );

  // return (
  //   <TableContainer>
  //     <Table variant="simple">
  //       <Thead>
  //         <Tr>
  //           <Th></Th>
  //           <Th>クライアント名</Th>
  //           <Th>説明</Th>
  //           <Th>作成日</Th>
  //           <Th></Th>
  //         </Tr>
  //       </Thead>
  //       <Tbody>
  //         {clients
  //           ? clients.map(v => {
  //               return (
  //                 <Tr key={`client-list-${v.client_id}`}>
  //                   <Td>
  //                     <Center>
  //                       <Avatar src={v.image ?? ''} size="sm" />
  //                     </Center>
  //                   </Td>
  //                   <Td>{v.name}</Td>
  //                   <Td
  //                     maxW="200px"
  //                     textOverflow="ellipsis"
  //                     whiteSpace="nowrap"
  //                     overflowX="hidden"
  //                   >
  //                     {v.description}
  //                   </Td>
  //                   <Td>
  //                     <AgoTime time={v.created_at} />
  //                   </Td>
  //                   <Td>
  //                     <Button
  //                       size="sm"
  //                       colorScheme="cateiru"
  //                       as={Link}
  //                       href={`/client/${v.client_id}`}
  //                     >
  //                       詳細
  //                     </Button>
  //                   </Td>
  //                 </Tr>
  //               );
  //             })
  //           : Array(5)
  //               .fill(0)
  //               .map((_, i) => {
  //                 return (
  //                   <Tr key={`loading-client-${i}`}>
  //                     <Td>
  //                       <Center>
  //                         <Skeleton w="32px" h="32px" borderRadius="50%" />
  //                       </Center>
  //                     </Td>
  //                     <Td>
  //                       <Skeleton w="100px" h="16px" />
  //                     </Td>
  //                     <Td>
  //                       <Skeleton w="200px" h="16px" />
  //                     </Td>
  //                     <Td>
  //                       <Skeleton w="100px" h="16px" />
  //                     </Td>
  //                     <Td>
  //                       <Skeleton w="52px" h="32px" borderRadius="0.375rem" />
  //                     </Td>
  //                   </Tr>
  //                 );
  //               })}
  //       </Tbody>
  //     </Table>
  //   </TableContainer>
  // );
};
