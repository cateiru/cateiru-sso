import {Button, Spacer, Stack, Text} from '@chakra-ui/react';
import {useSecondaryColor} from '../../../Common/useColor';

export const OtpDisableText: React.FC<{onOpen: () => void}> = ({onOpen}) => {
  const textColor = useSecondaryColor();

  return (
    <>
      <Stack
        direction={{base: 'column', md: 'row'}}
        alignItems={{base: undefined, md: 'center'}}
      >
        <Text color={textColor}>二段階認証は現在設定されていません。</Text>
        <Spacer />
        <Button colorScheme="cateiru" onClick={onOpen}>
          二段階認証を設定する
        </Button>
      </Stack>
    </>
  );
};
