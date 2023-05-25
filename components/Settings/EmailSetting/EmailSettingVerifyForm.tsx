import {Center, Text, useColorModeValue} from '@chakra-ui/react';
import React from 'react';
import {TbArrowBigDown} from 'react-icons/tb';
import {config} from '../../../utils/config';
import {EmailVerifyForm} from '../../Common/Form/EmailVerifyForm';

interface Props {
  onSubmit: (data: EmailVerifyForm) => Promise<void>;
}

export const EmailSettingVerifyForm: React.FC<Props> = props => {
  const textColor = useColorModeValue('gray.500', 'gray.400');

  return (
    <>
      <Center mb=".5rem" mt="2rem">
        <TbArrowBigDown size="30px" />
      </Center>
      <Text textAlign="center" color={textColor} mb=".5rem">
        メールアドレスに送られた{config.emailCodeLength}
        桁の確認コードを入力してください
      </Text>
      <Center>
        <EmailVerifyForm
          onSubmit={props.onSubmit}
          emailCodeLength={config.emailCodeLength}
        />
      </Center>
    </>
  );
};
