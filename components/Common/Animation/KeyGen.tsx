import {Box, Button, Text} from '@chakra-ui/react';
import React from 'react';
import {aa, defaultTable} from '../../../utils/keygenAsciiArt';

export const KeyGen = () => {
  const [AA, setAA] = React.useState<string>(defaultTable());
  const [reset, setReset] = React.useState(false);
  const [key, setKey] = React.useState('');

  React.useEffect(() => {
    const [keyGen, key] = aa();
    setKey(key);

    const generator = keyGen.run();
    const interval = setInterval(() => {
      const iter = generator.next();
      if (iter.done) {
        return;
      }
      setAA(iter.value);
    }, 30);
    return () => clearInterval(interval);
  }, [reset]);

  return (
    <Box>
      <Text fontSize=".8rem" textAlign="center">
        {key}
      </Text>
      <Text
        textAlign="center"
        as="pre"
        fontSize="1.5rem"
        fontFamily="Source Code Pro"
        lineHeight="1.45rem"
      >
        <code>{AA}</code>
      </Text>
      <Button
        size="sm"
        onClick={() => setReset(e => !e)}
        m="auto"
        colorScheme="cateiru"
        display="block"
      >
        リセット
      </Button>
    </Box>
  );
};
