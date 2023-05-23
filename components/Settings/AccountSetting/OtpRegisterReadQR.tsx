import {
  Box,
  Center,
  IconButton,
  Input,
  InputGroup,
  InputRightElement,
  Text,
  Tooltip,
  useClipboard,
  useColorModeValue,
} from '@chakra-ui/react';
import QRcode from 'qrcode.react';
import React from 'react';
import {TbCheck, TbCopy} from 'react-icons/tb';
import {colorTheme} from '../../../utils/theme';

interface Props {
  token: string;
}

export const OtpRegisterReadQR: React.FC<Props> = props => {
  const bgColor = useColorModeValue(
    colorTheme.lightBackground,
    colorTheme.darkBackground
  );
  const fgColor = useColorModeValue('#572bcf', '#2bc4cf');
  const checkMarkColor = useColorModeValue('#68D391', '#38A169');
  const textColor = useColorModeValue('gray.500', 'gray.400');

  const {hasCopied, onCopy} = useClipboard(props.token);

  return (
    <>
      <Box width="100%">
        <Text mb="1rem" color={textColor}>
          アプリでQRコードを読み込むか、URLをコピーしてワンタイムパスワードを生成してください。
        </Text>
        <Center>
          <QRcode
            value={props.token}
            size={200}
            bgColor={bgColor}
            fgColor={fgColor}
          />
        </Center>
        <Text mt="1rem">もしくは、セットアップキーを直接入力してください</Text>
        <InputGroup mt=".2rem">
          <Input defaultValue={props.token} onFocus={e => e.target.select()} />
          <Tooltip label="コピー" hasArrow borderRadius="7px">
            <InputRightElement>
              <IconButton
                aria-label="copy"
                size="sm"
                onClick={onCopy}
                icon={
                  hasCopied ? (
                    <TbCheck
                      size="25px"
                      color={checkMarkColor}
                      strokeWidth="3px"
                    />
                  ) : (
                    <TbCopy size="25px" />
                  )
                }
              />
            </InputRightElement>
          </Tooltip>
        </InputGroup>
      </Box>
    </>
  );
};
