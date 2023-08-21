import {
  Box,
  Center,
  Heading,
  Image,
  Stack,
  Text,
  useColorModeValue,
} from '@chakra-ui/react';
import {config} from '../../utils/config';
import {useSecondaryColor, useShadowColor} from '../Common/useColor';

export const PageTwo = () => {
  const textColor = useSecondaryColor();
  const shadowColor = useShadowColor();
  const image = useColorModeValue('/top_1.png', '/top_2.png');

  return (
    <Box h="100vh">
      <Heading textAlign="center">{config.title}とは？</Heading>
      <Center mt="3rem">
        <Stack
          direction={{base: 'column', lg: 'row'}}
          spacing="10"
          w={{base: '100%', lg: '1000px', xl: '1200px'}}
        >
          <Box w={{base: '95%', lg: '700px'}} mx="auto">
            <Text
              color={textColor}
              fontSize="1.2rem"
              pt={{base: '0', lg: '3rem'}}
            >
              {config.title} は
              <Text
                as="span"
                fontWeight="bold"
                px=".5rem"
                background="linear-gradient(128deg, #E23D3D 0%, #EC44BD 100%)"
                backgroundClip="text"
              >
                Identity Platform (IdP)
              </Text>
              です。
              <br />
              あなたの使用しているアカウントを1つにまとめることができます。
            </Text>
          </Box>

          <Box w={{base: '95%', lg: '100%'}} mx="auto">
            <Image
              src={image}
              alt="ユーザートップの画像"
              mx="auto"
              mt="1rem"
              width="100%"
              boxShadow={`0px 0px 7px -2px ${shadowColor}`}
              borderRadius="10px"
            />
            <Text textAlign="center" color={textColor} mt=".5rem">
              ユーザートップページ
            </Text>
          </Box>
        </Stack>
      </Center>
    </Box>
  );
};
