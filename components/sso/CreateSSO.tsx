import {
  Modal,
  ModalOverlay,
  ModalContent,
  ModalHeader,
  ModalFooter,
  ModalBody,
  ModalCloseButton,
  Button,
  Box,
  useColorMode,
  useDisclosure,
  Center,
  FormControl,
  FormLabel,
  Input,
  FormErrorMessage,
  useToast,
  Text,
} from '@chakra-ui/react';
import React from 'react';
import {useForm, FormProvider, SubmitHandler} from 'react-hook-form';
import {IoAddOutline} from 'react-icons/io5';
import {setSSOs, Service} from '../../utils/api/proSSO';
import FromURLs, {FromURLForm} from './Form/FromURLs';
import ToURLs, {ToURLForm} from './Form/ToURLs';

interface Form extends FromURLForm, ToURLForm {
  name: string;
  roles: string;
}

const CreateSSO: React.FC<{setService: (s: Service) => void}> = ({
  setService,
}) => {
  const toast = useToast();

  const {colorMode} = useColorMode();
  const createModal = useDisclosure();
  const methods = useForm<Form>({
    defaultValues: {fromUrls: [{url: ''}], toUrls: [{url: ''}]},
  });
  const {
    handleSubmit,
    register,
    reset,
    formState: {errors},
  } = methods;

  const submit: SubmitHandler<Form> = values => {
    const f = async () => {
      const fromURL = values.fromUrls.map(v => v.url);
      const toURL = values.toUrls.map(v => v.url);

      const roles = values.roles.length === 0 ? [] : values.roles.split(',');

      try {
        const service = await setSSOs(values.name, fromURL, toURL, roles);
        setService(service);
        toast({
          title: '作成しました',
          status: 'info',
          isClosable: true,
          duration: 9000,
        });
        createModal.onClose();
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
    <Box
      maxWidth="350px"
      minWidth="300px"
      height="10rem"
      borderRadius="23px"
      cursor="pointer"
      transition="all 0.5s"
      onClick={createModal.onOpen}
      color={colorMode === 'dark' ? 'gray.500' : 'gray.400'}
      boxShadow={
        colorMode === 'dark'
          ? '10px 10px 30px #000000CC, -10px -10px 30px #4A5568CC, inset 10px 10px 30px transparent, inset -10px -10px 30px transparent;'
          : '10px 10px 30px #A0AEC0B3, -10px -10px 30px #F7FAFCE6, inset 10px 10px 30px transparent, inset -10px -10px 30px transparent;'
      }
      _hover={{
        boxShadow:
          colorMode === 'dark'
            ? '10px 10px 30px transparent, -10px -10px 30px transparent, inset 10px 10px 30px #000000CC, inset -10px -10px 30px #4A5568CC;'
            : '10px 10px 30px transparent, -10px -10px 30px transparent, inset 10px 10px 30px #A0AEC0B3, inset -10px -10px 30px #F7FAFCE6;',
        color: colorMode === 'dark' ? 'gray.600' : 'gray.300',
      }}
    >
      <Center height="100%">
        <IoAddOutline size="80px" />
      </Center>
      <Modal
        isOpen={createModal.isOpen}
        onClose={() => {
          reset();
          createModal.onClose();
        }}
        isCentered
      >
        <ModalOverlay />
        <ModalContent>
          <ModalHeader>サービスを新しく作成する</ModalHeader>
          <ModalCloseButton size="lg" />
          <FormProvider {...methods}>
            <form onSubmit={handleSubmit(submit)}>
              <ModalBody>
                <Text color="yellow.500" mb="1rem">
                  * アイコン画像は作成後に追加できます
                </Text>
                <FormControl isInvalid={Boolean(errors.name)}>
                  <FormLabel htmlFor="name">サービス名</FormLabel>
                  <Input
                    id="name"
                    type="text"
                    placeholder="サービス名"
                    {...register('name', {
                      required: 'サービス名 の入力が必要です',
                      maxLength: {
                        value: 20,
                        message: '20文字以内で入力してください',
                      },
                      minLength: {
                        value: 1,
                        message: '1文字以上で入力してください',
                      },
                    })}
                  />
                  <FormErrorMessage>
                    {errors.name && errors.name.message}
                  </FormErrorMessage>
                </FormControl>
                <FromURLs />
                <ToURLs />
                <FormControl isInvalid={Boolean(errors.roles)}>
                  <FormLabel htmlFor="roles" mt="1rem">
                    ロール（オプション）
                  </FormLabel>
                  <Input
                    id="roles"
                    type="text"
                    placeholder="ロール（,区切り）"
                    {...register('roles', {
                      pattern: {
                        value: /(,?[0-9a-z]+)*/,
                        message: 'ロールはコンマ区切りで入力してください',
                      },
                    })}
                  />
                  <FormErrorMessage>
                    {errors.roles && errors.roles.message}
                  </FormErrorMessage>
                </FormControl>
              </ModalBody>
              <ModalFooter>
                <Button colorScheme="blue" type="submit" mr={3}>
                  作成
                </Button>
              </ModalFooter>
            </form>
          </FormProvider>
        </ModalContent>
      </Modal>
    </Box>
  );
};

export default CreateSSO;
