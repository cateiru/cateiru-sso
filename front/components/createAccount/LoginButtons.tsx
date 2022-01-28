import {Stack, Button} from '@chakra-ui/react';
import Link from 'next/link';

const LoginButtons = () => {
  return (
    <Stack
      direction={['column', 'row']}
      spacing="10px"
      width={{base: '100%', sm: 'auto'}}
    >
      <Button colorScheme="blue" variant="solid">
        ログイン
      </Button>
      <Link href="/create" passHref>
        <Button colorScheme="blue" variant="outline">
          新規登録
        </Button>
      </Link>
    </Stack>
  );
};

export default LoginButtons;
