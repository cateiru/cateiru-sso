import type {Meta, StoryObj} from '@storybook/react';
import {PasswordForm} from '../../../components/Login/PasswordForm';

const meta: Meta<typeof PasswordForm> = {
  title: 'CateiruSSO/Login/Password/Form',
  component: PasswordForm,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof PasswordForm>;

export const Default: Story = {
  args: {
    onSubmit: async data => {
      window.alert(JSON.stringify(data));
    },
  },
};
