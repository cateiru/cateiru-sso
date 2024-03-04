import type {Meta, StoryObj} from '@storybook/react';
import React from 'react';
import {FormProvider, useForm} from 'react-hook-form';
import {
  UserNameForm,
  UserNameFormData,
} from '../../../components/Common/Form/UserNameForm';

const Form: typeof UserNameForm = props => {
  const methods = useForm<UserNameFormData>();

  return (
    <FormProvider {...methods}>
      <UserNameForm {...props} />
    </FormProvider>
  );
};

const meta: Meta<typeof UserNameForm> = {
  title: 'CateiruSSO/Common/Form/UserNameForm',
  component: Form,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof UserNameForm>;

export const Default: Story = {};
