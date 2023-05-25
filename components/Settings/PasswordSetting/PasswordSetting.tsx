'use client';

import {Skeleton} from '@chakra-ui/react';
import useSWR from 'swr';
import {userAccountCertificatesFeather} from '../../../utils/swr/featcher';
import {Error} from '../../Common/Error/Error';

export const PasswordSetting = () => {
  const {data, error} = useSWR(
    '/v2/account/certificates',
    userAccountCertificatesFeather
  );

  if (error) {
    return <Error {...error} />;
  }

  if (!data) {
    return <Skeleton w="100%" h="100px" />;
  }

  return <></>;
};
