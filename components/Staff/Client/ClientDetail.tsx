'use client';

import React from 'react';
import useSWR from 'swr';
import {staffClientDetailFeather} from '../../../utils/swr/staff';
import {ErrorType} from '../../../utils/types/error';
import {ClientDetail as ClientDetailType} from '../../../utils/types/staff';
import {Error} from '../../Common/Error/Error';
import {ClientDetailContent} from './ClientDetailContent';

interface Props {
  id: string;
}

export const ClientDetail: React.FC<Props> = ({id}) => {
  const {data, error} = useSWR<ClientDetailType, ErrorType>(
    `/v2/admin/client_detail?client_id=${id}`,
    () => staffClientDetailFeather(id)
  );

  if (error) {
    return <Error {...error} />;
  }

  return <ClientDetailContent data={data} />;
};
