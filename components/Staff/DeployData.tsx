'use client';

import {Table, TableContainer, Tbody, Td, Tr} from '@chakra-ui/react';
import {config} from '../../utils/config';

export const DeployData = () => {
  return (
    <TableContainer>
      <Table variant="simple">
        <Tbody>
          <Tr>
            <Td fontWeight="bold">モード</Td>
            <Td>{config.mode}</Td>
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
            <Td>{config.apiHost}</Td>
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
