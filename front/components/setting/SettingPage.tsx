import {Center, Box, Heading, Tabs, TabList, Tab} from '@chakra-ui/react';
import {useRouter} from 'next/router';

const SettingPage: React.FC<{index: number; children: React.ReactNode}> = ({
  index,
  children,
}) => {
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
      case 3:
        path = '/setting/account';
        break;
      default:
        break;
    }

    router.push(path);
  };

  return (
    <Center>
      <Box width={{base: '100%', lg: '1000px'}} mt="2.5rem">
        <Center mb="1.7rem">
          <Heading>設定</Heading>
        </Center>
        <Box
          overflow="auto"
          css={{
            '&::-webkit-scrollbar': {display: 'none'},
            scrollbarWidth: 'none',
          }}
        >
          <Box width={{base: '800px', md: '100%'}}>
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
                <Tab>アカウント関連</Tab>
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
