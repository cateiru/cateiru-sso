import {
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
import {Avatar} from '../Avatar';
import {Logo} from '../Icons/Logo';
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
              w="32px"
              h="32px"
              minW="32px"
              minH="32px"
              borderRadius="50%"
            >
              <Avatar src={user.user.avatar ?? ''} size="sm" />
            </MenuButton>
          </Menu>
        </>
      );
    }

    return (
      <>
        <Skeleton w="25px" h="25px" borderRadius="15%" mr="1rem" />
        <Skeleton w="32px" h="32px" borderRadius="50%" />
      </>
    );
  }, [user]);

  return (
    <Flex
      as="header"
      w="100%"
      h="60px"
      alignItems="center"
      pr=".5rem"
      pl=".5rem"
    >
      <NextLink href="/">
        <Flex alignItems="center">
          <Logo size={40} />
          <Text
            fontWeight="bold"
            ml=".5rem"
            fontSize="1.3rem"
            background="linear-gradient(124deg, #2bc4cf, #572bcf, #cf2ba1)"
            backgroundClip="text"
          >
            CateiruSSO
          </Text>
        </Flex>
      </NextLink>
      <Spacer />
      <HeaderTools />
    </Flex>
  );
});

Header.displayName = 'Header';
