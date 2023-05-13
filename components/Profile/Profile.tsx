import {Center, Heading, Text, useColorModeValue} from '@chakra-ui/react';
import {useRecoilValue} from 'recoil';
import {UserState} from '../../utils/state/atom';
import {Margin} from '../Common/Margin';
import {UpdateAvatar} from './UpdateAvatar';

export const Profile = () => {
  const userNameColor = useColorModeValue('gray.500', 'gray.400');
  const user = useRecoilValue(UserState);

  return (
    <Margin>
      <Heading textAlign="center" mt="3rem">
        My プロフィール
      </Heading>
      <Center mt="3rem">
        <UpdateAvatar />
      </Center>
      {user?.user.family_name ||
      user?.user.middle_name ||
      user?.user.given_name ? (
        <>
          <Text
            textAlign="center"
            mt="1rem"
            fontSize="1.5rem"
            fontWeight="bold"
          >
            {user?.user.family_name}{' '}
            {user?.user.middle_name ? `${user?.user.middle_name} ` : ''}
            {user?.user.given_name}
          </Text>
          <Text
            textAlign="center"
            fontSize="1rem"
            fontWeight="bold"
            color={userNameColor}
          >
            &#064;{user?.user.user_name}
          </Text>
        </>
      ) : (
        <Text
          textAlign="center"
          mt="1rem"
          fontSize="1.5rem"
          fontWeight="bold"
          color={userNameColor}
        >
          &#064;{user?.user.user_name || '???'}
        </Text>
      )}
    </Margin>
  );
};
