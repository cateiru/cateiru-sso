import {
  Avatar,
  Button,
  Flex,
  MenuButton,
  Skeleton,
  Spacer,
  Text,
} from '@chakra-ui/react';
import NextLink from 'next/link';
import React from 'react';
import {useRecoilValue} from 'recoil';
import {UserState} from '../../../utils/state/atom';
import {ColorButton} from './ColorButton';
import {Menu} from './Menu';

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
          <Menu>
            <MenuButton
              as={Button}
              variant="unstyled"
              ml=".5rem"
              h="48px"
              borderRadius="50%"
            >
              <Avatar src={user.user.avatar ?? ''} />
            </MenuButton>
          </Menu>
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
      <NextLink href="/">
        <Text pl="1rem" fontSize="1.5rem" fontWeight="bold">
          たいとる
        </Text>
      </NextLink>
      <Spacer />
      <HeaderTools />
    </Flex>
  );
});

Header.displayName = 'Header';
