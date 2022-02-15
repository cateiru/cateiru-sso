import {
  Box,
  Heading,
  Center,
  SimpleGrid,
  useColorMode,
  useToast,
  Avatar,
  Text,
} from '@chakra-ui/react';
import React from 'react';
import {getUserSSO, ServiceLogInfo} from '../../utils/api/userSSO';
import {hawManyDaysAgo} from '../../utils/date';

const ConnectedAccountPage = () => {
  const {colorMode} = useColorMode();
  const [services, setServices] = React.useState<ServiceLogInfo[]>([]);
  const toast = useToast();

  React.useEffect(() => {
    const f = async () => {
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
    };

    f();
  }, []);

  const ServiceElement = (s: ServiceLogInfo) => {
    const latestDate = s.logs.reduce((p, v) => {
      if (Date.parse(p.accept_date) - Date.parse(v.accept_date) > 0) {
        return p;
      }
      return v;
    });

    return (
      <Box
        key={s.client_id}
        width="100%"
        maxWidth="500px"
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
      >
        <Avatar mt="1.2rem" ml="2rem" name={s.name} src={s.service_icon} />
        <Heading mt=".5rem" ml="2rem">
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
          {services.map(ServiceElement)}
        </SimpleGrid>
      </Center>
    </Box>
  );
};

export default ConnectedAccountPage;
