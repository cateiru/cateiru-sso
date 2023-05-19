import {Button} from '@chakra-ui/react';
import React from 'react';
import {FormProvider, useForm} from 'react-hook-form';
import {
  RegisterPasswordForm as RegisterPasswordFormBase,
  type RegisterPasswordFormData,
} from '../../components/Common/Form/RegisterPasswordForm';

interface Props {
  onSubmit: (data: RegisterPasswordFormData) => Promise<void>;
}

export const RegisterPasswordForm: React.FC<Props> = props => {
  const methods = useForm<RegisterPasswordFormData>();
  const {
    handleSubmit,
    clearErrors,
    formState: {isSubmitting},
  } = methods;

  const [ok, setOk] = React.useState(false);

  const onSubmit = async (data: RegisterPasswordFormData) => {
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
        <RegisterPasswordFormBase ok={ok} setOk={setOk} />
        <Button
          mt="1rem"
          isLoading={isSubmitting}
          colorScheme="cateiru"
          type="submit"
          w="100%"
        >
          パスワードを登録する
        </Button>
      </form>
    </FormProvider>
  );
};
