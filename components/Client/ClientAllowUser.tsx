'use client';

import {
  Button,
  Heading,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalHeader,
  ModalOverlay,
  useDisclosure,
} from '@chakra-ui/react';
import {useParams} from 'next/navigation';
import {Margin} from '../Common/Margin';
import {useRequest} from '../Common/useRequest';
import {ClientAllowUserTable} from './ClientAllowUserTable';

export const ClientAllowUser = () => {
  const {id} = useParams();

  const {isOpen, onOpen, onClose} = useDisclosure();
  const {request} = useRequest('/v2/client/allow_user');

  return (
    <Margin>
      <Heading textAlign="center" mt="3rem">
        許可ユーザーの編集
      </Heading>
      <Button w="100%" mt="1.5rem" colorScheme="cateiru" onClick={onOpen}>
        ルールを追加
      </Button>
      <Modal isOpen={isOpen} onClose={onClose} isCentered>
        <ModalOverlay />
        <ModalContent>
          <ModalHeader>ルールを追加</ModalHeader>
          <ModalCloseButton size="lg" />
          <ModalBody></ModalBody>
        </ModalContent>
      </Modal>
      <ClientAllowUserTable id={id} />
    </Margin>
  );
};
