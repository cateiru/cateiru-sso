import {Box, Flex, Text, Tooltip} from '@chakra-ui/react';
import React from 'react';
import {TbDeviceLaptop, TbDeviceMobile} from 'react-icons/tb';

interface Props {
  device: string | null;
  os: string | null;
  browser: string | null;
  isMobile: boolean | null;
}

export const Device: React.FC<Props> = props => {
  const DeviceIcon = () => {
    if (props.isMobile) {
      return (
        <Tooltip label="モバイル" hasArrow borderRadius="7px" placement="top">
          <Box>
            <TbDeviceMobile size="25px" />
          </Box>
        </Tooltip>
      );
    }
    return (
      <Tooltip label="デスクトップ" hasArrow borderRadius="7px" placement="top">
        <Box>
          <TbDeviceLaptop size="25px" />
        </Box>
      </Tooltip>
    );
  };

  return (
    <Flex>
      <DeviceIcon />
      <Text fontWeight="bold" ml=".5rem">
        {props.browser}（{props.device === '' ? props.os : props.device}）
      </Text>
    </Flex>
  );
};
