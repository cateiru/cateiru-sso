'use client';

import {
  Badge,
  Center,
  Divider,
  Heading,
  Link,
  Skeleton,
  Text,
  useColorModeValue,
} from '@chakra-ui/react';
import React from 'react';
import {TbExternalLink} from 'react-icons/tb';
import useSWR from 'swr';
import {badgeColor} from '../../utils/color';
import {orgDetailFeather} from '../../utils/swr/featcher';
import {ErrorType} from '../../utils/types/error';
import {PublicOrganizationDetail} from '../../utils/types/organization';
import {Avatar} from '../Common/Chakra/Avatar';
import {Error} from '../Common/Error/Error';
import {Margin} from '../Common/Margin';

interface Props {
  id: string;
}

export const OrganizationDetail: React.FC<Props> = ({id}) => {
  const textColor = useColorModeValue('gray.500', 'gray.400');
  const {data, error} = useSWR<PublicOrganizationDetail, ErrorType>(
    `/v2/org/detail?org_id=${id}`,
    () => orgDetailFeather(id)
  );

  return (
    <Margin>
      <Heading textAlign="center" mt="3rem">
        組織詳細
      </Heading>
      {error ? (
        <Center mt="1rem">
          <Error {...error} />
        </Center>
      ) : (
        <>
          <Center mt="1rem">
            <Avatar src={data?.image ?? ''} size="lg" />
            {data ? (
              <Text ml=".7rem" fontSize="1.5rem" fontWeight="bold">
                {data?.name}
              </Text>
            ) : (
              <Skeleton h="2rem" w="7rem" ml=".7rem" />
            )}
          </Center>
          {data ? (
            <>
              <Center mt="1rem">
                {data.link && (
                  <Text textAlign="center">
                    <Link href={data.link} isExternal>
                      {data.link}{' '}
                      <TbExternalLink
                        size="1rem"
                        style={{
                          display: 'inline-block',
                          verticalAlign: 'middle',
                          marginLeft: '0.1rem',
                        }}
                      />
                    </Link>
                  </Text>
                )}
                <Badge
                  verticalAlign="bottom"
                  ml="1rem"
                  colorScheme={badgeColor(data.role)}
                >
                  {data.role}
                </Badge>
              </Center>
            </>
          ) : (
            <Center mt="1rem">
              <Skeleton w="15rem" h="1.2rem" />
              <Skeleton ml="1rem" w="4rem" h="1.2rem" />
            </Center>
          )}
          {data ? (
            <Text mt=".5rem" color={textColor} textAlign="center">
              組織作成日: {new Date(data.created_at).toLocaleString()} - 加入日:{' '}
              {new Date(data.join_date).toLocaleString()}
            </Text>
          ) : (
            <Center mt=".7rem">
              <Skeleton w="25rem" h="1.2rem" />
            </Center>
          )}
          <Divider my="1rem" />
        </>
      )}
    </Margin>
  );
};
