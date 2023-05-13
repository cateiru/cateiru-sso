import {
  Menu as ChakraMenu,
  MenuDivider,
  MenuGroup,
  MenuItem,
  MenuList,
} from '@chakra-ui/react';
import Link from 'next/link';
import React from 'react';
import {TbLogout, TbSettings, TbUser, TbUsers} from 'react-icons/tb';
import {useLogout} from '../useLogout';

const fontSize = {base: '1.5rem', sm: '1rem'};
const height = {base: '55px', sm: '32px'};

export const Menu: React.FC<{children: React.ReactNode}> = ({children}) => {
  const {logout} = useLogout();

  return (
    <ChakraMenu>
      {children}
      <MenuList>
        <MenuGroup title="プロフィール" fontSize={fontSize}>
          <MenuItem
            as={Link}
            href="/profile"
            icon={<TbUser size="20px" />}
            fontSize={fontSize}
            h={height}
          >
            My プロフィール
          </MenuItem>
          <MenuItem
            as={Link}
            href="/settings"
            icon={<TbSettings size="20px" />}
            fontSize={fontSize}
            h={height}
          >
            設定
          </MenuItem>
        </MenuGroup>
        <MenuDivider />
        <MenuGroup title="アカウント設定" fontSize={fontSize}>
          <MenuItem
            as={Link}
            href="/switch_account"
            icon={<TbUsers size="20px" />}
            fontSize={fontSize}
            h={height}
          >
            アカウントを切り替える
          </MenuItem>
          <MenuItem
            icon={<TbLogout size="20px" />}
            onClick={logout}
            fontSize={fontSize}
            h={height}
          >
            ログアウト
          </MenuItem>
        </MenuGroup>
      </MenuList>
    </ChakraMenu>
  );
};
