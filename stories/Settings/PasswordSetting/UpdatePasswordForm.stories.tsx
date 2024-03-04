import type {Meta, StoryObj} from '@storybook/react';
import {UpdatePasswordForm} from '../../../components/Settings/PasswordSetting/UpdatePasswordForm';

const meta: Meta<typeof UpdatePasswordForm> = {
  title: 'CateiruSSO/Settings/PasswordSetting/UpdatePasswordForm',
  component: UpdatePasswordForm,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof UpdatePasswordForm>;

export const Default: Story = {
  args: {
    onSubmit: async data => {
      window.alert(JSON.stringify(data));
    },
  },
};
