'use client';

import {Box, Center} from '@chakra-ui/react';
import React from 'react';
import {useRecoilValue} from 'recoil';
import {UserState} from '../../utils/state/atom';
import {Error, OidcError} from '../Common/Error/Error';
import {Spinner} from '../Common/Icons/Spinner';
import {Consent} from './Consent';
import {useOidcRequire} from './useOidcRequire';

export const OidcRequirePage = () => {
  const {require, data, error, oidcError} = useOidcRequire();

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
        userImage={user?.user.avatar ?? undefined}
        clientName={data.client_name}
        description={data.client_description ?? undefined}
        clientImage={data.image ?? undefined}
        registerUserName={data.register_user_name}
        registerUserImage={data.register_user_image ?? undefined}
        orgName={data.org_name ?? undefined}
        orgImage={data.org_image ?? undefined}
        orgMemberOnly={data.org_member_only}
        scopes={data.scopes}
        redirectUri={data.redirect_uri}
        // TODO: API側実装してから
        onSubmit={async () => {
          alert('submit');
        }}
        onCancel={async () => {
          alert('cancel');
        }}
      />
    </Box>
  );
};
