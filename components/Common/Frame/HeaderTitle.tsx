import {Flex, Text} from '@chakra-ui/react';
import {useAtomValue} from 'jotai';
import {config} from '../../../utils/config';
import {UserState} from '../../../utils/state/atom';
import {Logo} from '../Icons/Logo';
import {Link as NextLink} from '../Next/Link';

export const HeaderTitle = () => {
  const user = useAtomValue(UserState);

  return (
    <NextLink href={user ? '/profile' : '/'}>
      <Flex alignItems="center">
        <Logo size="40px" />
        <Text
          fontWeight="bold"
          ml=".5rem"
          fontSize={{base: '1.1rem', md: '1.3rem'}}
          background="linear-gradient(124deg, #2bc4cf, #572bcf, #cf2ba1)"
          backgroundClip="text"
        >
          {config.title}
        </Text>
      </Flex>
    </NextLink>
  );
};
