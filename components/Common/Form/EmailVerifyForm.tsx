import {PinInput, PinInputField} from '@chakra-ui/react';
import React from 'react';

export interface EmailVerifyForm {
  code: string;
}

interface Props {
  emailCodeLength: number;
  onSubmit: (data: EmailVerifyForm) => Promise<void>;
}

export const EmailVerifyForm: React.FC<Props> = ({
  onSubmit,
  emailCodeLength,
}) => {
  const [value, setValue] = React.useState('');
  const [isDisabled, setIsDisabled] = React.useState(false);
  const [rerender, setRerender] = React.useState(false);

  React.useEffect(() => {
    if (value.length >= emailCodeLength) {
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
          {new Array(emailCodeLength).fill(0).map((_, i) => {
            return (
              <PinInputField
                ml={i !== 0 ? '.5rem' : undefined}
                key={`pin_input-${i}`}
              />
            );
          })}
        </PinInput>
      )}
    </>
  );
};
