import {
  Modal,
  ModalOverlay,
  ModalContent,
  ModalHeader,
  ModalFooter,
  ModalBody,
  ModalCloseButton,
  useDisclosure,
  Button,
  useToast,
  Center,
  Stack,
  Box,
  Input,
  InputGroup,
  InputRightElement,
  IconButton,
  useClipboard,
  Divider,
  Text,
  SimpleGrid,
  Skeleton,
} from '@chakra-ui/react';
import {useColorMode} from '@chakra-ui/react';
import QRcode from 'qrcode.react';
import React from 'react';
import {IoCopyOutline, IoCheckmarkSharp} from 'react-icons/io5';
import {useSetRecoilState, useRecoilState} from 'recoil';
import {isEnableOTP} from '../../utils/api/check';
import {
  OTPGetResponse,
  getToken,
  setToken,
  getBackups,
  deleteotp,
} from '../../utils/api/otp';
import {LoadState, OTPEnableState} from '../../utils/state/atom';
import {OTPState} from '../../utils/state/types';

const OTP = () => {
  const {isOpen, onOpen, onClose} = useDisclosure();
  const backSaveModal = useDisclosure();
  const deleteOtpModal = useDisclosure();

  const [otpGenerate, setOTPGenerate] = React.useState(false);
  const [otpEnable, setOTPEnable] = useRecoilState(OTPEnableState);
  const [otpToken, setOTPToken] = React.useState<OTPGetResponse>();
  const [passcode, setPasscode] = React.useState('');
  const [backups, setBackups] = React.useState<string[]>([]);
  const toast = useToast();
  const {colorMode} = useColorMode();
  const {hasCopied, onCopy} = useClipboard(otpToken?.otp_token || '');
  const backupCopy = useClipboard(backups.join(', '));
  const [isError, setIsError] = React.useState(false);

  const setLoad = useSetRecoilState(LoadState);

  // OTPが設定されているかを確認する
  React.useEffect(() => {
    const f = async () => {
      if (otpEnable !== OTPState.Loading) {
        return;
      }

      try {
        const isOTP = await isEnableOTP();
        if (isOTP) {
          setOTPEnable(OTPState.Enable);
        } else {
          setOTPEnable(OTPState.Disable);
        }
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
  }, []);

  // OTP設定モーダルを開く
  const setOTP = () => {
    const f = async () => {
      try {
        const resp = await getToken();
        setOTPToken(resp);
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

  const generate = () => {
    setOTPGenerate(true);
    setOTP();
  };

  const closeCreateModal = () => {
    setOTPGenerate(false);
    onClose();
  };

  // パスコードを送信してOTPを設定する
  const submitOTP = () => {
    const f = async () => {
      // パスコードが入力されていな場合はエラーにする
      if (!passcode) {
        setIsError(true);
        return;
      }
      setIsError(false);
      setLoad(true);

      try {
        const backups = await setToken(otpToken?.id || '', passcode);
        setBackups(backups);

        toast({
          title: '二段階認証を設定しました',
          status: 'info',
          isClosable: true,
          duration: 9000,
        });

        setPasscode('');
        onClose();
        setOTPEnable(OTPState.Enable);
        backSaveModal.onOpen();
      } catch (error) {
        if (error instanceof Error) {
          toast({
            title: 'ワンタイムパスワードの設定に失敗しました',
            description: error.message,
            status: 'error',
            isClosable: true,
            duration: 9000,
          });
        }
        setOTPGenerate(false);
        onClose();
      }

      setLoad(false);
    };

    f();
  };

  // OTPを削除する
  const deleteOTP = () => {
    const f = async () => {
      setLoad(true);
      try {
        // パスコードが入力されていな場合はエラーにする
        if (!passcode) {
          setIsError(true);
          return;
        }
        setIsError(false);

        await deleteotp(passcode);

        deleteOtpModal.onClose();

        setPasscode('');
        setOTPEnable(OTPState.Disable);

        toast({
          title: '二段階認証を削除しました',
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
        deleteOtpModal.onClose();
      }
      setLoad(false);
    };

    f();
  };

  // バックアップコードモーダルを開く
  const openBackups = () => {
    const f = async () => {
      setLoad(true);
      try {
        const backups = await getBackups();
        setBackups(backups);
        backSaveModal.onOpen();
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
      setLoad(false);
    };

    f();
  };

  switch (otpEnable) {
    case OTPState.Disable:
      return (
        <>
          <Button onClick={onOpen} width={{base: '100%', sm: 'auto'}}>
            二段階認証を設定する
          </Button>
          <Modal isOpen={isOpen} onClose={closeCreateModal} isCentered>
            <ModalOverlay />
            <ModalContent>
              <ModalHeader>二段階認証を設定します</ModalHeader>
              <ModalCloseButton size="lg" />
              <ModalBody>
                {otpToken && otpGenerate ? (
                  <Center>
                    <Box width="100%">
                      <Text mb="1rem">
                        アプリでQRコードを読み込むか、URLをコピーしてワンタイムパスワードを生成してください。
                      </Text>
                      <Center>
                        <QRcode
                          value={otpToken?.otp_token}
                          size={200}
                          bgColor={colorMode === 'dark' ? '#2D3748' : '#FFFFFF'}
                          fgColor={colorMode === 'dark' ? '#FFFFFF' : '#000000'}
                        />
                      </Center>
                      <InputGroup mt="1rem">
                        <Input
                          placeholder="OTPのURL"
                          type="url"
                          defaultValue={otpToken?.otp_token}
                        />
                        <InputRightElement>
                          <IconButton
                            variant="ghost"
                            aria-label="copy"
                            size="sm"
                            onClick={onCopy}
                            icon={
                              hasCopied ? (
                                <IoCheckmarkSharp size="20px" color="#38A169" />
                              ) : (
                                <IoCopyOutline size="20px" />
                              )
                            }
                          />
                        </InputRightElement>
                      </InputGroup>
                      <Divider my="1rem" />
                      <Input
                        placeholder="確認のため、生成されたコードを入力"
                        type="number"
                        onChange={e => setPasscode(e.target.value)}
                        isInvalid={isError}
                      />
                    </Box>
                  </Center>
                ) : (
                  <>
                    <Text>
                      二段階認証を設定すると、アカウントのセキュリティがより強化されます。
                    </Text>
                    <Text mt=".2rem">
                      この機能を使用するには、
                      <strong>ワンタイムパスワードを生成できるアプリ</strong>
                      が必要です。
                    </Text>
                    <Center my="2rem">
                      <Button
                        onClick={generate}
                        colorScheme="blue"
                        width="100%"
                      >
                        ワンタイムパスワードを有効にする
                      </Button>
                    </Center>
                  </>
                )}
              </ModalBody>

              {otpGenerate && (
                <ModalFooter>
                  <Button colorScheme="blue" mr={3} onClick={submitOTP}>
                    設定する
                  </Button>
                </ModalFooter>
              )}
            </ModalContent>
          </Modal>
        </>
      );
    case OTPState.Enable:
      return (
        <>
          <Stack direction={['column', 'row']} spacing="1rem">
            <Button onClick={openBackups}>バックアップコードを表示する</Button>
            <Button onClick={deleteOtpModal.onOpen} colorScheme="red">
              二段階認証を削除する
            </Button>
          </Stack>
          <Modal
            isOpen={backSaveModal.isOpen}
            onClose={backSaveModal.onClose}
            isCentered
          >
            <ModalOverlay />
            <ModalContent>
              <ModalHeader>バックアップコード</ModalHeader>
              <ModalCloseButton size="lg" />
              <ModalBody>
                <Text color="red.500" fontWeight="bold">
                  *必ず大切に保管してください
                </Text>
                <Text mt=".5rem">
                  バックアップコードはワンタイムパスワードを忘れてしまった、削除されてしまった場合に入力できるコードです。
                </Text>
                <Text mt=".5rem">コードは1つにつき1回入力できます。</Text>
                <Divider my="1rem" />
                <SimpleGrid columns={2} spacing="10px" my="1rem">
                  {backups.map(v => (
                    <Text key={v} textAlign="center">
                      {v}
                    </Text>
                  ))}
                </SimpleGrid>
                <Center mb="1rem">
                  <Button
                    onClick={backupCopy.onCopy}
                    leftIcon={
                      backupCopy.hasCopied ? (
                        <IoCheckmarkSharp size="20px" color="#38A169" />
                      ) : (
                        <IoCopyOutline size="20px" />
                      )
                    }
                  >
                    コピーする
                  </Button>
                </Center>
              </ModalBody>
            </ModalContent>
          </Modal>
          <Modal
            isOpen={deleteOtpModal.isOpen}
            onClose={deleteOtpModal.onClose}
            isCentered
          >
            <ModalOverlay />
            <ModalContent>
              <ModalHeader>二段階認証を削除しますか</ModalHeader>
              <ModalCloseButton size="lg" />
              <ModalBody>
                <Text>
                  二段階認証を削除するとアカウントが危険にさらされる恐れがあります。
                </Text>
                <Text mt=".5rem">
                  削除する場合は、ワンタイムパスワードを入力して「削除する」を押してください
                </Text>
                <Divider my="1rem" />
                <Input
                  placeholder="ワンタイムパスワードを入力"
                  type="text"
                  onChange={e => setPasscode(e.target.value)}
                  isInvalid={isError}
                />
              </ModalBody>
              <ModalFooter>
                <Button colorScheme="red" mr={3} onClick={deleteOTP}>
                  削除する
                </Button>
              </ModalFooter>
            </ModalContent>
          </Modal>
        </>
      );
    default:
      return (
        <Skeleton>
          <Button></Button>
        </Skeleton>
      );
  }
};

export default OTP;
