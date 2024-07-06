import type {Meta, StoryObj} from '@storybook/react';
import {ForgetForm} from '../../components/ReregistrationPassword/ForgetForm';

const meta: Meta<typeof ForgetForm> = {
  title: 'CateiruSSO/ReregistrationPassword/ForgetForm',
  component: ForgetForm,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof ForgetForm>;

export const Default: Story = {
  args: {
    onSubmit: async data => {
      window.alert(data.email);
    },
  },
};
