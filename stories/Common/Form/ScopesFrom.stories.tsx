import type {Meta, StoryObj} from '@storybook/react';
import React from 'react';
import {FormProvider, useForm} from 'react-hook-form';

import {
  ScopesForm,
  ScopesFormValue,
} from '../../../components/Common/Form/ScopesFrom';

const Form: typeof ScopesForm = props => {
  const methods = useForm<ScopesFormValue>({
    defaultValues: {
      scopes: [
        {
          value: 'openid',
          isRequired: true,
        },
      ],
    },
  });

  return (
    <FormProvider {...methods}>
      <ScopesForm {...props} />
    </FormProvider>
  );
};

const meta: Meta<typeof ScopesForm> = {
  title: 'CateiruSSO/Common/Form/ScopesForm',
  component: Form,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof ScopesForm>;

export const Default: Story = {
  args: {
    scopes: [
      'openid',
      'profile',
      'email',
      'phone',
      'address',
      'offline_access',
    ],
  },
};
