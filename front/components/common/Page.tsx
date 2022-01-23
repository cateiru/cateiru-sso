import {Flex, Box} from '@chakra-ui/react';
import React from 'react';
import Footer from './Footer';
import Header from './Header';

const Page: React.FC = props => {
  return (
    <Flex flexDirection="column" minHeight="100vh">
      <Box>
        <Header />
        {props.children}
      </Box>
      <Box marginTop="auto">
        <Footer />
      </Box>
    </Flex>
  );
};

export default Page;
