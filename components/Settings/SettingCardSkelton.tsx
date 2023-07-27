import {Box, Skeleton, SkeletonText} from '@chakra-ui/react';
import {useBorderColor} from '../Common/useColor';

export const SettingCardSkelton = () => {
  const borderColor = useBorderColor();

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
