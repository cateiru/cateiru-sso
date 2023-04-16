import {PinInput, PinInputField} from '@chakra-ui/react';
import React from 'react';
import {config} from '../../utils/config';

export interface EmailForm {
  code: string;
}

interface Props {
  onSubmit: (data: EmailForm) => Promise<void>;
}

export const EmailVerifyForm: React.FC<Props> = ({onSubmit}) => {
  const [value, setValue] = React.useState('');
  const [isDisabled, setIsDisabled] = React.useState(false);
  const [rerender, setRerender] = React.useState(false);

  React.useEffect(() => {
    if (value.length >= config.registerAccountEmailCodeLength) {
      setIsDisabled(true);
      onSubmit({code: value})
        .then(() => {
          setIsDisabled(false);
        })
        .catch(() => {
          setIsDisabled(false);
          setValue('');

          // 無理やり再レンダリングする
          setRerender(true);
          setTimeout(() => {
            setRerender(false);
          }, 10);
        });
    }

    return () => setIsDisabled(false);
  }, [value]);

  return (
    <>
      {rerender || (
        <PinInput
          colorScheme="cateiru"
          isDisabled={isDisabled}
          otp
          autoFocus
          onChange={(v: string) => {
            setValue(v);
          }}
          value={value}
        >
          <PinInputField />
          <PinInputField ml=".5rem" />
          <PinInputField ml=".5rem" />
          <PinInputField ml=".5rem" />
          <PinInputField ml=".5rem" />
          <PinInputField ml=".5rem" />
        </PinInput>
      )}
    </>
  );
};
