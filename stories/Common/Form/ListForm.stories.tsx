import type {Meta, StoryObj} from '@storybook/react';
import React from 'react';
import {FormProvider, useForm} from 'react-hook-form';
import {ListForm} from '../../../components/Common/Form/ListForm';

interface Sample {
  sample: string[];
}

const Form = () => {
  const methods = useForm<Sample>({
    defaultValues: {
      sample: ['default1'],
    },
  });

  return (
    <FormProvider {...methods}>
      <ListForm name="sample" />
    </FormProvider>
  );
};

const meta: Meta<typeof ListForm> = {
  title: 'CateiruSSO/Common/Form/ListForm',
  component: Form,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof ListForm>;

export const Default: Story = {};
