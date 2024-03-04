import type {Meta, StoryObj} from '@storybook/react';
import React from 'react';
import {FormProvider, useForm} from 'react-hook-form';
import {
  RedirectUrlsForm,
  RedirectUrlsFormValue,
} from '../../../components/Common/Form/RedirectUrlsForm';

const Form: typeof RedirectUrlsForm = props => {
  const methods = useForm<RedirectUrlsFormValue>({
    defaultValues: {
      redirectUrls: [{value: ''}],
    },
  });

  return (
    <FormProvider {...methods}>
      <RedirectUrlsForm {...props} />
    </FormProvider>
  );
};

const meta: Meta<typeof RedirectUrlsForm> = {
  title: 'CateiruSSO/Common/Form/RedirectUrlsForm',
  component: Form,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof RedirectUrlsForm>;

export const Default: Story = {
  args: {
    maxCreatedCount: 5,
  },
};
