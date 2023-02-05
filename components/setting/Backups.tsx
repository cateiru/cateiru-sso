import {
  Modal,
  ModalOverlay,
  ModalContent,
  ModalBody,
  ModalCloseButton,
  useDisclosure,
  Button,
  useToast,
  Input,
  InputGroup,
  InputRightElement,
  IconButton,
  Text,
  FormControl,
  FormErrorMessage,
} from '@chakra-ui/react';
import React from 'react';
import {useForm, SubmitHandler} from 'react-hook-form';
import {TbEye, TbEyeOff} from 'react-icons/tb';
import {getBackups} from '../../utils/api/otp';
import ShowBackup from './ShowBackup';

type InputsPassword = {
  password: string;
};

interface Props {
  onClose: () => void;
  isOpen: boolean;
}

const Backups: React.FC<Props> = ({onClose, isOpen}) => {
  const showBackupModal = useDisclosure();

  const [show, setShow] = React.useState(false);
  const [backups, setBackups] = React.useState<string[]>([]);
  const [load, setLoad] = React.useState(false);

  const {
    register,
    handleSubmit,
    formState: {errors},
    reset,
  } = useForm<InputsPassword>();

  const toast = useToast();

  // バックアップコードモーダルを開く
  const openBackups: SubmitHandler<InputsPassword> = data => {
    setLoad(true);
    const f = async () => {
      try {
        const backups = await getBackups(data.password);
        setBackups(backups);
        onClose();
        showBackupModal.onOpen();
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

      setShow(false);
      reset();
      setLoad(false);
    };

    f();
  };

  return (
    <>
      <Modal
        isOpen={isOpen}
        onClose={() => {
          setShow(false);
          reset();
          onClose();
        }}
        isCentered
      >
        <ModalOverlay />
        <ModalContent>
          <ModalCloseButton size="lg" />
          <ModalBody my="1.5rem" paddingX="1rem">
            <Text mt="1rem" mb=".5rem">
              バックアップコードを表示するにはパスワードが必要です。
            </Text>
            <form onSubmit={handleSubmit(openBackups)}>
              <FormControl isInvalid={Boolean(errors.password)}>
                <InputGroup>
                  <Input
                    id="password"
                    type={show ? 'text' : 'password'}
                    placeholder="パスワード"
                    {...register('password', {
                      required: '入力してください',
                    })}
                  />
                  <InputRightElement>
                    <IconButton
                      variant="ghost"
                      aria-label="show password"
                      icon={
                        show ? <TbEye size="25px" /> : <TbEyeOff size="25px" />
                      }
                      size="sm"
                      onClick={() => setShow(!show)}
                    />
                  </InputRightElement>
                </InputGroup>
                <FormErrorMessage>
                  {errors.password && errors.password.message}
                </FormErrorMessage>
              </FormControl>
              <Button
                marginTop="1rem"
                colorScheme="blue"
                isLoading={load}
                type="submit"
                width="100%"
              >
                表示する
              </Button>
            </form>
          </ModalBody>
        </ModalContent>
      </Modal>
      <ShowBackup
        backups={backups}
        isOpen={showBackupModal.isOpen}
        onClose={showBackupModal.onClose}
      />
    </>
  );
};

export default Backups;
