import type {Meta, StoryObj} from '@storybook/react';
import {OtpBackupModal} from '../../../../components/Settings/AccountSetting/OtpSetting/OtpBackupModal';

const meta: Meta<typeof OtpBackupModal> = {
  title: 'CateiruSSO/Settings/AccountSetting/OtpSetting/OtpBackupModal',
  component: OtpBackupModal,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof OtpBackupModal>;

export const Default: Story = {
  args: {
    isOpen: true,
  },
};
