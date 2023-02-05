import {
  SimpleGrid,
  useToast,
  Box,
  Heading,
  Text,
  useColorMode,
  Center,
  useDisclosure,
} from '@chakra-ui/react';
import React from 'react';
import {useSetRecoilState} from 'recoil';
import {getSSOs, Service} from '../../utils/api/proSSO';
import {LoadState} from '../../utils/state/atom';
import Avatar from '../common/Avatar';
import CreateSSO from './CreateSSO';
import ServiceDetails from './ServiceDetails';

const SSOPage = () => {
  const [services, setServices] = React.useState<Service[]>([]);
  const toast = useToast();
  const {colorMode} = useColorMode();
  const [selectService, setSelectService] = React.useState<Service>();
  const [selectIndex, setSelectIndex] = React.useState(0);
  const detailsModal = useDisclosure();
  const setLoad = useSetRecoilState(LoadState);

  React.useEffect(() => {
    const f = async () => {
      setLoad(true);
      try {
        const getServices = await getSSOs();
        setServices(getServices);
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

  const ServiceItem = (s: Service, i: number) => {
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
          setSelectIndex(i);
          detailsModal.onOpen();
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
          {s.name} ({s.login_count || 0})
        </Heading>
        <Text
          textOverflow="ellipsis"
          overflow="hidden"
          whiteSpace="nowrap"
          width="60%"
          mx="2rem"
        >
          ID: {s.client_id}
        </Text>
      </Box>
    );
  };

  return (
    <Box minHeight="80vh">
      <Heading textAlign="center" mb="2rem" mt="2.5rem">
        SSO Services
      </Heading>
      <Center>
        <SimpleGrid
          columns={{base: 1, sm: 1, md: 2, lg: 3}}
          spacing="2.5rem"
          mx="1rem"
        >
          {services.map(ServiceItem)}
          <CreateSSO
            setService={s => {
              setServices(v => [s, ...v]);
            }}
          />
        </SimpleGrid>
      </Center>
      <ServiceDetails
        service={selectService}
        isOpen={detailsModal.isOpen}
        onClose={detailsModal.onClose}
        changeService={s => {
          const cloneServices = [...services];
          if (s) {
            cloneServices.splice(selectIndex, 1, s);
          } else {
            cloneServices.splice(selectIndex, 1);
          }
          setServices(cloneServices);
        }}
      />
    </Box>
  );
};

export default SSOPage;
