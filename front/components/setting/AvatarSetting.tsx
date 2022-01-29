import {
  Box,
  Avatar,
  Text,
  Modal,
  ModalOverlay,
  ModalCloseButton,
  ModalBody,
  ModalFooter,
  Button,
  ModalContent,
  ModalHeader,
  useDisclosure,
  Input,
  Slider,
  SliderTrack,
  SliderFilledTrack,
  SliderThumb,
  Center,
  useColorMode,
  useToast,
} from '@chakra-ui/react';
import React from 'react';
import AvatarEditor from 'react-avatar-editor';
import {useRecoilState} from 'recoil';
import {setAvatar} from '../../utils/api/avatar';
import {UserState} from '../../utils/state/atom';

const AvatarSetting = () => {
  const inputRef = React.useRef<HTMLInputElement>(null);
  const editorRef = React.useRef<AvatarEditor>(null);

  const [user, setUser] = useRecoilState(UserState);
  const {isOpen, onOpen, onClose} = useDisclosure();
  const {colorMode} = useColorMode();
  const toast = useToast();

  const [image, setImage] = React.useState<File>(new File([], ''));
  const [zoom, setZoom] = React.useState(1);

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
    onClose();
  };

  const apply = () => {
    if (editorRef.current) {
      const canvas = editorRef.current.getImageScaledToCanvas();

      canvas.toBlob(blob => {
        if (blob) {
          const file = new File([blob], 'avatar', {type: 'image/*'});

          const f = async () => {
            try {
              const url = await setAvatar(file);

              // urlを更新する
              if (user) {
                setUser({
                  ...user,
                  avatar_url: url,
                });
              }

              toast({
                title: '変更しました',
                description: '画像が変更されない場合はリロードしてください',
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
      }, 'image/*');
    }

    onClose();
  };
  return (
    <Box>
      <Box width={{base: '80px', md: '120px'}}>
        <Avatar src={user?.avatar_url} size="full" />
      </Box>
      <label htmlFor="filename">
        <Text
          mt=".3rem"
          textAlign="center"
          cursor="pointer"
          _hover={{textDecoration: 'underline'}}
        >
          変更する
        </Text>
        <Input
          ref={inputRef}
          display="none"
          id="filename"
          type="file"
          onChange={handleChange}
        />
      </label>
      <Modal isOpen={isOpen} onClose={closeModal} size="md" isCentered>
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
            <Button colorScheme="blue" mr={3} onClick={apply}>
              変更する
            </Button>
          </ModalFooter>
        </ModalContent>
      </Modal>
    </Box>
  );
};

export default AvatarSetting;
