'use client';

import React from 'react';
import useSWR from 'swr';
import {brandFeather} from '../../../utils/swr/featcher';
import {ErrorType} from '../../../utils/types/error';
import {Brand} from '../../../utils/types/staff';
import {Error} from '../../Common/Error/Error';
import {BrandDetailContent} from './BrandDetailContent';

interface Props {
  id: string;
}

export const BrandDetail: React.FC<Props> = ({id}) => {
  const {data, error} = useSWR<Brand, ErrorType>(
    `/v2/admin/brand?brand_id=${id}`,
    () => brandFeather(id)
  );

  if (error) {
    return <Error {...error} />;
  }

  return <BrandDetailContent brand={data} />;
};
