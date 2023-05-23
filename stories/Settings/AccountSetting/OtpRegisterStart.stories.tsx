import type {Meta, StoryObj} from '@storybook/react';
import {OtpRegisterStart} from '../../../components/Settings/AccountSetting/OtpRegisterStart';

const meta: Meta<typeof OtpRegisterStart> = {
  title: 'CateiruSSO/Settings/AccountSetting/OtpRegisterStart',
  component: OtpRegisterStart,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof OtpRegisterStart>;

export const Default: Story = {
  args: {
    onRegisterStart: () => {
      window.alert('onRegisterStart');
    },
  },
};
