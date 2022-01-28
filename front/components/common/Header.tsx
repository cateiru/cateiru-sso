import {
  Box,
  Flex,
  IconButton,
  Spacer,
  useColorMode,
  Tooltip,
  Center,
  Avatar,
} from '@chakra-ui/react';
import NextLink from 'next/link';
import Link from 'next/link';
import React from 'react';
import {IoSettingsOutline} from 'react-icons/io5';
import {MdOutlineDarkMode, MdOutlineLightMode} from 'react-icons/md';
import {useRecoilValue} from 'recoil';
import {UserState} from '../../utils/state/atom';
import Logo from './Logo';

const Header = React.memo(() => {
  const {colorMode, toggleColorMode} = useColorMode();
  const user = useRecoilValue(UserState);

  const Setting = () => {
    return (
      <Center mx=".5rem">
        <Tooltip label="設定" hasArrow borderRadius="4px">
          <Box>
            <Link href="/setting" passHref>
              <IconButton
                aria-label="change color mode"
                icon={<IoSettingsOutline size="30px" />}
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
      <Center>
        <Avatar size="md" src={user?.avatar_url} />
      </Center>
    );
  };

  return (
    <Box width="100%">
      <Flex
        paddingLeft="1rem"
        paddingRight={{base: '1rem'}}
        paddingTop={{base: '1rem', sm: '0'}}
        marginY=".5rem"
        height="50px"
      >
        <NextLink href="/">
          <Center width="160px" cursor="pointer">
            <Logo />
          </Center>
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
                  <MdOutlineDarkMode size="30px" />
                ) : (
                  <MdOutlineLightMode size="30px" />
                )
              }
              variant="ghost"
              onClick={toggleColorMode}
            ></IconButton>
          </Tooltip>
        </Center>
        {user !== null && user !== undefined ? (
          <>
            <Setting />
            <UserAvatar />
          </>
        ) : (
          <></>
        )}
      </Flex>
    </Box>
  );
});

Header.displayName = 'Header';

export default Header;
