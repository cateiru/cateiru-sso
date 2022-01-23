import {Flex} from '@chakra-ui/react';
import {useRouter} from 'next/router';
import React from 'react';
import SelectCreate, {CreateType} from './SelectCreate';

const CreateAccountPage: React.FC = () => {
  const [mailToken, setMailToken] = React.useState('');
  const [selectType, setSelectType] = React.useState(CreateType.Initialize);

  const router = useRouter();
  React.useEffect(() => {
    if (!router.isReady) return;
    const query = router.query;

    if (typeof query['token'] === 'string') {
      setMailToken(query['token']);
      setSelectType(CreateType.ValidateMail);
    }
  }, [router.isReady, router.query]);

  return (
    <Flex
      justifyContent="center"
      alignItems="center"
      flexDirection="column"
      width="100%"
      height="80vh"
      px={{base: '1rem', md: '5rem'}}
    >
      <SelectCreate
        selectType={selectType}
        mailToken={mailToken}
        setSelectType={setSelectType}
      />
    </Flex>
  );
};

export default CreateAccountPage;
