import {
  Stack,
  Button,
  Modal,
  ModalOverlay,
  ModalContent,
  ModalHeader,
  ModalFooter,
  ModalBody,
  ModalCloseButton,
  useDisclosure,
  useToast,
  Text,
} from '@chakra-ui/react';
import {useRouter} from 'next/router';
import {useSetRecoilState} from 'recoil';
import {logout, deleteAccount} from '../../utils/api/logout';
import {UserState} from '../../utils/state/atom';

const LogoutDelete = () => {
  const deleteModal = useDisclosure();
  const logoutModal = useDisclosure();
  const toast = useToast();
  const setUser = useSetRecoilState(UserState);
  const router = useRouter();

  const logoutHandle = () => {
    const f = async () => {
      try {
        await logout();
        toast({
          title: 'ログインしました',
          status: 'info',
          isClosable: true,
        });
        setUser(null);
        router.replace('/');
      } catch (error) {
        if (error instanceof Error) {
          toast({
            title: error.message,
            status: 'error',
            isClosable: true,
            duration: 9000,
          });
        }
      }
    };
    f();
  };

  const deleteHandle = () => {
    const f = async () => {
      try {
        await deleteAccount();
        toast({
          title: 'アカウントを削除しました',
          status: 'info',
          isClosable: true,
        });
        setUser(null);
        router.replace('/');
      } catch (error) {
        if (error instanceof Error) {
          toast({
            title: error.message,
            status: 'error',
            isClosable: true,
            duration: 9000,
          });
        }
      }
    };
    f();
  };

  return (
    <>
      <Stack direction={['column', 'row']} spacing="1rem">
        <Button colorScheme="blue" onClick={logoutModal.onOpen}>
          ログアウト
        </Button>
        <Button variant="ghost" colorScheme="red" onClick={deleteModal.onOpen}>
          アカウント削除
        </Button>
      </Stack>
      <Modal
        isOpen={logoutModal.isOpen}
        onClose={logoutModal.onClose}
        isCentered
      >
        <ModalOverlay />
        <ModalContent>
          <ModalHeader>ログアウトしますか？</ModalHeader>
          <ModalCloseButton size="lg" />
          <ModalBody>再度ログインすることができます。</ModalBody>

          <ModalFooter>
            <Button colorScheme="red" mr={3} onClick={logoutHandle}>
              ログアウト
            </Button>
            <Button variant="ghost" onClick={logoutModal.onClose}>
              閉じる
            </Button>
          </ModalFooter>
        </ModalContent>
      </Modal>
      <Modal
        isOpen={deleteModal.isOpen}
        onClose={deleteModal.onClose}
        isCentered
      >
        <ModalOverlay />
        <ModalContent>
          <ModalHeader>アカウントを削除しますか？</ModalHeader>
          <ModalCloseButton size="lg" />
          <ModalBody>
            <Text>
              あなたの情報はすべて削除され、再度ログインすることはできなくなります。
            </Text>
            <Text color="red.500">*この操作は元には戻せません</Text>
          </ModalBody>

          <ModalFooter>
            <Button colorScheme="red" mr={3} onClick={deleteHandle}>
              アカウント削除
            </Button>
            <Button variant="ghost" onClick={deleteModal.onClose}>
              閉じる
            </Button>
          </ModalFooter>
        </ModalContent>
      </Modal>
    </>
  );
};

export default LogoutDelete;
