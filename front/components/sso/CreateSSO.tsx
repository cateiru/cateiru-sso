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
  IconButton,
  ButtonGroup,
  useToast,
  Text,
} from '@chakra-ui/react';
import React from 'react';
import {useForm} from 'react-hook-form';
import type {FieldValues} from 'react-hook-form';
import {IoAddOutline, IoRemoveOutline} from 'react-icons/io5';
import {setSSOs, Service} from '../../utils/api/proSSO';

const CreateSSO: React.FC<{setService: (s: Service) => void}> = ({
  setService,
}) => {
  const [fromURLs, setFromURLs] = React.useState(1);
  const [toURLs, setToURLs] = React.useState(1);
  const toast = useToast();

  const {colorMode} = useColorMode();
  const createModal = useDisclosure();
  const {
    handleSubmit,
    register,
    formState: {errors},
  } = useForm();

  const submit = (values: FieldValues) => {
    const f = async () => {
      const fromURL = new Array(fromURLs)
        .fill('')
        .map((_, index) => values[`fromurl${index}`]);
      const toURL = new Array(toURLs)
        .fill('')
        .map((_, index) => values[`tourl${index}`]);

      try {
        const service = await setSSOs(values.name, fromURL, toURL);
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
      width="100%"
      maxWidth="500px"
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
        onClose={createModal.onClose}
        isCentered
      >
        <ModalOverlay />
        <ModalContent>
          <ModalHeader>サービスを新しく作成する</ModalHeader>
          <ModalCloseButton size="lg" />
          <form onSubmit={handleSubmit(submit)}>
            <ModalBody>
              <Text color="yellow.500" mb="1rem">
                * アイコン画像は作成後に追加できます
              </Text>
              <FormControl isInvalid={errors.name}>
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
              <FormLabel mt="1rem">送信元URL</FormLabel>
              {new Array(fromURLs).fill(0).map((_, index) => (
                <FormControl
                  isInvalid={errors[`fromurl${index}`]}
                  key={index}
                  my=".5rem"
                >
                  <Input
                    id={`fromurl${index}`}
                    type="text"
                    placeholder={`送信元URL ${index + 1}`}
                    {...register(`fromurl${index}`, {
                      required: '送信元URL の入力が必要です',
                      pattern: {
                        value:
                          /(https:\/\/[\w/:%#$&?()~.=+-]+|http:\/\/localhost|direct)/,
                        message: 'URLの形式が違うようです',
                      },
                    })}
                  />
                  <FormErrorMessage>
                    {errors[`fromurl${index}`] &&
                      errors[`fromurl${index}`].message}
                  </FormErrorMessage>
                </FormControl>
              ))}
              <ButtonGroup isAttached>
                <IconButton
                  aria-label="add"
                  icon={<IoAddOutline size="25px" />}
                  onClick={() => {
                    setFromURLs(v => {
                      if (v >= 5) {
                        return v;
                      }
                      return (v += 1);
                    });
                  }}
                />
                <IconButton
                  aria-label="add"
                  icon={<IoRemoveOutline size="25px" />}
                  onClick={() => {
                    setFromURLs(v => {
                      if (v <= 1) {
                        return v;
                      }
                      return (v -= 1);
                    });
                  }}
                />
              </ButtonGroup>
              <FormLabel mt="1rem">
                リダイレクトURL（しない場合はdirect）
              </FormLabel>
              {new Array(toURLs).fill(0).map((_, index) => (
                <FormControl
                  isInvalid={errors[`tourl${index}`]}
                  key={index}
                  my=".5rem"
                >
                  <Input
                    id={`tourl${index}`}
                    type="text"
                    placeholder={`リダイレクトURL ${index + 1}`}
                    {...register(`tourl${index}`, {
                      required: 'リダイレクトURL の入力が必要です',
                      pattern: {
                        value:
                          /(https:\/\/[\w/:%#$&?()~.=+-]+|http:\/\/localhost|direct)/,
                        message: 'URLの形式が違うようです',
                      },
                    })}
                  />
                  <FormErrorMessage>
                    {errors[`tourl${index}`] && errors[`tourl${index}`].message}
                  </FormErrorMessage>
                </FormControl>
              ))}
              <ButtonGroup isAttached>
                <IconButton
                  aria-label="add"
                  icon={<IoAddOutline size="25px" />}
                  onClick={() => {
                    setToURLs(v => {
                      if (v >= 5) {
                        return v;
                      }
                      return (v += 1);
                    });
                  }}
                />
                <IconButton
                  aria-label="add"
                  icon={<IoRemoveOutline size="25px" />}
                  onClick={() => {
                    setToURLs(v => {
                      if (v <= 1) {
                        return v;
                      }
                      return (v -= 1);
                    });
                  }}
                />
              </ButtonGroup>
            </ModalBody>
            <ModalFooter>
              <Button colorScheme="blue" type="submit" mr={3}>
                作成
              </Button>
            </ModalFooter>
          </form>
        </ModalContent>
      </Modal>
    </Box>
  );
};

export default CreateSSO;
