import {
  Avatar,
  Box,
  Button,
  Center,
  Input,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalHeader,
  ModalOverlay,
  Slider,
  SliderFilledTrack,
  SliderThumb,
  SliderTrack,
  Text,
  useColorModeValue,
  useDisclosure,
  useToast,
} from '@chakra-ui/react';
import React from 'react';
import AvatarEditor from 'react-avatar-editor';
import {useRecoilState} from 'recoil';
import {UserState} from '../../utils/state/atom';
import {UserAvatarSchema} from '../../utils/types/user';
import {useRequest} from '../Common/useRequest';

export const UpdateAvatar = () => {
  const [user, setUser] = useRecoilState(UserState);

  const inputRef = React.useRef<HTMLInputElement>(null);
  const editorRef = React.useRef<AvatarEditor>(null);

  const {isOpen, onOpen, onClose} = useDisclosure();
  const toast = useToast();

  const {request} = useRequest('/v2/user/avatar', {
    errorCallback: () => {
      // エラー起きたらモーダル閉じる
      closeModal();
    },
  });

  const shadowColor = useColorModeValue(
    '0 2px 10px 000',
    '0 5px 20px 0 #171923'
  );
  const sliderThumbColor = useColorModeValue('gray.500', 'white');

  const [image, setImage] = React.useState<File>(new File([], ''));
  const [zoom, setZoom] = React.useState(1);
  const [loading, setLoading] = React.useState(false);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files) {
      setImage(e.target.files[0]);
      onOpen();
    }
  };

  const closeModal = () => {
    if (inputRef.current) {
      inputRef.current.value = '';
    }
    setZoom(1);
    setLoading(false);

    onClose();
  };

  const apply = () => {
    if (editorRef.current) {
      const canvas = editorRef.current.getImageScaledToCanvas();

      canvas.toBlob(blob => {
        if (blob) {
          const file = new File([blob], 'image', {type: 'image/png'});

          const f = async () => {
            setLoading(true);

            const form = new FormData();
            form.append('image', file);

            const res = await request({
              method: 'POST',
              body: form,
              mode: 'cors',
              credentials: 'include',
            });

            if (res) {
              const data = UserAvatarSchema.safeParse(await res.json());

              if (data.success) {
                toast({
                  title: 'アバターを変更しました',
                  status: 'success',
                  duration: 3000,
                  isClosable: true,
                });

                // 一旦avatarをnullにしてから、画像を追加する
                // 更新時などでもURLは同一なため
                setUser(u => {
                  if (u) {
                    const user = {...u.user};
                    user.avatar = null;
                    return {
                      ...u,
                      user: user,
                    };
                  }
                  return u;
                });

                setTimeout(() => {
                  setUser(u => {
                    if (u) {
                      const user = {...u.user};
                      user.avatar = data.data.avatar;
                      return {
                        ...u,
                        user: user,
                      };
                    }
                    return u;
                  });
                }, 100);
              }
              closeModal();
            }
          };

          f();
        }
      }, 'image/png');
    }
  };

  return (
    <Box position="relative">
      <Avatar src={user?.user.avatar ?? ''} size="xl" />
      <label htmlFor="filename">
        <Box
          position="absolute"
          top="0"
          left="0px"
          w="96px"
          h="96px"
          borderRadius="50%"
          opacity="0"
          _hover={{
            opacity: '0.7',
          }}
          cursor="pointer"
        >
          <Text
            w="96px"
            h="48px"
            mt="48px"
            bgColor="#171923"
            borderRadius="0 0 100px 100px"
            color="#fff"
            fontWeight="bold"
            textAlign="center"
          >
            変更する
          </Text>
        </Box>
        <Input
          ref={inputRef}
          display="none"
          id="filename"
          type="file"
          accept="image/*"
          onChange={handleChange}
        />
      </label>
      <Modal isOpen={isOpen} onClose={closeModal} isCentered>
        <ModalOverlay />
        <ModalContent>
          <ModalHeader>画像をトリミングする</ModalHeader>
          <ModalCloseButton size="lg" />
          <ModalBody>
            <Center>
              <Box>
                <Box boxShadow={shadowColor} mb="1rem">
                  <AvatarEditor
                    ref={editorRef}
                    image={image}
                    width={350}
                    height={350}
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
                  max={5}
                  min={0.6}
                  onChange={v => setZoom(v)}
                >
                  <SliderTrack>
                    <SliderFilledTrack bgColor="my.secondary" />
                  </SliderTrack>
                  <SliderThumb bgColor={sliderThumbColor} />
                </Slider>
              </Box>
            </Center>
            <Button
              colorScheme="cateiru"
              mr={3}
              onClick={apply}
              w="100%"
              mt="1rem"
              isLoading={loading}
            >
              変更する
            </Button>
          </ModalBody>
        </ModalContent>
      </Modal>
    </Box>
  );
};
