import {
  Box,
  Center,
  List,
  ListIcon,
  ListItem,
  Text,
  useColorModeValue,
} from '@chakra-ui/react';
import {TbCheck} from 'react-icons/tb';
import {useSecondaryColor} from '../Common/useColor';

export const PageThree = () => {
  const textColor = useSecondaryColor();
  const checkMarkColor = useColorModeValue('#68D391', '#38A169');

  return (
    <Box h="100vh">
      <Center>
        <List fontSize="2rem" color={textColor} fontWeight="bold">
          <ListItem>
            <ListIcon
              as={TbCheck}
              color={checkMarkColor}
              strokeWidth="3px"
              fontSize="2.3rem"
            />
            パスキー対応で
            <Text
              as="span"
              fontWeight="bold"
              px=".5rem"
              background="linear-gradient(128deg, #E23D3D 0%, #EC44BD 100%)"
              backgroundClip="text"
            >
              3秒
            </Text>
            ログイン
          </ListItem>
          <ListItem>
            <ListIcon
              as={TbCheck}
              color={checkMarkColor}
              strokeWidth="3px"
              fontSize="2.3rem"
            />
            複数アカウントを
            <Text
              as="span"
              fontWeight="bold"
              px=".5rem"
              background="linear-gradient(128deg, #E23D3D 0%, #EC44BD 100%)"
              backgroundClip="text"
            >
              1クリック
            </Text>
            で切り替え
          </ListItem>
          <ListItem>
            <ListIcon
              as={TbCheck}
              color={checkMarkColor}
              strokeWidth="3px"
              fontSize="2.3rem"
            />
            <Text
              as="span"
              fontWeight="bold"
              px=".5rem"
              background="linear-gradient(128deg, #E23D3D 0%, #EC44BD 100%)"
              backgroundClip="text"
            >
              二段階認証
            </Text>
            に対応
          </ListItem>
          <ListItem
            background="linear-gradient(128deg, #E23D3D 0%, #EC44BD 100%)"
            backgroundClip="text"
            fontWeight="bold"
          >
            <ListIcon
              as={TbCheck}
              color={checkMarkColor}
              strokeWidth="3px"
              fontSize="2.3rem"
            />
            ソースコード公開
          </ListItem>
        </List>
      </Center>
    </Box>
  );
};
