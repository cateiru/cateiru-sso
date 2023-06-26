'use client';

import React from 'react';
import useSWR from 'swr';
import {staffUserDetailFeather} from '../../../utils/swr/featcher';
import {ErrorType} from '../../../utils/types/error';
import type {UserDetail as UserDetailType} from '../../../utils/types/staff';
import {Error} from '../../Common/Error/Error';
import {UserDetailContent} from './UserDetailContent';

interface Props {
  id: string;
}

export const UserDetail: React.FC<Props> = props => {
  const {data, error} = useSWR<UserDetailType, ErrorType>(
    `/v2/admin/user_detail?user_id=${props.id}`,
    () => staffUserDetailFeather(props.id)
  );

  if (error) {
    return <Error {...error} />;
  }

  console.log(data);

  return <UserDetailContent data={data} />;
};
