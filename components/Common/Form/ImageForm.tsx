import {
  Box,
  Button,
  Center,
  Flex,
  Input,
  InputGroup,
  InputRightElement,
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
  useColorModeValue,
  useDisclosure,
} from '@chakra-ui/react';
import React from 'react';
import AvatarEditor from 'react-avatar-editor';
import {useFormContext} from 'react-hook-form';
import {Avatar} from '../Chakra/Avatar';

interface Props {
  imageUrl?: string;
  clearImage?: () => void;
}

export interface ImageFormValue {
  image?: File;
}

export const ImageForm: React.FC<Props> = props => {
  const {setValue} = useFormContext<ImageFormValue>();

  const {isOpen, onOpen, onClose} = useDisclosure();

  const shadowColor = useColorModeValue(
    '0 2px 10px 000',
    '0 5px 20px 0 #171923'
  );
  const sliderThumbColor = useColorModeValue('gray.500', 'white');

  const inputRef = React.useRef<HTMLInputElement>(null);
  const editorRef = React.useRef<AvatarEditor>(null);

  const [image, setImage] = React.useState<File>(new File([], ''));
  const [zoom, setZoom] = React.useState(1);
  const [imageUrl, setImageUrl] = React.useState('');

  React.useEffect(() => {
    if (props.imageUrl) {
      setImageUrl(props.imageUrl);
    }
  }, [props.imageUrl]);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files) {
      setImage(e.target.files[0]);
      onOpen();
    }
  };

  const closeModal = () => {
    clearImage();
    onClose();
  };

  const clearImage = () => {
    if (inputRef.current) {
      inputRef.current.value = '';
    }
    setImage(new File([], ''));
    setImageUrl('');
    setZoom(1);

    props.clearImage && props.clearImage();
  };

  const apply = () => {
    if (editorRef.current) {
      const canvas = editorRef.current.getImageScaledToCanvas();

      canvas.toBlob(blob => {
        if (blob) {
          const file = new File([blob], 'image', {type: 'image/png'});

          setImageUrl(URL.createObjectURL(file));
          setValue('image', file);
          setImage(new File([], ''));
          onClose();
        }
      }, 'image/png');
    }
  };

  return (
    <>
      <Flex alignItems="center">
        <Avatar src={imageUrl} size="sm" mr=".5rem" />
        <InputGroup>
          <Input
            ref={inputRef}
            id="filename"
            type="file"
            accept="image/*"
            onChange={handleChange}
          />
          <InputRightElement w="5.5rem" mr=".2rem">
            <Button size="sm" onClick={clearImage}>
              画像を削除
            </Button>
          </InputRightElement>
        </InputGroup>
      </Flex>

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
            >
              OK
            </Button>
          </ModalBody>
        </ModalContent>
      </Modal>
    </>
  );
};
