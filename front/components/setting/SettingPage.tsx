import {Center, Box, Heading, Tabs, TabList, Tab} from '@chakra-ui/react';
import {useRouter} from 'next/router';
import Setting from './Setting';

const SettingPage: React.FC<{index: number}> = ({index, children}) => {
  const router = useRouter();

  const handleChange = (index: number) => {
    let path = '/setting';
    switch (index) {
      case 0:
        path = '/setting';
        break;
      case 1:
        path = '/setting/mail';
        break;
      case 2:
        path = '/setting/password';
        break;
      default:
        break;
    }

    router.push(path);
  };

  return (
    <Center>
      <Box width={{base: '100%', lg: '75%'}} mt="2.5rem">
        <Center mb="1.7rem">
          <Heading>設定</Heading>
        </Center>
        <Box overflow="auto">
          <Box width={{base: '500px', sm: '100%'}}>
            <Tabs
              isFitted
              size="lg"
              defaultIndex={index}
              onChange={handleChange}
            >
              <TabList>
                <Tab>ユーザ情報</Tab>
                <Tab>メールアドレス</Tab>
                <Tab>パスワード</Tab>
              </TabList>
            </Tabs>
          </Box>
        </Box>

        {children}
      </Box>
    </Center>
  );
};

export default SettingPage;
