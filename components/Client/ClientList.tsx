'use client';

import {Box, IconButton} from '@chakra-ui/react';
import {useParams} from 'next/navigation';
import {TbPlus} from 'react-icons/tb';
import useSWR from 'swr';
import {clientFetcher} from '../../utils/swr/client';
import type {ClientListResponse} from '../../utils/types/client';
import {ErrorType} from '../../utils/types/error';
import {Tooltip} from '../Common/Chakra/Tooltip';
import {Error} from '../Common/Error/Error';
import {Link} from '../Common/Next/Link';
import {ClientListTable} from './ClientListTable';

export const ClientList = () => {
  const {id} = useParams();

  const {data, error} = useSWR<ClientListResponse, ErrorType>(
    id ? `/v2/client?org_id=${id}` : '/client',
    () => clientFetcher(undefined, id)
  );

  if (error) {
    return <Error {...error} />;
  }

  return (
    <>
      <Box
        position="fixed"
        bottom="5rem"
        right={{base: '1rem', md: '4rem'}}
        zIndex="100"
      >
        <Tooltip
          placement="top-end"
          label={
            data?.can_register_client
              ? `あと、${data?.remaining_creatable_quantity}件のクライアントが作成可能です`
              : 'クライアントの作成上限を超えています'
          }
        >
          <IconButton
            icon={<TbPlus size="30px" />}
            aria-label="クライアント新規追加"
            borderRadius="50%"
            size="lg"
            colorScheme="cateiru"
            isDisabled={!data?.can_register_client}
            as={Link}
            href={id ? `/client/register?org_id=${id}` : '/client/register'}
          />
        </Tooltip>
      </Box>
      <ClientListTable clients={data?.clients} />
    </>
  );
};
