import {
  Button,
  ButtonGroup,
  Center,
  Link,
  ListItem,
  Table,
  TableContainer,
  Tbody,
  Td,
  Text,
  Tr,
  UnorderedList,
  useDisclosure,
} from '@chakra-ui/react';
import {useRouter} from 'next/navigation';
import React from 'react';
import {OrganizationDetail} from '../../../utils/types/staff';
import {Avatar} from '../../Common/Chakra/Avatar';
import {Confirm} from '../../Common/Confirm/Confirm';
import {Link as NextLink} from '../../Common/Next/Link';
import {useSecondaryColor} from '../../Common/useColor';
import {useRequest} from '../../Common/useRequest';
import {OrgClient} from './OrgClient';
import {OrgUser} from './OrgUser';

export const OrgDetailContent: React.FC<OrganizationDetail> = data => {
  const textColor = useSecondaryColor();
  const deleteModal = useDisclosure();
  const router = useRouter();

  const {request} = useRequest('/v2/admin/org');

  const onDeleteOrg = async () => {
    const params = new URLSearchParams();
    params.append('org_id', data.org.id);

    const res = await request(
      {
        method: 'DELETE',
        mode: 'cors',
        credentials: 'include',
      },
      params
    );

    if (res) {
      router.replace('/staff/orgs');
    }
  };

  return (
    <>
      <Center mt="3rem">
        <Avatar src={data.org.image ?? ''} size="lg" />
      </Center>
      <Text
        mt="2rem"
        mb="1rem"
        fontSize="1.5rem"
        color={textColor}
        fontWeight="bold"
      >
        組織情報
      </Text>
      <TableContainer>
        <Table variant="simple">
          <Tbody>
            <Tr>
              <Td fontWeight="bold" color={textColor}>
                組織ID
              </Td>
              <Td>{data.org.id}</Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold" color={textColor}>
                組織名
              </Td>
              <Td>{data.org.name}</Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold" color={textColor}>
                組織URL
              </Td>
              <Td>
                {data.org.link ? (
                  <Link href={data.org.link} isExternal>
                    {data.org.link}
                  </Link>
                ) : (
                  ''
                )}
              </Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold" color={textColor}>
                組織作成日
              </Td>
              <Td>{new Date(data.org.created_at).toLocaleString()}</Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold" color={textColor}>
                組織更新日
              </Td>
              <Td>{new Date(data.org.updated_at).toLocaleString()}</Td>
            </Tr>
          </Tbody>
        </Table>
      </TableContainer>
      <ButtonGroup mt="1rem" w="100%">
        <Button
          w="100%"
          colorScheme="cateiru"
          as={NextLink}
          href={`/staff/org/edit/${data.org.id}`}
        >
          組織を編集
        </Button>
        <Button w="100%" onClick={deleteModal.onOpen}>
          組織削除
        </Button>
      </ButtonGroup>
      <Confirm
        onSubmit={onDeleteOrg}
        isOpen={deleteModal.isOpen}
        onClose={deleteModal.onClose}
        text={{
          confirmHeader: `「${data.org.name}」を削除しますか？`,
          confirmOkText: '削除',
          confirmOkTextColor: 'red',
        }}
      >
        <UnorderedList>
          <ListItem>この操作は元に戻せません。</ListItem>
          <ListItem>削除すると、全ユーザーの加入が解除されます。</ListItem>
        </UnorderedList>
      </Confirm>
      <OrgUser users={data.users} orgId={data.org.id} />
      <OrgClient clients={data.clients} />
    </>
  );
};
