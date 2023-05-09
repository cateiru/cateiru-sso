import {Box, Button, Center} from '@chakra-ui/react';
import React from 'react';
import {type RecoilState, useRecoilState} from 'recoil';

interface Props<T> {
  children: React.ReactNode;
  recoilState: RecoilState<T>;
  defaultValue: T;
  setValues: {key: string; value: T}[];
}

export const RecoilController = <T extends any>(props: Props<T>) => {
  const [state, setState] = useRecoilState(props.recoilState);

  const onDefault = () => {
    setState(props.defaultValue);
  };

  return (
    <>
      {props.children}
      <Box>
        <Center>
          <Button onClick={onDefault}>Set Default</Button>
          {props.setValues.map(value => {
            return (
              <Button
                onClick={() => {
                  setState(value.value);
                }}
                ml=".5rem"
              >
                Set {value.key}
              </Button>
            );
          })}
        </Center>
        <Box as="pre">{JSON.stringify(state, null, '\t')}</Box>
      </Box>
    </>
  );
};
