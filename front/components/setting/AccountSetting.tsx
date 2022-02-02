import {Center, Box, Heading, Button} from '@chakra-ui/react';
import Link from 'next/link';
import LogoutDelete from './LogoutDelete';
import OTP from './OTP';

const AccountSetting = () => {
  return (
    <Center>
      <Box width={{base: '100%', lg: '800px'}} mx=".5rem" mt="2.5rem">
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
