'use client';

import {Button} from '@chakra-ui/react';
import {useParams} from 'next/navigation';
import useSWR from 'swr';
import {clientFetcher} from '../../utils/swr/client';
import type {ClientListResponse} from '../../utils/types/client';
import {ErrorType} from '../../utils/types/error';
import {Tooltip} from '../Common/Chakra/Tooltip';
import {Error} from '../Common/Error/Error';
import {ClientListTable} from './ClientListTable';

export const MyClientList = () => {
  const param = useParams();

  const {data, error} = useSWR<ClientListResponse, ErrorType>(
    param.id ? `/v2/client?org_id=${param.id}` : '/v2/client',
    () => clientFetcher(undefined, param.id ?? '')
  );

  if (error) {
    return <Error {...error} />;
  }

  return (
    <>
      <Tooltip
        placement="top"
        label={
          data?.can_register_client
            ? `あと、${data?.remaining_creatable_quantity}件のクライアントが作成可能です`
            : 'クライアントの作成上限を超えています'
        }
      >
        <Button
          w="100%"
          my=".5rem"
          colorScheme="cateiru"
          isDisabled={!data?.can_register_client}
        >
          クライアントを新規作成
        </Button>
      </Tooltip>
      <ClientListTable clients={data?.clients} />
    </>
  );
};
