'use client';

import {Box, Center} from '@chakra-ui/react';
import React from 'react';
import {useRecoilValue} from 'recoil';
import {UserState} from '../../utils/state/atom';
import {Error, OidcError} from '../Common/Error/Error';
import {Spinner} from '../Common/Icons/Spinner';
import {Consent} from './Consent';
import {useOidcRequire} from './useOidcRequire';
import {useOidcSubmit} from './useOidcSubmit';

export const OidcRequirePage = () => {
  const {submit, cancel} = useOidcSubmit();
  const {require, data, error, oidcError} = useOidcRequire(submit);

  const user = useRecoilValue(UserState);

  React.useEffect(() => {
    require();
  }, []);

  if (oidcError) {
    return <OidcError {...oidcError} />;
  }

  if (error) {
    return <Error {...error} />;
  }

  if (data === null) {
    return (
      <Center h="80vh">
        <Spinner size="xl" />
      </Center>
    );
  }

  return (
    <Box mt="3rem">
      <Consent
        userName={user?.user.user_name ?? ''}
        userImage={user?.user.avatar ?? undefined}
        data={data}
        onSubmit={submit}
        onCancel={cancel}
      />
    </Box>
  );
};
