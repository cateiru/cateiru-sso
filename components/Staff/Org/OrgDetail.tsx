'use client';

import React from 'react';
import useSWR from 'swr';
import {orgDetailFeather} from '../../../utils/swr/featcher';
import {ErrorType} from '../../../utils/types/error';
import {OrganizationDetail} from '../../../utils/types/staff';
import {Error} from '../../Common/Error/Error';
import {OrgDetailContent} from './OrgDetailContent';

interface Props {
  id: string;
}

export const OrgDetail: React.FC<Props> = ({id}) => {
  const {data, error} = useSWR<OrganizationDetail, ErrorType>(
    `/v2/admin/org?org_id=${id}`,
    () => orgDetailFeather(id)
  );

  if (error) {
    return <Error {...error} />;
  }

  if (!data) {
    return <></>;
  }

  return <OrgDetailContent {...data} />;
};
