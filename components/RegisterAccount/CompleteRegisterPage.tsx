import {Heading, VStack, useColorModeValue} from '@chakra-ui/react';
import {CheckMark, CheckmarkProps} from './CheckMark';

export const CompleteRegisterPage = () => {
  const checkmarkProps = useColorModeValue<CheckmarkProps, CheckmarkProps>(
    {
      size: 100,
      bgColor: '#572bcf',
      color: '#fff',
    },
    {
      size: 100,
      bgColor: '#2bc4cf',
      color: '#fff',
    }
  );

  return (
    <VStack mt="3rem">
      <CheckMark {...checkmarkProps} />
      <Heading textAlign="center" mt=".5rem" color={checkmarkProps.bgColor}>
        アカウントを作成しました🎉
      </Heading>
    </VStack>
  );
};
