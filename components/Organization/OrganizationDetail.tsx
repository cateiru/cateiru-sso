'use client';

import {Heading} from '@chakra-ui/react';
import React from 'react';
import useSWR from 'swr';
import {orgDetailFeather} from '../../utils/swr/featcher';
import {ErrorType} from '../../utils/types/error';
import {PublicOrganizationDetail} from '../../utils/types/organization';
import {Margin} from '../Common/Margin';

interface Props {
  id: string;
}

export const OrganizationDetail: React.FC<Props> = ({id}) => {
  const {data, error} = useSWR<PublicOrganizationDetail, ErrorType>(
    `/v2/admin/org?org_id=${id}`,
    () => orgDetailFeather(id)
  );

  return (
    <Margin>
      <Heading textAlign="center" mt="3rem">
        組織詳細
      </Heading>
    </Margin>
  );
};
