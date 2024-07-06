import {Button, Center, Heading, Text} from '@chakra-ui/react';
import React from 'react';
import {TbDeviceMobile} from 'react-icons/tb';
import {useSecondaryColor} from '../../../Common/useColor';

interface Props {
  onRegisterStart: () => Promise<void>;
}

export const OtpRegisterStart: React.FC<Props> = props => {
  const textColor = useSecondaryColor();
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
