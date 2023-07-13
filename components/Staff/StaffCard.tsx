import {Box, Text, useColorModeValue} from '@chakra-ui/react';
import React from 'react';

interface Props {
  title: string;
  description?: React.ReactNode | string;
  children: React.ReactNode;
}

export const StaffCard: React.FC<Props> = props => {
  const borderColor = useColorModeValue('gray.300', 'gray.600');
  const textColor = useColorModeValue('gray.500', 'gray.400');

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
      {props.description && (
        <Text color={textColor} mb="1rem">
          {props.description}
        </Text>
      )}
      {props.children}
    </Box>
  );
};
