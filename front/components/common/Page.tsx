import {Flex, Box} from '@chakra-ui/react';
import React from 'react';
import Footer from './Footer';

const Page: React.FC = ({children}) => {
  return (
    <Flex flexDirection="column" minHeight="100vh">
      <Box>
        {/* <Header /> */}
        {children}
      </Box>
      <Box marginTop="auto">
        <Footer />
      </Box>
    </Flex>
  );
};

export default Page;
