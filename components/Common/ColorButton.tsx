import {IconButton, Tooltip, useColorMode} from '@chakra-ui/react';
import React from 'react';
import {TbMoon, TbSun} from 'react-icons/tb';

export const ColorButton = () => {
  const {colorMode, toggleColorMode} = useColorMode();

  if (colorMode === 'light') {
    return (
      <Tooltip
        label="ダークテーマに変更"
        hasArrow
        placement="bottom-end"
        borderRadius="7px"
      >
        <IconButton
          onClick={toggleColorMode}
          aria-label="to dark"
          icon={<TbSun size="25px" />}
          variant="ghost"
        />
      </Tooltip>
    );
  }

  return (
    <Tooltip
      label="ライトテーマに変更"
      hasArrow
      placement="bottom-end"
      borderRadius="7px"
    >
      <IconButton
        onClick={toggleColorMode}
        aria-label="to light"
        icon={<TbMoon size="25px" />}
        variant="ghost"
      />
    </Tooltip>
  );
};
