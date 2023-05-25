import type {Meta, StoryObj} from '@storybook/react';
import React from 'react';
import {FormProvider, useForm} from 'react-hook-form';
import {
  RegisterPasswordFormContext,
  type RegisterPasswordFormContextData,
} from '../../../components/Common/Form/RegisterPasswordFormContext';

const Form = () => {
  const methods = useForm<RegisterPasswordFormContextData>();
  const [ok, setOk] = React.useState(false);

  return (
    <FormProvider {...methods}>
      <RegisterPasswordFormContext ok={ok} setOk={setOk} />
    </FormProvider>
  );
};

const meta: Meta<typeof RegisterPasswordFormContext> = {
  title: 'CateiruSSO/Common/Form/RegisterPasswordFormContext',
  component: Form,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof RegisterPasswordFormContext>;

export const Default: Story = {};
