import type {Meta, StoryObj} from '@storybook/react';
import {ReregistrationPasswordForm} from '../../components/ReregistrationPassword/ReregistrationPasswordForm';

const meta: Meta<typeof ReregistrationPasswordForm> = {
  title: 'CateiruSSO/ReregistrationPassword/ReregistrationPasswordForm',
  component: ReregistrationPasswordForm,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof ReregistrationPasswordForm>;

export const Default: Story = {
  args: {
    onSubmit: async data => {
      window.alert(data.new_password);
    },
  },
};
