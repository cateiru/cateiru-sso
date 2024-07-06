import {
  Button,
  Flex,
  IconButton,
  MenuButton,
  Skeleton,
  Spacer,
} from '@chakra-ui/react';
import {useAtomValue} from 'jotai';
import React from 'react';
import {TbLogin} from 'react-icons/tb';
import {UserState} from '../../../utils/state/atom';
import {Avatar} from '../Chakra/Avatar';
import {Tooltip} from '../Chakra/Tooltip';
import {Link} from '../Next/Link';
import {ColorButton} from './ColorButton';
import {HeaderTitle} from './HeaderTitle';
import {Menu} from './Menu';

export const Header = React.memo(() => {
  const user = useAtomValue(UserState);

  const HeaderTools = React.useCallback(() => {
    if (user === null) {
      return (
        <>
          <ColorButton />
          <Tooltip label="ログイン" placement="bottom-end">
            <IconButton
              as={Link}
              href="/login"
              icon={<TbLogin size="25px" />}
              aria-label="login"
              variant="ghost"
              borderRadius="50%"
            />
          </Tooltip>
        </>
      );
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
        <Skeleton w="32px" h="32px" borderRadius="50%" mr="1rem" />
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
