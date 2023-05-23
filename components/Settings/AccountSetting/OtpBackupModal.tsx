import {
  Button,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalHeader,
  ModalOverlay,
  Text,
  useColorModeValue,
  useToast,
} from '@chakra-ui/react';
import React from 'react';
import {FormProvider, useForm} from 'react-hook-form';
import {z} from 'zod';
import {ErrorUniqueMessage} from '../../../utils/types/error';
import {PasswordForm, PasswordFormData} from '../../Common/Form/PasswordForm';
import {useRequest} from '../../Common/useRequest';
import {OtpBackups} from './OtpBackups';

interface Props {
  isOpen: boolean;
  onClose: () => void;
}

export const OtpBackupModal: React.FC<Props> = props => {
  const textColor = useColorModeValue('gray.500', 'gray.400');
  const toast = useToast();

  const [backups, setBackups] = React.useState<string[]>([]);

  const {request: backupsRequest} = useRequest('/v2/account/otp/backups', {
    customError: e => {
      if (e.unique_code === 8) {
        toast({
          title: 'パスワードが違います',
          status: 'error',
          duration: 5000,
          isClosable: true,
        });

        return;
      }

      const message = e.unique_code
        ? ErrorUniqueMessage[e.unique_code] ?? e.message
        : e.message;

      toast({
        title: message,
        status: 'error',
        duration: 5000,
        isClosable: true,
      });
    },
  });

  const onSubmit = async (data: PasswordFormData) => {
    const form = new FormData();
    form.append('password', data.password);

    const res = await backupsRequest({
      method: 'POST',
      mode: 'cors',
      credentials: 'include',
      body: form,
    });

    if (res) {
      const data = z.array(z.string()).safeParse(await res.json());
      if (data.success) {
        setBackups(data.data);
      } else {
        console.error(data.error);
      }
    }
  };

  const methods = useForm<PasswordFormData>();
  const {
    handleSubmit,
    reset,
    formState: {isSubmitting},
  } = methods;

  const closeModal = () => {
    setBackups([]);
    reset();
    props.onClose();
  };

  return (
    <Modal isOpen={props.isOpen} onClose={closeModal} isCentered size="lg">
      <ModalOverlay />
      <ModalContent>
        <ModalHeader>バックアップコード</ModalHeader>
        <ModalCloseButton size="lg" />
        <ModalBody mb="1rem">
          {backups.length === 0 ? (
            <>
              <Text color={textColor}>
                二段階認証を削除するにはアカウントのパスワードを入力する必要があります。
              </Text>
              <FormProvider {...methods}>
                <form onSubmit={handleSubmit(onSubmit)}>
                  <PasswordForm enableWebauthn={false} />
                  <Button
                    mt="1rem"
                    isLoading={isSubmitting}
                    colorScheme="cateiru"
                    type="submit"
                    w="100%"
                  >
                    バックアップコードを表示
                  </Button>
                </form>
              </FormProvider>
            </>
          ) : (
            <OtpBackups backups={backups} />
          )}
        </ModalBody>
      </ModalContent>
    </Modal>
  );
};
