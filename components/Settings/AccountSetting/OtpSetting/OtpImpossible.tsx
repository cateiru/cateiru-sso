import {Button, Text} from '@chakra-ui/react';
import {Link} from '../../../Common/Next/Link';
import {useSecondaryColor} from '../../../Common/useColor';

export const OtpImpossible = () => {
  const textColor = useSecondaryColor();

  return (
    <>
      <Text color={textColor}>
        パスワードが設定されていないため二段階認証は現在使用できません。
      </Text>
      <Button
        mt="1rem"
        w="100%"
        colorScheme="cateiru"
        as={Link}
        href="/settings/password"
      >
        パスワードを設定する
      </Button>
    </>
  );
};
