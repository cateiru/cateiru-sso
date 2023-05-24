'use client';

import {
  Box,
  Center,
  Divider,
  Heading,
  IconButton,
  Text,
  useColorModeValue,
} from '@chakra-ui/react';
import Link from 'next/link';
import {TbHistory, TbSettings} from 'react-icons/tb';
import {useRecoilValue} from 'recoil';
import {UserState} from '../../utils/state/atom';
import {Tooltip} from '../Common/Chakra/Tooltip';
import {Margin} from '../Common/Margin';
import {LoginDevice} from '../Histories/LoginDevice';
import {ProfileForm} from './ProfileForm';
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
      <Center mt="1rem">
        <Tooltip label="設定">
          <IconButton
            aria-label="設定"
            icon={<TbSettings size="25px" />}
            borderRadius="50%"
            as={Link}
            href="/settings"
            mr=".5rem"
          />
        </Tooltip>
        <Tooltip label="履歴">
          <IconButton
            aria-label="履歴"
            icon={<TbHistory size="25px" />}
            borderRadius="50%"
            as={Link}
            href="/histories"
          />
        </Tooltip>
      </Center>
      <ProfileForm />
      <Divider my="2rem" />
      <Box mb="3rem">
        <Text textAlign="center" mb="1rem" fontSize="1.2rem" fontWeight="bold">
          ログインしているデバイス
        </Text>
        <LoginDevice />
      </Box>
    </Margin>
  );
};
