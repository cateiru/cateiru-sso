import {
  Modal,
  ModalOverlay,
  ModalContent,
  ModalHeader,
  ModalFooter,
  ModalBody,
  ModalCloseButton,
  Button,
  Input,
  ListItem,
  IconButton,
  UnorderedList,
  useToast,
  Text,
  useClipboard,
  InputGroup,
  InputRightElement,
  Avatar,
  Center,
  Box,
  useDisclosure,
  FormControl,
  FormLabel,
  ButtonGroup,
  FormErrorMessage,
  Switch,
  Slider,
  SliderTrack,
  SliderFilledTrack,
  SliderThumb,
  useColorMode,
  Stack,
} from '@chakra-ui/react';
import React from 'react';
import AvatarEditor from 'react-avatar-editor';
import {useForm} from 'react-hook-form';
import type {FieldValues} from 'react-hook-form';
import {
  IoCopyOutline,
  IoCheckmarkSharp,
  IoAddOutline,
  IoRemoveOutline,
} from 'react-icons/io5';
import {
  deleteService,
  editImage,
  editSSO,
  Service,
} from '../../utils/api/proSSO';

const ServiceDetails: React.FC<{
  service: Service | undefined;
  isOpen: boolean;
  onClose: () => void;
  changeService: (s: Service | null) => void;
}> = ({service, isOpen, onClose, changeService}) => {
  const toast = useToast();
  const {colorMode} = useColorMode();

  // モーダル操作
  const editModal = useDisclosure();
  const imageModal = useDisclosure();
  const deleteModal = useDisclosure();

  // 詳細
  const clientIDCopy = useClipboard(service?.client_id || '');
  const tokenSecretCopy = useClipboard(service?.token_secret || '');

  // 変更
  const [fromURLs, setFromURLs] = React.useState(1);
  const [toURLs, setToURLs] = React.useState(1);
  const {
    handleSubmit,
    register,
    formState: {errors},
    reset,
  } = useForm();

  // 画像
  const inputRef = React.useRef<HTMLInputElement>(null);
  const editorRef = React.useRef<AvatarEditor>(null);
  const [image, setImage] = React.useState<File>(new File([], ''));
  const [zoom, setZoom] = React.useState(1);

  React.useEffect(() => {
    setFromURLs(service?.from_url.length || 1);
  }, [service?.from_url]);

  React.useEffect(() => {
    setToURLs(service?.to_url.length || 1);
  }, [service?.to_url]);

  const submit = (values: FieldValues) => {
    const fromURL = new Array(fromURLs)
      .fill('')
      .map((_, index) => values[`fromurl${index}`])
      .sort();
    const toURL = new Array(toURLs)
      .fill('')
      .map((_, index) => values[`tourl${index}`])
      .sort();

    let changedName = '';
    let changedFromURL: string[] = [];
    let changedToURL: string[] = [];
    let isChanged = false;
    if (service?.name !== values.name) {
      changedName = values.name;
      isChanged = true;
    }
    if (JSON.stringify(service?.from_url.sort()) !== JSON.stringify(fromURL)) {
      changedFromURL = fromURL;
      isChanged = true;
    }
    if (JSON.stringify(service?.to_url.sort()) !== JSON.stringify(toURL)) {
      changedToURL = toURL;
      isChanged = true;
    }
    if (values.secret) {
      isChanged = true;
    }

    if (!isChanged) {
      return;
    }

    const f = async () => {
      try {
        const newService = await editSSO(
          service?.client_id || '',
          changedName,
          changedFromURL,
          changedToURL,
          values.secret
        );
        changeService(newService);
        toast({
          title: '更新しました',
          status: 'info',
          isClosable: true,
          duration: 9000,
        });
        editModal.onClose();
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

  const handlerSetImage = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files) {
      setImage(e.target.files[0]);
      onClose();
      imageModal.onOpen();
    }
  };

  const submitImage = () => {
    if (editorRef.current) {
      const canvas = editorRef.current.getImageScaledToCanvas();

      canvas.toBlob(blob => {
        if (blob) {
          const file = new File([blob], 'avatar', {type: 'image/png'});

          const f = async () => {
            try {
              const newService = await editImage(
                file,
                service?.client_id || ''
              );

              changeService(newService);

              toast({
                title: '変更しました',
                description:
                  '画像がキャッシュされている場合はしばらくしてからリロードしてください',
                status: 'info',
                isClosable: true,
              });
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
        }
      }, 'image/png');
    }

    imageModal.onClose();
  };

  const handleDelete = () => {
    const f = async () => {
      try {
        await deleteService(service?.client_id || '');
        changeService(null);
        deleteModal.onClose();
        toast({
          title: '削除しました',
          status: 'info',
          isClosable: true,
          duration: 9000,
        });
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
      <Modal isOpen={isOpen} onClose={onClose}>
        <ModalOverlay />
        <ModalContent>
          <ModalHeader pr="5rem">{service?.name}の詳細</ModalHeader>
          <ModalCloseButton size="lg" />
          <ModalBody>
            <Center>
              <Avatar
                size="lg"
                name={service?.name}
                src={service?.service_icon}
              />
            </Center>
            <Text textAlign="center" fontWeight="bold">
              {service?.name}
            </Text>
            <Text mt="1rem" mb=".2rem" fontSize="1.2rem">
              Client ID
            </Text>
            <InputGroup>
              <Input
                defaultValue={service?.client_id}
                placeholder="client id"
              />
              <InputRightElement>
                <IconButton
                  variant="ghost"
                  size="sm"
                  aria-label="copy"
                  icon={
                    clientIDCopy.hasCopied ? (
                      <IoCheckmarkSharp size="25px" color="#38A169" />
                    ) : (
                      <IoCopyOutline size="25px" />
                    )
                  }
                  onClick={clientIDCopy.onCopy}
                />
              </InputRightElement>
            </InputGroup>
            <Text mt="1rem" mb=".2rem" fontSize="1.2rem">
              Token Secret
            </Text>
            <InputGroup>
              <Input
                defaultValue={service?.token_secret}
                placeholder="token secret"
              />
              <InputRightElement>
                <IconButton
                  variant="ghost"
                  size="sm"
                  aria-label="copy"
                  icon={
                    tokenSecretCopy.hasCopied ? (
                      <IoCheckmarkSharp size="25px" color="#38A169" />
                    ) : (
                      <IoCopyOutline size="25px" />
                    )
                  }
                  onClick={tokenSecretCopy.onCopy}
                />
              </InputRightElement>
            </InputGroup>
            <Text mt="1rem" mb=".2rem" fontSize="1.2rem">
              ログイン数
            </Text>
            <Text fontSize="1.2rem" fontWeight="bold">
              {service?.login_count || 0}
            </Text>
            <Text mt="1rem" mb=".2rem" fontSize="1.2rem">
              アクセス元URL
            </Text>
            <UnorderedList>
              {service?.from_url.map(v => {
                return (
                  <ListItem key={v} ml="1rem">
                    {v}
                  </ListItem>
                );
              })}
            </UnorderedList>
            <Text mt="1rem" mb=".2rem" fontSize="1.2rem">
              リダイレクトURL
            </Text>
            <UnorderedList>
              {service?.to_url.map(v => {
                return (
                  <ListItem key={v} ml="1rem">
                    {v}
                  </ListItem>
                );
              })}
            </UnorderedList>
            <Text mt="1rem" mb=".2rem" fontSize="1.2rem">
              編集
            </Text>
            <Stack
              direction={['column', 'row']}
              spacing="10px"
              width={{base: '100%', sm: 'auto'}}
            >
              <label htmlFor="filename">
                <Button
                  type="submit"
                  mr={3}
                  as="p"
                  cursor="pointer"
                  width="100%"
                >
                  画像を設定する
                </Button>
                <Input
                  ref={inputRef}
                  display="none"
                  id="filename"
                  type="file"
                  accept="image/*"
                  onChange={handlerSetImage}
                />
              </label>
              <Button
                type="submit"
                mr={3}
                onClick={() => {
                  onClose();
                  reset();
                  editModal.onOpen();
                }}
              >
                名前、URLを変更する
              </Button>
            </Stack>
            <Text mt="1rem" mb=".2rem" fontSize="1.2rem">
              削除
            </Text>
            <Button
              variant="ghost"
              colorScheme="red"
              type="submit"
              mr={3}
              onClick={() => {
                onClose();
                deleteModal.onOpen();
              }}
            >
              削除する
            </Button>
          </ModalBody>
          <ModalFooter>
            <Button type="submit" mr={3} onClick={onClose}>
              閉じる
            </Button>
          </ModalFooter>
        </ModalContent>
      </Modal>

      <Modal isOpen={editModal.isOpen} onClose={editModal.onClose}>
        <ModalOverlay />
        <ModalContent>
          <ModalHeader pr="5rem">{service?.name}の編集</ModalHeader>
          <ModalCloseButton size="lg" />
          <form onSubmit={handleSubmit(submit)}>
            <ModalBody>
              <FormControl isInvalid={errors.name}>
                <FormLabel htmlFor="name">サービス名</FormLabel>
                <Input
                  id="name"
                  type="text"
                  placeholder="サービス名"
                  defaultValue={service?.name}
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
                    defaultValue={service?.from_url[index]}
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
                    defaultValue={service?.to_url[index]}
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
              <FormLabel mt="1rem">Token Secret</FormLabel>
              <Switch size="md" {...register('secret')}>
                Token Secretを更新する
              </Switch>
            </ModalBody>
            <ModalFooter>
              <Button colorScheme="blue" mr={3} type="submit">
                これでOK
              </Button>
            </ModalFooter>
          </form>
        </ModalContent>
      </Modal>

      <Modal isOpen={imageModal.isOpen} onClose={imageModal.onClose}>
        <ModalOverlay />
        <ModalContent>
          <ModalHeader>画像をトリミングする</ModalHeader>
          <ModalCloseButton size="lg" />
          <ModalBody>
            <Center>
              <Box>
                <Box
                  boxShadow={
                    colorMode === 'dark'
                      ? '0 5px 20px 0 #171923'
                      : '0 2px 10px 000'
                  }
                  mb="1rem"
                >
                  <AvatarEditor
                    ref={editorRef}
                    image={image}
                    width={250}
                    height={250}
                    border={0}
                    color={[113, 128, 150, 0.7]} // RGBA
                    scale={zoom}
                    rotate={0}
                    borderRadius={255}
                  />
                </Box>
                <Slider
                  colorScheme="blue"
                  aria-label="zoom"
                  defaultValue={1}
                  step={0.01}
                  max={3}
                  min={1}
                  onChange={v => setZoom(v)}
                >
                  <SliderTrack bg="gray.400">
                    <SliderFilledTrack />
                  </SliderTrack>
                  <SliderThumb />
                </Slider>
              </Box>
            </Center>
          </ModalBody>

          <ModalFooter>
            <Button colorScheme="blue" mr={3} onClick={submitImage}>
              変更する
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
          <ModalHeader>このSSOサービスを削除しますか？</ModalHeader>
          <ModalCloseButton size="lg" />
          <ModalBody>この操作は戻せません。</ModalBody>

          <ModalFooter>
            <Button colorScheme="red" mr={3} onClick={handleDelete}>
              削除する
            </Button>
            <Button mr={3} onClick={deleteModal.onClose}>
              なにもしない
            </Button>
          </ModalFooter>
        </ModalContent>
      </Modal>
    </>
  );
};

export default ServiceDetails;
