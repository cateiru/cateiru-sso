import {
  Box,
  Button,
  Stack,
  Text,
  useColorModeValue,
  useDisclosure,
} from '@chakra-ui/react';
import React from 'react';
import {OtpBackupModal} from './OtpBackupModal';
import {OtpDeleteModal} from './OtpDeleteModal';

interface Props {
  updatedAt: Date;
}

export const OtpEnableText: React.FC<Props> = props => {
  const textColor = useColorModeValue('gray.500', 'gray.400');

  const deleteModal = useDisclosure();
  const backupModal = useDisclosure();

  return (
    <Box>
      <Text color={textColor}>二段階認証は設定されています。</Text>
      <Text color={textColor} mb=".5rem">
        設定日時:
        <Text as="span" fontWeight="bold" ml=".3rem">
          {props.updatedAt.toLocaleString()}
        </Text>
      </Text>
      <Stack direction={{base: 'column', md: 'row'}}>
        <Button w="100%" colorScheme="cateiru" onClick={backupModal.onOpen}>
          バックアップコードを表示する
        </Button>
        <Button w="100%" onClick={deleteModal.onOpen}>
          二段階認証を削除する
        </Button>
      </Stack>
      <OtpDeleteModal
        isOpen={deleteModal.isOpen}
        onClose={deleteModal.onClose}
      />
      <OtpBackupModal
        isOpen={backupModal.isOpen}
        onClose={backupModal.onClose}
      />
    </Box>
  );
};
