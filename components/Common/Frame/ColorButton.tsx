import {IconButton, useColorMode} from '@chakra-ui/react';
import React from 'react';
import {TbMoon, TbSun} from 'react-icons/tb';
import {Tooltip} from '../Chakra/Tooltip';

export const ColorButton = () => {
  const {colorMode, toggleColorMode} = useColorMode();

  if (colorMode === 'light') {
    return (
      <Tooltip label="ダークテーマに変更" placement="bottom-end">
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
    <Tooltip label="ライトテーマに変更" placement="bottom-end">
      <IconButton
        onClick={toggleColorMode}
        aria-label="to light"
        icon={<TbMoon size="25px" />}
        variant="ghost"
      />
    </Tooltip>
  );
};
