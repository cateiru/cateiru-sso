import {Box, Text, useColorModeValue} from '@chakra-ui/react';
import React from 'react';

interface Props {
  title: string;
  children: React.ReactNode;
}

export const SettingCard: React.FC<Props> = props => {
  const borderColor = useColorModeValue('gray.300', 'gray.600');

  return (
    <Box w="100%" margin="auto" my="2.5rem">
      <Text
        fontWeight="bold"
        fontSize="1.2rem"
        borderBottom="1px"
        pb=".5rem"
        pl=".5rem"
        borderColor={borderColor}
        mb="1rem"
      >
        {props.title}
      </Text>
      {props.children}
    </Box>
  );
};
