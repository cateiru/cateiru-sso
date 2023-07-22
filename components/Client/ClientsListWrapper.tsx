'use client';

import {
  Center,
  Heading,
  Link,
  Select,
  Skeleton,
  Text,
  useColorModeValue,
} from '@chakra-ui/react';
import {useParams, usePathname, useRouter} from 'next/navigation';
import React from 'react';
import {TbExternalLink} from 'react-icons/tb';
import useSWR from 'swr';
import {routeChangeStart} from '../../utils/event';
import {orgSimpleListFeather} from '../../utils/swr/organization';
import {ErrorType, ErrorUniqueMessage} from '../../utils/types/error';
import {SimpleOrganizationList} from '../../utils/types/organization';
import {Margin} from '../Common/Margin';
import {UserName} from '../Common/UserName';

interface Props {
  children: React.ReactNode;
}

export const ClientsListWrapper: React.FC<Props> = ({children}) => {
  const textColor = useColorModeValue('gray.500', 'gray.400');
  const pathname = usePathname();
  const router = useRouter();
  const params = useParams();

  const id: string | undefined = params?.id;

  const {data, error} = useSWR<SimpleOrganizationList, ErrorType>(
    id ? `/v2/org/list/simple?org_id=${id}` : '/v2/org/list',
    () => orgSimpleListFeather(id)
  );

  const title = React.useMemo(() => {
    if (pathname.match(/^\/clients\/org\/.+/)) {
      return '組織のクライアント一覧';
    }
    return 'クライアント一覧';
  }, [pathname]);

  const description = React.useMemo(() => {
    if (pathname.match(/^\/clients\/org\/.+/)) {
      if (error) {
        return ErrorUniqueMessage[error.unique_code ?? 0] ?? error.message;
      }

      return 'あなたの組織で作成されたクライアントの一覧が表示されます。';
    }

    return 'あなたの作成したクライアントの一覧が表示されます。';
  }, [pathname, error]);

  return (
    <Margin>
      <Heading textAlign="center" mt="3rem">
        {title}
      </Heading>
      <Text color={textColor} mt=".5rem" textAlign="center" mb=".5rem">
        {description}
      </Text>
      <Center mb="1rem">
        <UserName mb="0" />
        {error || !id ? (
          <></>
        ) : (
          <Link
            href={`/org/${id}`}
            ml="1rem"
            fontWeight="bold"
            color={textColor}
            isExternal
          >
            組織の詳細
            <TbExternalLink
              size="1rem"
              style={{
                display: 'inline-block',
                verticalAlign: 'middle',
                marginLeft: '0.1rem',
              }}
            />
          </Link>
        )}
      </Center>
      {error ? (
        <></>
      ) : data ? (
        <Select
          w={{base: '100%', md: '400px'}}
          mb="1rem"
          size="md"
          mx="auto"
          onChange={v => {
            routeChangeStart();
            router.replace(v.target.value);
          }}
          defaultValue={pathname}
        >
          <option value="/clients">クライアント一覧</option>
          {data.map(v => {
            return (
              <option value={`/clients/org/${v.id}`} key={v.id}>
                {v.name} のクライアント一覧
              </option>
            );
          })}
        </Select>
      ) : (
        <Center>
          <Skeleton h="40px" w="400px" borderRadius="7px" />
        </Center>
      )}
      {children}
    </Margin>
  );
};
