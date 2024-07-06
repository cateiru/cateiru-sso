import {Button, FormControl, FormErrorMessage, Input} from '@chakra-ui/react';
import React from 'react';
import {useForm} from 'react-hook-form';

export interface OtpFormData {
  otp: string;
}

interface Props {
  onSubmit: (data: OtpFormData, reset: () => void) => Promise<void>;
}

export const OtpForm: React.FC<Props> = props => {
  const submitRef = React.useRef<HTMLButtonElement>(null);
  const [otpForm, setOtpForm] = React.useState('');

  const {
    handleSubmit,
    register,
    reset,
    formState: {errors, isSubmitting},
  } = useForm<OtpFormData>();

  React.useEffect(() => {
    if (/[0-9]{6}/.test(otpForm)) {
      setTimeout(() => {
        submitRef.current?.click();
      }, 10);
    }
  }, [otpForm]);

  const onSubmitWrapper = async (data: OtpFormData) => {
    await props.onSubmit(data, reset);
  };

  return (
    <form onSubmit={handleSubmit(onSubmitWrapper)}>
      <FormControl isInvalid={!!errors.otp}>
        <Input
          id="otp"
          type="text"
          autoComplete="one-time-code"
          {...register('otp', {
            required: 'この値は必須です',
            onChange: e => setOtpForm(e.target.value),
          })}
        />
        <FormErrorMessage>{errors.otp && errors.otp.message}</FormErrorMessage>
      </FormControl>
      <Button
        mt="1rem"
        isLoading={isSubmitting}
        colorScheme="cateiru"
        type="submit"
        w="100%"
        ref={submitRef}
      >
        ログイン
      </Button>
    </form>
  );
};
