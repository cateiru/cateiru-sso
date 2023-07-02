import {
  Button,
  ButtonGroup,
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
  useColorModeValue,
  useDisclosure,
} from '@chakra-ui/react';
import React from 'react';
import {useForm} from 'react-hook-form';
import {useRecoilValue} from 'recoil';
import {useSWRConfig} from 'swr';
import {UserState} from '../../../utils/state/atom';
import {Staff} from '../../../utils/types/staff';
import {useRequest} from '../../Common/useRequest';

interface FormTypes {
  memo?: string;
}

interface Props {
  staff?: Staff | null;
  userId?: string;
}

export const UserDetailStaff = React.memo<Props>(({staff, userId}) => {
  const textColor = useColorModeValue('gray.500', 'gray.400');

  const {request: staffRequest} = useRequest('/v2/admin/staff');
  const {mutate} = useSWRConfig();
  const user = useRecoilValue(UserState);

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
      {user?.user.id !== userId && (
        <>
          {staff ? (
            <Button
              w="100%"
              onClick={() => (staff ? onToggleStaff(false) : undefined)}
            >
              スタッフから外す
            </Button>
          ) : (
            <Button
              w="100%"
              colorScheme="cateiru"
              onClick={staff ? undefined : onOpen}
            >
              スタッフにする
            </Button>
          )}
        </>
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
          <ModalHeader>スタッフにしますか？</ModalHeader>
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
                スタッフにする
              </Button>
            </form>
          </ModalBody>
        </ModalContent>
      </Modal>
    </>
  );
});

UserDetailStaff.displayName = 'UserDetailStaff';