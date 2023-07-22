'use client';

import {useParams} from 'next/navigation';
import useSWR from 'swr';
import {clientFetcher} from '../../utils/swr/client';
import type {ClientList as ClientListType} from '../../utils/types/client';
import {ErrorType} from '../../utils/types/error';
import {Error} from '../Common/Error/Error';
import {ClientListTable} from './ClientListTable';

export const MyClientList = () => {
  const param = useParams();

  const {data, error} = useSWR<ClientListType, ErrorType>(
    param.id ? `/v2/client/?org_id=${param.id}` : '/v2/client/',
    () => clientFetcher(undefined, param.id ?? '')
  );

  if (error) {
    return <Error {...error} />;
  }

  return <ClientListTable clients={data} />;
};
