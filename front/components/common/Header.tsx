import {
  Box,
  Flex,
  IconButton,
  Spacer,
  useColorMode,
  Tooltip,
} from '@chakra-ui/react';
import NextLink from 'next/link';
import React from 'react';
import {MdOutlineDarkMode, MdOutlineLightMode} from 'react-icons/md';
import Logo from './Logo';

const Header: React.FC = () => {
  const {colorMode, toggleColorMode} = useColorMode();
  return (
    <Box width="100%">
      <Flex
        paddingLeft="1rem"
        paddingRight={{base: '1rem'}}
        paddingTop={{base: '1rem', sm: '0'}}
        marginY=".5rem"
      >
        <NextLink href="/">
          <Box width="10rem" cursor="pointer">
            <Logo />
          </Box>
        </NextLink>
        <Spacer />
        <Tooltip
          label={`${colorMode === 'light' ? 'ダーク' : 'ライト'}モードに変更`}
          hasArrow
          borderRadius="4px"
        >
          <IconButton
            aria-label="change color mode"
            icon={
              colorMode === 'light' ? (
                <MdOutlineDarkMode size="25px" />
              ) : (
                <MdOutlineLightMode size="25px" />
              )
            }
            variant="ghost"
            onClick={toggleColorMode}
          ></IconButton>
        </Tooltip>
      </Flex>
    </Box>
  );
};

export default Header;
