import type {Meta, StoryObj} from '@storybook/react';
import {OtpBackupModal} from '../../../components/Settings/AccountSetting/OtpBackupModal';

const meta: Meta<typeof OtpBackupModal> = {
  title: 'CateiruSSO/Settings/AccountSetting/OtpBackupModal',
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
