import {Button, Center, Box, Heading, Stack} from '@chakra-ui/react';
import {logout, deleteAccount} from '../../utils/api/logout';
import OTP from './OTP';

const AccountSetting = () => {
  const logoutHandle = () => {
    const f = async () => {
      await logout();
    };
    f();
  };

  const deleteHandle = () => {
    const f = async () => {
      await deleteAccount();
    };
    f();
  };

  return (
    <Center>
      <Box width={{base: '100%', lg: '800px'}} mx=".5rem" mt="2.5rem">
        <Box my="1rem">
          <Heading size="1.8rem" mb=".7rem">
            二段階認証設定
          </Heading>
          <OTP />
        </Box>
        <Box my="2.5rem">
          <Heading size="1.8rem" mb=".7rem">
            ログアウト
          </Heading>
          <Stack direction={['column', 'row']} spacing="1rem">
            <Button colorScheme="blue">ログアウト</Button>
            <Button variant="ghost" colorScheme="red">
              アカウント削除
            </Button>
          </Stack>
        </Box>
      </Box>
    </Center>
  );
};

export default AccountSetting;
