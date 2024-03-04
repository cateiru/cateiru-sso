import {Button} from '@chakra-ui/react';
import React from 'react';
import {FormProvider, useForm} from 'react-hook-form';
import {
  type RegisterPasswordFormContextData,
  RegisterPasswordFormContext,
} from './RegisterPasswordFormContext';

interface Props {
  buttonText: string;
  onSubmit: (data: RegisterPasswordFormContextData) => Promise<void>;
}

export const RegisterPasswordForm: React.FC<Props> = props => {
  const methods = useForm<RegisterPasswordFormContextData>();
  const {
    handleSubmit,
    clearErrors,
    formState: {isSubmitting},
  } = methods;

  const [ok, setOk] = React.useState(false);

  const onSubmit = async (data: RegisterPasswordFormContextData) => {
    if (!ok) {
      return;
    } else {
      clearErrors('new_password');
    }

    await props.onSubmit(data);
  };

  return (
    <FormProvider {...methods}>
      <form onSubmit={handleSubmit(onSubmit)}>
        <RegisterPasswordFormContext ok={ok} setOk={setOk} />
        <Button
          mt="1rem"
          isLoading={isSubmitting}
          colorScheme="cateiru"
          type="submit"
          w="100%"
        >
          {props.buttonText}
        </Button>
      </form>
    </FormProvider>
  );
};
