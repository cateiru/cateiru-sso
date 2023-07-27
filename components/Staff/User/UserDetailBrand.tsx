import {
  Button,
  Center,
  Link,
  Table,
  TableContainer,
  Tbody,
  Td,
  Text,
  Th,
  Thead,
  Tr,
} from '@chakra-ui/react';
import React from 'react';
import {TbPlugConnectedX} from 'react-icons/tb';
import {useSWRConfig} from 'swr';
import {UserBrand} from '../../../utils/types/staff';
import {DeleteButton} from '../../Common/DeleteButton';
import {Link as NextLink} from '../../Common/Next/Link';
import {useSecondaryColor} from '../../Common/useColor';
import {useRequest} from '../../Common/useRequest';
import {AddUser} from '../Brand/AddUser';

interface Props {
  userId: string;
  brands: UserBrand[];
}

export const UserDetailBrand: React.FC<Props> = props => {
  const textColor = useSecondaryColor();
  const {mutate} = useSWRConfig();
  const {request} = useRequest('/v2/admin/user/brand');

  const onDelete = async (brandId: string) => {
    const param = new URLSearchParams();
    param.append('user_id', props.userId);
    param.append('brand_id', brandId);

    const res = await request(
      {
        method: 'DELETE',
        mode: 'cors',
        credentials: 'include',
      },
      param
    );

    if (res) {
      purge();
    }
  };

  const purge = () => {
    mutate(
      key =>
        typeof key === 'string' &&
        key.startsWith(`/v2/admin/user_detail?user_id=${props.userId}`),
      undefined,
      {revalidate: true}
    );
  };

  return (
    <>
      <Text
        mt="2rem"
        fontSize="1.5rem"
        color={textColor}
        fontWeight="bold"
        mb=".5rem"
      >
        ブランド情報
      </Text>
      <Text color={textColor} mb="1rem">
        ブランドを使用することで、ユーザーに特殊な権限を与えたりグループ化することができます。
        <br />
        ユーザーにグループする権限を与えたりしたい場合は、orgを使用してください。
      </Text>
      <Link
        as={NextLink}
        href="/staff/brands"
        isExternal
        color={textColor}
        mb=".3rem"
        ml=".2rem"
      >
        ブランド一覧
      </Link>
      <AddUser userId={props.userId} handleSuccess={purge} />
      <TableContainer mt=".5rem">
        <Table variant="simple">
          <Thead>
            <Tr>
              <Th></Th>
              <Th>ブランド名</Th>
              <Th>加入日</Th>
              <Th></Th>
            </Tr>
          </Thead>
          <Tbody>
            {props.brands.map(brand => {
              return (
                <Tr key={`brand-${brand.id}`}>
                  <Td>
                    <Center>
                      <DeleteButton
                        onSubmit={async () => {
                          await onDelete(brand.brand_id);
                        }}
                        tooltipLabel="ユーザーからこのブランドを削除"
                        text={{
                          confirmHeader: `ユーザーと${brand.brand_name}の連携を解除しますか？`,
                          confirmOkTextColor: 'red',
                          confirmOkText: '削除',
                        }}
                        icon={<TbPlugConnectedX size="25px" />}
                      />
                    </Center>
                  </Td>
                  <Td>{brand.brand_name}</Td>
                  <Td>{new Date(brand.created_at).toLocaleString()}</Td>
                  <Td>
                    <Button
                      size="sm"
                      colorScheme="cateiru"
                      as={NextLink}
                      href={`/staff/brand/${brand.brand_id}`}
                    >
                      詳細
                    </Button>
                  </Td>
                </Tr>
              );
            })}
          </Tbody>
        </Table>
      </TableContainer>
    </>
  );
};
