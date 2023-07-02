import {
  Button,
  IconButton,
  Link,
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
import {TbXboxX} from 'react-icons/tb';
import {useSWRConfig} from 'swr';
import {UserBrand} from '../../../utils/types/staff';
import {Tooltip} from '../../Common/Chakra/Tooltip';
import {Link as NextLink} from '../../Common/Next/Link';
import {useRequest} from '../../Common/useRequest';
import {AddUser} from '../Brand/AddUser';

interface Props {
  userId: string;
  brands: UserBrand[];
}

export const UserDetailBrand: React.FC<Props> = props => {
  const textColor = useColorModeValue('gray.500', 'gray.400');
  const {mutate} = useSWRConfig();
  const {request} = useRequest('/v2/admin/user/brand');

  const onDelete = (brandId: string) => {
    const f = async () => {
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

    f();
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
        mb="1rem"
        fontSize="1.5rem"
        color={textColor}
        fontWeight="bold"
      >
        ブランド情報
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
                    <Tooltip label="ユーザーからこのブランドを削除">
                      <IconButton
                        icon={<TbXboxX size="20px" />}
                        variant="ghost"
                        aria-label="delete"
                        size="sm"
                        onClick={() => onDelete(brand.brand_id)}
                      />
                    </Tooltip>
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
