import {
  Button,
  FormControl,
  FormErrorMessage,
  FormLabel,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalHeader,
  ModalOverlay,
  Table,
  TableContainer,
  Tbody,
  Td,
  Text,
  Textarea,
  Tr,
  useDisclosure,
} from '@chakra-ui/react';
import {useAtomValue} from 'jotai';
import React from 'react';
import {useForm} from 'react-hook-form';
import {useSWRConfig} from 'swr';
import {UserState} from '../../../utils/state/atom';
import {Staff} from '../../../utils/types/staff';
import {useSecondaryColor} from '../../Common/useColor';
import {useRequest} from '../../Common/useRequest';

interface FormTypes {
  memo?: string;
}

interface Props {
  staff?: Staff | null;
  userId?: string;
}

export const UserDetailStaff = React.memo<Props>(({staff, userId}) => {
  const textColor = useSecondaryColor();

  const {request: staffRequest} = useRequest('/admin/staff');
  const {mutate} = useSWRConfig();
  const user = useAtomValue(UserState);

  const {isOpen, onOpen, onClose} = useDisclosure();

  const {
    register,
    reset,
    formState: {errors, isSubmitting},
    handleSubmit,
  } = useForm<FormTypes>();

  const onToggleStaff = async (isStaff: boolean, memo?: string) => {
    if (!userId) return;

    const form = new FormData();

    form.append('user_id', userId);
    form.append('memo', memo ?? '');
    form.append('is_staff', isStaff ? 'true' : 'false');

    const res = await staffRequest({
      method: 'POST',
      body: form,
      mode: 'cors',
      credentials: 'include',
    });

    if (res) {
      // キャッシュを飛ばして再ロードする
      mutate(
        key =>
          typeof key === 'string' &&
          key.startsWith(`/v2/admin/user_detail?user_id=${userId}`),
        undefined,
        {revalidate: true}
      );
    }
  };

  const onSubmitForm = async (data: FormTypes) => {
    await onToggleStaff(true, data.memo);
    onClose();
    reset();
  };

  return (
    <>
      <Text
        mt="2rem"
        mb="1rem"
        fontSize="1.5rem"
        color={textColor}
        fontWeight="bold"
      >
        スタッフ情報
      </Text>
      <Text color={textColor} mb="1rem">
        ユーザーをスタッフにすることで、ユーザーは管理画面にアクセスすることができます。
      </Text>
      <TableContainer>
        <Table variant="simple">
          <Tbody>
            <Tr>
              <Td fontWeight="bold" color={textColor}>
                スタッフ
              </Td>
              <Td>{staff ? 'はい' : 'いいえ'}</Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold" color={textColor}>
                メモ
              </Td>
              <Td>{staff?.memo}</Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold" color={textColor}>
                作成日
              </Td>
              <Td>
                {staff?.created_at
                  ? new Date(staff?.created_at).toLocaleString()
                  : ''}
              </Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold" color={textColor}>
                更新日
              </Td>
              <Td>
                {staff?.updated_at
                  ? new Date(staff?.updated_at).toLocaleString()
                  : ''}
              </Td>
            </Tr>
          </Tbody>
        </Table>
      </TableContainer>

      <Button w="100%" colorScheme="cateiru" onClick={onOpen}>
        {staff ? 'スタッフを更新する' : 'スタッフにする'}
      </Button>
      {user?.user.id !== userId && staff && (
        <Button
          mt=".5rem"
          w="100%"
          onClick={() => (staff ? onToggleStaff(false) : undefined)}
        >
          スタッフから外す
        </Button>
      )}
      <Modal
        isOpen={isOpen}
        onClose={() => {
          onClose();
          reset();
        }}
        isCentered
      >
        <ModalOverlay />
        <ModalContent>
          <ModalHeader>
            {staff ? 'スタッフを更新する' : 'スタッフにする'}
          </ModalHeader>
          <ModalCloseButton size="lg" />
          <ModalBody>
            <form onSubmit={handleSubmit(onSubmitForm)}>
              <FormControl isInvalid={!!errors.memo} mt=".5rem">
                <FormLabel htmlFor="memo">メモ（オプション）</FormLabel>
                <Textarea id="memo" {...register('memo')}></Textarea>
                <FormErrorMessage>
                  {errors.memo && errors.memo.message}
                </FormErrorMessage>
              </FormControl>

              <Button
                mt="1rem"
                isLoading={isSubmitting}
                colorScheme="cateiru"
                type="submit"
                w="100%"
              >
                {staff ? 'スタッフを更新する' : 'スタッフにする'}
              </Button>
            </form>
          </ModalBody>
        </ModalContent>
      </Modal>
    </>
  );
});

UserDetailStaff.displayName = 'UserDetailStaff';
