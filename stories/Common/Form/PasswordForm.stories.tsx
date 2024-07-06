import type {Meta, StoryObj} from '@storybook/react';
import React from 'react';
import {FormProvider, useForm} from 'react-hook-form';
import {
  PasswordForm,
  type PasswordFormData,
} from '../../../components/Common/Form/PasswordForm';

const Form: typeof PasswordForm = props => {
  const methods = useForm<PasswordFormData>();

  return (
    <FormProvider {...methods}>
      <PasswordForm {...props} />
    </FormProvider>
  );
};

const meta: Meta<typeof PasswordForm> = {
  title: 'CateiruSSO/Common/Form/PasswordForm',
  component: Form,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof PasswordForm>;

export const Default: Story = {};
