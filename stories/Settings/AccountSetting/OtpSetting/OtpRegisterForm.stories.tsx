import type {Meta, StoryObj} from '@storybook/react';
import {OtpRegisterForm} from '../../../../components/Settings/AccountSetting/OtpSetting/OtpRegisterForm';

const meta: Meta<typeof OtpRegisterForm> = {
  title: 'CateiruSSO/Settings/AccountSetting/OtpSetting/OtpRegisterForm',
  component: OtpRegisterForm,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof OtpRegisterForm>;

export const Default: Story = {
  args: {
    onSubmit: async data => {
      window.alert(JSON.stringify(data));
    },
  },
};
