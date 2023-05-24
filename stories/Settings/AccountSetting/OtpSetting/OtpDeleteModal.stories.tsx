import type {Meta, StoryObj} from '@storybook/react';
import {OtpDeleteModal} from '../../../../components/Settings/AccountSetting/OtpSetting/OtpDeleteModal';

const meta: Meta<typeof OtpDeleteModal> = {
  title: 'CateiruSSO/Settings/AccountSetting/OtpSetting/OtpDeleteModal',
  component: OtpDeleteModal,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof OtpDeleteModal>;

export const Default: Story = {
  args: {
    isOpen: true,
  },
};
