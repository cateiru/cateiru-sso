import {Flex} from '@chakra-ui/react';
import React from 'react';
import SelectCreate from './SelectCreate';

const CreateAccountPage: React.FC = () => {
  return (
    <Flex
      justifyContent="center"
      alignItems="center"
      flexDirection="column"
      width="100%"
      height="80vh"
      px={{base: '1rem', md: '5rem'}}
    >
      <SelectCreate />
    </Flex>
  );
};

export default CreateAccountPage;
