import {Text, Button} from '@chakra-ui/react';
import React from 'react';
import {aa, defaultTable} from '../../utils/keygenAsciiArt';

const KeyAA = () => {
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
    <>
      <Text fontSize=".8rem">{key}</Text>
      <Text
        as="pre"
        fontSize="1.5rem"
        fontFamily="Source Code Pro"
        lineHeight="1.45rem"
      >
        <code>{AA}</code>
      </Text>
      <Button size="sm" onClick={() => setReset(e => !e)}>
        リセット
      </Button>
    </>
  );
};

export default KeyAA;
