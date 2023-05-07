import {Avatar, Flex, Skeleton, Spacer} from '@chakra-ui/react';
import React from 'react';
import {useRecoilValue} from 'recoil';
import {UserState} from '../../utils/state/atom';
import {ColorButton} from './ColorButton';

export const Header = React.memo(() => {
  const user = useRecoilValue(UserState);

  const HeaderTools = React.useCallback(() => {
    if (user === null) {
      return <ColorButton />;
    }

    if (typeof user !== 'undefined') {
      return (
        <>
          <ColorButton />
          <Avatar src={user.user.avatar ?? ''} ml=".5rem" />
        </>
      );
    }

    return (
      <>
        <Skeleton w="40px" h="40px" borderRadius="15%" />
        <Skeleton w="48px" h="48px" borderRadius="50%" ml=".5rem" />
      </>
    );
  }, [user]);

  return (
    <Flex as="header" w="100%" h="60px" alignItems="center" pr=".5rem">
      <Spacer />
      <HeaderTools />
    </Flex>
  );
});

Header.displayName = 'Header';
