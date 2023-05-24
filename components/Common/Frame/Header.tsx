import {Button, Flex, MenuButton, Skeleton, Spacer} from '@chakra-ui/react';
import React from 'react';
import {useRecoilValue} from 'recoil';
import {UserState} from '../../../utils/state/atom';
import {Avatar} from '../Chakra/Avatar';
import {ColorButton} from './ColorButton';
import {HeaderTitle} from './HeaderTitle';
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
      pr=".8rem"
      pl=".8rem"
    >
      <HeaderTitle />
      <Spacer />
      <HeaderTools />
    </Flex>
  );
});

Header.displayName = 'Header';
