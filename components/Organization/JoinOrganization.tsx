import {
  Button,
  FormControl,
  FormErrorMessage,
  FormLabel,
  Input,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalHeader,
  ModalOverlay,
  Select,
  Text,
  useColorModeValue,
  useDisclosure,
  useToast,
} from '@chakra-ui/react';
import React from 'react';
import {useForm} from 'react-hook-form';
import {emailRegex} from '../../utils/regex';
import {ErrorUniqueMessage} from '../../utils/types/error';
import {useSecondaryColor} from '../Common/useColor';
import {useRequest} from '../Common/useRequest';

interface Props {
  orgId: string;
  handleSuccess: (isOrgMember: boolean) => void;
}

interface JoinFormData {
  user_name_or_email: string;
  role: string;
}

interface InviteFormData {
  email: string;
}

interface InviteFormProps extends Props {
  email: string;
}

const InviteForm: React.FC<InviteFormProps> = ({
  orgId,
  handleSuccess,
  email,
}) => {
  const {
    handleSubmit,
    register,
    reset,
    formState: {errors, isSubmitting},
  } = useForm<InviteFormData>({
    defaultValues: {
      email,
    },
  });

  const {request} = useRequest('/org/member/invite');

  const onSubmit = async (data: InviteFormData) => {
    const form = new FormData();

    form.append('org_id', orgId);
    form.append('email', data.email);

    const res = await request({
      method: 'POST',
      body: form,
    });

    if (res) {
      reset();
      handleSuccess(false);
    }
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <FormControl isInvalid={!!errors.email}>
        <FormLabel htmlFor="email">メールアドレス</FormLabel>
        <Input
          type="email"
          id="email"
          autoComplete="email"
          {...register('email', {
            required: 'メールアドレスは必須です',
          })}
        />
        <FormErrorMessage>
          {errors.email && errors.email.message}
        </FormErrorMessage>
      </FormControl>
      <Button
        mt="1rem"
        w="100%"
        colorScheme="cateiru"
        type="submit"
        isLoading={isSubmitting}
      >
        ユーザーを招待
      </Button>
    </form>
  );
};

interface JoinFormProps extends Props {
  setInvite: (el: boolean) => void;
  setUserOrEmail: (el: string) => void;
}

const JoinForm: React.FC<JoinFormProps> = ({
  orgId,
  handleSuccess,
  setInvite,
  setUserOrEmail,
}) => {
  const toast = useToast();

  const {request} = useRequest('/org/member', {
    customError: e => {
      // 10 = NotFoundUser の場合は別のフォームを開く
      if (e.unique_code === 10) {
        setInvite(true);
        return;
      }

      toast({
        title: ErrorUniqueMessage[e.unique_code ?? 0] ?? e.message,
        status: 'error',
      });
    },
  });

  const {
    handleSubmit,
    register,
    reset,
    formState: {errors, isSubmitting},
  } = useForm<JoinFormData>({
    defaultValues: {
      role: 'guest',
    },
  });

  // org登録
  const onSubmit = async (data: JoinFormData) => {
    const form = new FormData();

    form.append('org_id', orgId);
    form.append('user_name_or_email', data.user_name_or_email);
    form.append('role', data.role);

    if (emailRegex.test(data.user_name_or_email)) {
      setUserOrEmail(data.user_name_or_email);
    }

    const res = await request({
      method: 'POST',
      body: form,
    });

    if (res) {
      reset();
      handleSuccess(true);
    }
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <FormControl isInvalid={!!errors.user_name_or_email}>
        <FormLabel htmlFor="user_name_or_email">
          ユーザー名またはメールアドレス
        </FormLabel>
        <Input
          type="text"
          id="user_name_or_email"
          {...register('user_name_or_email', {
            required: 'ユーザーは必須です',
          })}
        />
        <FormErrorMessage>
          {errors.user_name_or_email && errors.user_name_or_email.message}
        </FormErrorMessage>
      </FormControl>
      <FormControl isInvalid={!!errors.role} mt=".5rem">
        <FormLabel htmlFor="role">ロール</FormLabel>
        <Select
          {...register('role', {
            required: 'ロールは必須です',
          })}
        >
          <option value="owner">管理者</option>
          <option value="member">メンバー</option>
          <option value="guest">ゲスト</option>
        </Select>

        <FormErrorMessage>
          {errors.role && errors.role.message}
        </FormErrorMessage>
      </FormControl>

      <Button
        mt="1rem"
        w="100%"
        colorScheme="cateiru"
        type="submit"
        isLoading={isSubmitting}
      >
        ユーザーを追加
      </Button>
    </form>
  );
};

export const JoinOrganization: React.FC<Props> = props => {
  const textColor = useSecondaryColor();
  const buttonColor = useColorModeValue('my.primary', 'my.secondary');
  const joinOrgModal = useDisclosure();

  const [invite, setInvite] = React.useState(false);
  const [userOrEmail, setUserOrEmail] = React.useState('');

  const closeModal = () => {
    joinOrgModal.onClose();
    setInvite(false);
  };

  const handleSuccess = (isOrgMember: boolean) => {
    props.handleSuccess(isOrgMember);
    closeModal();
  };

  return (
    <>
      <Button w="100%" colorScheme="cateiru" onClick={joinOrgModal.onOpen}>
        組織に招待する
      </Button>
      <Modal isOpen={joinOrgModal.isOpen} onClose={closeModal} isCentered>
        <ModalOverlay />
        <ModalContent>
          <ModalHeader>組織に招待する</ModalHeader>
          <ModalCloseButton size="lg" />
          <ModalBody mb=".5rem">
            {invite ? (
              <>
                <Text mb="1rem" color={textColor}>
                  メールアドレスに対して招待メールを送信します。すでにユーザーが存在している場合は、
                  <Button
                    variant="link"
                    onClick={() => setInvite(false)}
                    fontWeight="bold"
                    color={buttonColor}
                  >
                    メールアドレスまたはユーザー名を指定した招待
                  </Button>
                  を使用してください。
                </Text>
                <InviteForm
                  orgId={props.orgId}
                  handleSuccess={handleSuccess}
                  email={userOrEmail}
                />
              </>
            ) : (
              <>
                <JoinForm
                  orgId={props.orgId}
                  handleSuccess={handleSuccess}
                  setInvite={setInvite}
                  setUserOrEmail={setUserOrEmail}
                />
              </>
            )}
          </ModalBody>
        </ModalContent>
      </Modal>
    </>
  );
};
