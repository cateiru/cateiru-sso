import {
  Modal,
  ModalOverlay,
  ModalContent,
  ModalHeader,
  ModalBody,
  ModalCloseButton,
  Button,
  Center,
  Divider,
  Text,
  SimpleGrid,
  useClipboard,
} from '@chakra-ui/react';
import React from 'react';
import {IoCopyOutline, IoCheckmarkSharp} from 'react-icons/io5';

interface Props {
  backups: string[];
  isOpen: boolean;
  onClose: () => void;
}

const ShowBackup: React.FC<Props> = ({backups, isOpen, onClose}) => {
  const backupCopy = useClipboard(backups.join(', '));

  return (
    <Modal isOpen={isOpen} onClose={onClose} isCentered>
      <ModalOverlay />
      <ModalContent>
        <ModalHeader>バックアップコード</ModalHeader>
        <ModalCloseButton size="lg" />
        <ModalBody>
          <Text color="red.500" fontWeight="bold">
            *必ず大切に保管してください
          </Text>
          <Text mt=".5rem">
            バックアップコードはワンタイムパスワードを忘れてしまった、削除されてしまった場合に入力できるコードです。
          </Text>
          <Text mt=".5rem">コードは1つにつき1回入力できます。</Text>
          <Divider my="1rem" />
          <SimpleGrid
            columns={2}
            spacing="10px"
            my="1rem"
            fontFamily="Source Code Pro"
          >
            {backups.map(v => (
              <Text key={v} textAlign="center">
                {v}
              </Text>
            ))}
          </SimpleGrid>
          <Center mb="1rem">
            <Button
              onClick={backupCopy.onCopy}
              leftIcon={
                backupCopy.hasCopied ? (
                  <IoCheckmarkSharp size="20px" color="#38A169" />
                ) : (
                  <IoCopyOutline size="20px" />
                )
              }
            >
              コピーする
            </Button>
          </Center>
        </ModalBody>
      </ModalContent>
    </Modal>
  );
};

export default ShowBackup;
