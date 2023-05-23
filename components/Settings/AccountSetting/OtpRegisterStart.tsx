import {
  Button,
  Center,
  Heading,
  Text,
  useColorModeValue,
} from '@chakra-ui/react';
import React from 'react';
import {TbDeviceMobile} from 'react-icons/tb';

interface Props {
  onRegisterStart: () => Promise<void>;
}

export const OtpRegisterStart: React.FC<Props> = props => {
  const textColor = useColorModeValue('gray.500', 'gray.400');
  const [loading, setLoading] = React.useState(false);

  const onClick = () => {
    setLoading(true);
    props
      .onRegisterStart()
      .then(() => setLoading(false))
      .catch(() => setLoading(false));
  };

  return (
    <>
      <Center mb=".5rem">
        <TbDeviceMobile size="60px" />
      </Center>
      <Heading fontSize="1.3rem" textAlign="center" mb="1rem">
        二段階認証を設定します
      </Heading>
      <Text color={textColor} mb="1rem">
        Authenticatorアプリを準備してください。
      </Text>
      <Button
        onClick={onClick}
        colorScheme="cateiru"
        w="100%"
        isLoading={loading}
      >
        二段階認証を設定する
      </Button>
    </>
  );
};
