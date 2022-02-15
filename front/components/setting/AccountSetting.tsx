import {Center, Box, Heading, Button} from '@chakra-ui/react';
import Link from 'next/link';
import {useRecoilValue} from 'recoil';
import {UserState} from '../../utils/state/atom';
import LogoutDelete from './LogoutDelete';
import OTP from './OTP';

const AccountSetting = () => {
  const user = useRecoilValue(UserState);

  return (
    <Center>
      <Box width={{base: '100%', lg: '800px'}} mx=".5rem" mt="2.5rem">
        {user?.role.includes('admin') && (
          <Box my="1rem">
            <Heading size="1.8rem" mb=".7rem">
              Adminページ
            </Heading>
            <Link href="/admin" passHref>
              <Button width={{base: '100%', sm: 'auto'}}>全ユーザー参照</Button>
            </Link>
          </Box>
        )}
        <Box my="1rem">
          <Heading size="1.8rem" mb=".7rem">
            SSOアカウント
          </Heading>
          <Link href="/setting/connected" passHref>
            <Button width={{base: '100%', sm: 'auto'}}>
              連携しているSSOアカウント確認
            </Button>
          </Link>
        </Box>
        <Box my="1rem">
          <Heading size="1.8rem" mb=".7rem">
            ログイン履歴確認
          </Heading>
          <Link href="/log/login" passHref>
            <Button width={{base: '100%', sm: 'auto'}}>ログイン履歴確認</Button>
          </Link>
        </Box>
        <Box my="2.5rem">
          <Heading size="1.8rem" mb=".7rem">
            二段階認証設定
          </Heading>
          <OTP />
        </Box>
        <Box my="2.5rem">
          <Heading size="1.8rem" mb=".7rem">
            ログアウト
          </Heading>
          <LogoutDelete />
        </Box>
      </Box>
    </Center>
  );
};

export default AccountSetting;
