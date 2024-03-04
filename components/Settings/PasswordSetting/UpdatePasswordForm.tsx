import {Button} from '@chakra-ui/react';
import React from 'react';
import {FormProvider, useForm} from 'react-hook-form';
import {
  PasswordForm,
  type PasswordFormData,
} from '../../Common/Form/PasswordForm';
import {
  type RegisterPasswordFormContextData,
  RegisterPasswordFormContext,
} from '../../Common/Form/RegisterPasswordFormContext';

export interface UpdatePasswordFormData
  extends PasswordFormData,
    RegisterPasswordFormContextData {}

interface Props {
  onSubmit: (data: UpdatePasswordFormData) => Promise<void>;
}

export const UpdatePasswordForm: React.FC<Props> = props => {
  const methods = useForm<UpdatePasswordFormData>();
  const {
    handleSubmit,
    clearErrors,
    reset,
    formState: {isSubmitting},
  } = methods;

  const [ok, setOk] = React.useState(false);

  const onSubmit = async (data: UpdatePasswordFormData) => {
    if (!ok) {
      return;
    } else {
      clearErrors('new_password');
    }

    await props.onSubmit(data);
    reset();
  };

  return (
    <FormProvider {...methods}>
      <form onSubmit={handleSubmit(onSubmit)}>
        <PasswordForm
          enableWebauthn={false}
          label="現在設定しているパスワード"
        />
        <RegisterPasswordFormContext
          ok={ok}
          setOk={setOk}
          label="新しいパスワード"
          mt="1rem"
        />
        <Button
          mt="1rem"
          isLoading={isSubmitting}
          colorScheme="cateiru"
          type="submit"
          w="100%"
        >
          パスワードを更新する
        </Button>
      </form>
    </FormProvider>
  );
};
