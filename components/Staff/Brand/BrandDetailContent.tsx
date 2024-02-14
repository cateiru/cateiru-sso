import {
  Button,
  ButtonGroup,
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
import {Brand} from '../../../utils/types/staff';
import {Confirm} from '../../Common/Confirm/Confirm';
import {Link} from '../../Common/Next/Link';
import {useSecondaryColor} from '../../Common/useColor';
import {useRequest} from '../../Common/useRequest';

interface Props {
  brand?: Brand;
}

export const BrandDetailContent: React.FC<Props> = ({brand}) => {
  const router = useRouter();
  const deleteModal = useDisclosure();
  const textColor = useSecondaryColor();

  const {request} = useRequest('/admin/brand');

  const onDelete = async () => {
    const u = new URLSearchParams();
    u.append('brand_id', brand?.id ?? '');

    const res = await request(
      {
        method: 'DELETE',
        mode: 'cors',
        credentials: 'include',
      },
      u
    );

    if (res) {
      router.replace('/staff/brands');
    }
  };

  return (
    <>
      <TableContainer>
        <Table variant="simple">
          <Tbody>
            <Tr>
              <Td fontWeight="bold" color={textColor}>
                ID
              </Td>
              <Td>{brand?.id}</Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold" color={textColor}>
                ブランド名
              </Td>
              <Td>{brand?.name}</Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold" color={textColor}>
                説明
              </Td>
              <Td>
                <Text as="pre">{brand?.description}</Text>
              </Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold" color={textColor}>
                アカウント作成日
              </Td>
              <Td>
                {brand?.created_at
                  ? new Date(brand.created_at).toLocaleString()
                  : ''}
              </Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold" color={textColor}>
                アカウント更新日
              </Td>
              <Td>
                {brand?.updated_at
                  ? new Date(brand.updated_at).toLocaleString()
                  : ''}
              </Td>
            </Tr>
          </Tbody>
        </Table>
      </TableContainer>
      <ButtonGroup mt="1rem" w="100%">
        <Button
          w="100%"
          colorScheme="cateiru"
          as={Link}
          href={`/staff/brand/edit/${brand?.id}`}
        >
          編集
        </Button>
        <Button w="100%" onClick={deleteModal.onOpen}>
          削除
        </Button>
      </ButtonGroup>
      <Confirm
        onSubmit={onDelete}
        isOpen={deleteModal.isOpen}
        onClose={deleteModal.onClose}
        text={{
          confirmHeader: `「${brand?.name}」を削除しますか？`,
          confirmOkText: '削除',
          confirmOkTextColor: 'red',
        }}
      >
        <UnorderedList>
          <ListItem>この操作は元に戻せません。</ListItem>
          <ListItem>削除すると、全ユーザーの加入が解除されます。</ListItem>
        </UnorderedList>
      </Confirm>
    </>
  );
};
