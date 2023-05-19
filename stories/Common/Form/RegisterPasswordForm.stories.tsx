import type {Meta, StoryObj} from '@storybook/react';
import React from 'react';
import {FormProvider, useForm} from 'react-hook-form';
import {
  RegisterPasswordForm,
  type RegisterPasswordFormData,
} from '../../../components/Common/Form/RegisterPasswordForm';

const Form = () => {
  const methods = useForm<RegisterPasswordFormData>();
  const [ok, setOk] = React.useState(false);

  return (
    <FormProvider {...methods}>
      <RegisterPasswordForm ok={ok} setOk={setOk} />
    </FormProvider>
  );
};

const meta: Meta<typeof RegisterPasswordForm> = {
  title: 'CateiruSSO/Common/Form/RegisterPasswordForm',
  component: Form,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof RegisterPasswordForm>;

export const Default: Story = {};
