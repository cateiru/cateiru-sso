import {
  Button,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalHeader,
  ModalOverlay,
  Text,
  useToast,
} from '@chakra-ui/react';
import React from 'react';
import {FormProvider, useForm} from 'react-hook-form';
import {useSWRConfig} from 'swr';
import {ErrorUniqueMessage} from '../../../../utils/types/error';
import {
  PasswordForm,
  PasswordFormData,
} from '../../../Common/Form/PasswordForm';
import {useSecondaryColor} from '../../../Common/useColor';
import {useRequest} from '../../../Common/useRequest';

interface Props {
  isOpen: boolean;
  onClose: () => void;
}

export const OtpDeleteModal: React.FC<Props> = props => {
  const textColor = useSecondaryColor();
  const toast = useToast();
  const {mutate} = useSWRConfig();

  const {request: deleteOtp} = useRequest('/account/otp/delete', {
    customError: e => {
      if (e.unique_code === 8) {
        toast({
          title: 'パスワードが違います',
          status: 'error',
        });

        return;
      }

      const message = e.unique_code
        ? ErrorUniqueMessage[e.unique_code] ?? e.message
        : e.message;

      toast({
        title: message,
        status: 'error',
      });
    },
  });

  const methods = useForm<PasswordFormData>();
  const {
    handleSubmit,
    reset,
    formState: {isSubmitting},
  } = methods;

  const onDelete = async (data: PasswordFormData) => {
    const form = new FormData();
    form.append('password', data.password);

    const res = await deleteOtp({
      method: 'POST',
      body: form,
    });

    if (res) {
      toast({
        title: '二段階認証を無効化しました。',
        status: 'success',
      });

      // パージする
      mutate(
        key =>
          typeof key === 'string' && key.startsWith('/account/certificates'),
        undefined,
        {revalidate: true}
      );

      closeModal();
    }
  };

  const closeModal = () => {
    reset();
    props.onClose();
  };

  return (
    <Modal isOpen={props.isOpen} onClose={closeModal} isCentered size="lg">
      <ModalOverlay />
      <ModalContent>
        <ModalHeader>二段階認証を削除します</ModalHeader>
        <ModalCloseButton size="lg" />
        <ModalBody mb="1rem">
          <Text color={textColor}>
            二段階認証を削除するにはアカウントのパスワードを入力する必要があります。
          </Text>
          <FormProvider {...methods}>
            <form onSubmit={handleSubmit(onDelete)}>
              <PasswordForm enableWebauthn={false} />
              <Button
                mt="1rem"
                isLoading={isSubmitting}
                colorScheme="cateiru"
                type="submit"
                w="100%"
              >
                二段階認証を削除
              </Button>
            </form>
          </FormProvider>
        </ModalBody>
      </ModalContent>
    </Modal>
  );
};
