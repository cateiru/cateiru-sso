'use client';

import {
  Box,
  Center,
  Divider,
  Heading,
  IconButton,
  Text,
} from '@chakra-ui/react';
import {
  TbAddressBook,
  TbBuildingSkyscraper,
  TbHistory,
  TbSettings,
  TbTool,
} from 'react-icons/tb';
import {useRecoilValue} from 'recoil';
import {UserState} from '../../utils/state/atom';
import {Tooltip} from '../Common/Chakra/Tooltip';
import {Margin} from '../Common/Margin';
import {Link} from '../Common/Next/Link';
import {useSecondaryColor} from '../Common/useColor';
import {LoginDevice} from '../Histories/LoginDevice';
import {EmailView} from './EmailView';
import {ProfileDatetime} from './ProfileDatetime';
import {ProfileForm} from './ProfileForm';
import {UpdateAvatar} from './UpdateAvatar';

export const Profile = () => {
  const textColor = useSecondaryColor();
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
            color={textColor}
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
          color={textColor}
        >
          &#064;{user?.user.user_name || '???'}
        </Text>
      )}
      <Center mt=".5rem">
        <EmailView email={user?.user.email ?? ''} />
      </Center>
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
            mr=".5rem"
          />
        </Tooltip>
        <Tooltip label="クライアント">
          <IconButton
            aria-label="クライアント"
            icon={<TbAddressBook size="25px" />}
            borderRadius="50%"
            as={Link}
            href="/clients"
          />
        </Tooltip>
        {user?.joined_organization && (
          <Tooltip label="組織一覧">
            <IconButton
              aria-label="組織一覧"
              icon={<TbBuildingSkyscraper size="25px" />}
              borderRadius="50%"
              as={Link}
              href="/orgs"
              ml=".5rem"
            />
          </Tooltip>
        )}
        {user?.is_staff && (
          <Tooltip label="スッタフ管理画面">
            <IconButton
              aria-label="スッタフ管理画面"
              icon={<TbTool size="25px" />}
              colorScheme="cateiru"
              variant="outline"
              borderRadius="50%"
              as={Link}
              href="/staff"
              ml=".5rem"
            />
          </Tooltip>
        )}
      </Center>
      <ProfileForm />
      <Divider my="2rem" />
      <Box mb="3rem">
        <Text textAlign="center" mb="1rem" fontSize="1.2rem" fontWeight="bold">
          ログインしているデバイス
        </Text>
        <LoginDevice />
      </Box>
      <ProfileDatetime />
    </Margin>
  );
};
