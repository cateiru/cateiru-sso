import type {Meta, StoryObj} from '@storybook/react';
import {OtpRegister} from '../../../../components/Settings/AccountSetting/OtpSetting/OtpRegister';

const meta: Meta<typeof OtpRegister> = {
  title: 'CateiruSSO/Settings/AccountSetting/OtpSetting/OtpRegister',
  component: OtpRegister,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof OtpRegister>;

export const Default: Story = {
  args: {
    isOpen: true,
  },
};
