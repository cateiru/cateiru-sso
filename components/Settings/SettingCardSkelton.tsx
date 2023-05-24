import {Box, Skeleton, SkeletonText, useColorModeValue} from '@chakra-ui/react';

export const SettingCardSkelton = () => {
  const borderColor = useColorModeValue('gray.300', 'gray.600');

  return (
    <Box w="100%" margin="auto" my="2.5rem">
      <Box
        borderBottom="1px"
        pb=".5rem"
        pl=".5rem"
        mb="1rem"
        borderColor={borderColor}
      >
        <Skeleton w="200px" fontSize="1.2rem">
          -
        </Skeleton>
      </Box>
      <SkeletonText />
    </Box>
  );
};
