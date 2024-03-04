import type {Meta, StoryObj} from '@storybook/react';
import React from 'react';
import {FormProvider, useForm} from 'react-hook-form';
import {
  ImageForm,
  ImageFormValue,
} from '../../../components/Common/Form/ImageForm';

const Form = () => {
  const methods = useForm<ImageFormValue>();

  return (
    <FormProvider {...methods}>
      <ImageForm />
    </FormProvider>
  );
};

const meta: Meta<typeof ImageForm> = {
  title: 'CateiruSSO/Common/Form/ImageForm',
  component: Form,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof ImageForm>;

export const Default: Story = {};
