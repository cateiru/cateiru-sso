import {
  Box,
  Heading,
  Center,
  SimpleGrid,
  useColorMode,
  useToast,
  Text,
  Modal,
  ModalOverlay,
  ModalContent,
  ModalHeader,
  ModalFooter,
  ModalBody,
  ModalCloseButton,
  useDisclosure,
  Button,
  Divider,
  UnorderedList,
  ListItem,
} from '@chakra-ui/react';
import Link from 'next/link';
import React from 'react';
import {IoArrowBackOutline} from 'react-icons/io5';
import {useSetRecoilState} from 'recoil';
import {getUserSSO, ServiceLogInfo, deleteSSO} from '../../utils/api/userSSO';
import {hawManyDaysAgo, formatDate} from '../../utils/date';
import {LoadState} from '../../utils/state/atom';
import Avatar from '../common/Avatar';

const ConnectedAccountPage = () => {
  const {colorMode} = useColorMode();
  const [services, setServices] = React.useState<ServiceLogInfo[]>([]);
  const toast = useToast();
  const [selectService, setSelectService] = React.useState<ServiceLogInfo>();
  const [selectIndex, setSelectIndex] = React.useState(0);
  const {isOpen, onOpen, onClose} = useDisclosure();
  const deleteModal = useDisclosure();
  const setLoad = useSetRecoilState(LoadState);

  React.useEffect(() => {
    const f = async () => {
      setLoad(true);
      try {
        const s = await getUserSSO();
        setServices(s);
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
  }, []);

  const ServiceElement = (s: ServiceLogInfo, i: number) => {
    const latestDate = s.logs.reduce((p, v) => {
      if (Date.parse(p.accept_date) - Date.parse(v.accept_date) > 0) {
        return p;
      }
      return v;
    });

    return (
      <Box
        key={s.client_id}
        maxWidth="350px"
        minWidth="300px"
        height="10rem"
        borderRadius="23px"
        cursor="pointer"
        transition="all 0.5s"
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
        onClick={() => {
          setSelectService(s);
          onOpen();
          setSelectIndex(i);
        }}
      >
        <Avatar mt="1.2rem" ml="2rem" name={s.name} src={s.service_icon} />
        <Heading
          mt=".5rem"
          ml="2rem"
          textOverflow="ellipsis"
          overflow="hidden"
          whiteSpace="nowrap"
        >
          {s.name}
        </Heading>
        <Text
          textOverflow="ellipsis"
          overflow="hidden"
          whiteSpace="nowrap"
          width="60%"
          mx="2rem"
        >
          {hawManyDaysAgo(new Date(latestDate.accept_date))}
        </Text>
      </Box>
    );
  };

  const deleteService = () => {
    const f = async () => {
      setLoad(true);
      try {
        await deleteSSO(selectService?.client_id || '');
        setServices(s => {
          const cloneServices = [...s];
          cloneServices.splice(selectIndex, 1);

          return cloneServices;
        });
        deleteModal.onClose();
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

  return (
    <Box minHeight="80vh">
      <Center>
        <Box mx=".5rem" width={{base: '100%', lg: '1000px'}} mt="2rem">
          <Link href="/setting/account" passHref>
            <Button
              pl=".5rem"
              variant="ghost"
              leftIcon={<IoArrowBackOutline size="25px" />}
            >
              戻る
            </Button>
          </Link>
        </Box>
      </Center>
      <Heading textAlign="center" mb="2rem" mt="2.5rem">
        接続しているSSOサービス
      </Heading>
      <Center>
        <SimpleGrid
          columns={{base: 1, sm: 1, md: 2, lg: 3}}
          spacing="2.5rem"
          mx="1rem"
        >
          {services.map(ServiceElement)}
        </SimpleGrid>
      </Center>
      <Modal isOpen={isOpen} onClose={onClose}>
        <ModalOverlay />
        <ModalContent>
          <ModalHeader pr="5rem">{selectService?.name}</ModalHeader>
          <ModalCloseButton size="lg" />
          <ModalBody>
            <Center>
              <Avatar
                size="xl"
                name={selectService?.name}
                src={selectService?.service_icon}
                isShadow
              />
            </Center>
            <Heading textAlign="center" fontWeight="bold" fontSize="1.5rem">
              {selectService?.name}
            </Heading>
            <Divider mt="1rem" mb="2rem" borderWidth="1px" />
            <Heading fontSize="1.2rem" mb=".5rem">
              ログイン履歴
            </Heading>
            <UnorderedList>
              {selectService?.logs
                .sort(
                  (a, b) =>
                    Date.parse(b.accept_date) - Date.parse(a.accept_date)
                )
                .map(v => {
                  const d = new Date(v.accept_date);

                  return (
                    <ListItem key={v.log_id} ml="1rem">
                      {hawManyDaysAgo(d)}（{formatDate(d)}）
                    </ListItem>
                  );
                })}
            </UnorderedList>
            <Heading fontSize="1.2rem" mb=".5rem" mt="1rem">
              連携解除
            </Heading>
            <Button
              variant="ghost"
              colorScheme="red"
              onClick={() => {
                onClose();
                deleteModal.onOpen();
              }}
            >
              連携を解除する
            </Button>
          </ModalBody>

          <ModalFooter>
            <Button mr={3} onClick={onClose}>
              閉じる
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
          <ModalHeader pr="5rem">{selectService?.name}の連携を解除</ModalHeader>
          <ModalCloseButton size="lg" />
          <ModalBody>
            連携を解除すると、そのサービスからログアウトされる可能性があります。
          </ModalBody>

          <ModalFooter>
            <Button colorScheme="red" onClick={deleteService}>
              連携を解除する
            </Button>
            <Button ml=".2rem" variant="ghost" onClick={deleteModal.onClose}>
              閉じる
            </Button>
          </ModalFooter>
        </ModalContent>
      </Modal>
    </Box>
  );
};

export default ConnectedAccountPage;
