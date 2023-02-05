import {
  Box,
  Flex,
  IconButton,
  Spacer,
  useColorMode,
  Tooltip,
  Center,
  Spinner,
} from '@chakra-ui/react';
import NextLink from 'next/link';
import Link from 'next/link';
import React from 'react';
import {TbMoon, TbSun} from 'react-icons/tb';
import {TbSettings, TbLayoutDashboard} from 'react-icons/tb';
import {useRecoilValue} from 'recoil';
import {UserState} from '../../utils/state/atom';
import Avatar from './Avatar';
import Logo from './Logo';

const Header = React.memo(() => {
  const {colorMode, toggleColorMode} = useColorMode();
  const user = useRecoilValue(UserState);

  const Setting = () => {
    return (
      <Center>
        <Tooltip label="設定" hasArrow borderRadius="4px">
          <Box>
            <Link href="/setting" passHref>
              <IconButton
                aria-label="change color mode"
                icon={<TbSettings size="25px" />}
                variant="ghost"
              ></IconButton>
            </Link>
          </Box>
        </Tooltip>
      </Center>
    );
  };

  const UserAvatar = () => {
    return (
      <NextLink href="/hello">
        <Center cursor="pointer">
          <Avatar size="sm" src={user?.avatar_url} ml=".5rem" />
        </Center>
      </NextLink>
    );
  };

  const Dashboard = () => {
    return (
      <Center>
        <Tooltip label="ダッシュボード" hasArrow borderRadius="4px">
          <Box>
            <Link href="/dashboard" passHref>
              <IconButton
                aria-label="change color mode"
                icon={<TbLayoutDashboard size="25px" />}
                variant="ghost"
              ></IconButton>
            </Link>
          </Box>
        </Tooltip>
      </Center>
    );
  };

  return (
    <Box width="100%">
      <Flex paddingX="1rem" marginY=".3rem" height="50px">
        <NextLink href="/">
          <Box width="160px" cursor="pointer">
            <Logo />
          </Box>
        </NextLink>
        <Spacer />
        <Center>
          <Tooltip
            label={`${colorMode === 'light' ? 'ダーク' : 'ライト'}モードに変更`}
            hasArrow
            borderRadius="4px"
          >
            <IconButton
              aria-label="change color mode"
              icon={
                colorMode === 'light' ? (
                  <TbMoon size="25px" />
                ) : (
                  <TbSun size="25px" />
                )
              }
              variant="ghost"
              onClick={toggleColorMode}
            ></IconButton>
          </Tooltip>
        </Center>
        {user !== undefined ? (
          user !== null ? (
            // ログインしている
            <>
              {user.role.includes('pro') && <Dashboard />}
              <Setting />
              <UserAvatar />
            </>
          ) : (
            // ログインしていない
            <></>
          )
        ) : (
          // 読込中
          <Center ml=".5rem">
            <Spinner thickness="2px" />
          </Center>
        )}
      </Flex>
    </Box>
  );
});

Header.displayName = 'Header';

export default Header;
