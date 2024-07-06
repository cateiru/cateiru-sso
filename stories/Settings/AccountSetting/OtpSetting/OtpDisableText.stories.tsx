import type {Meta, StoryObj} from '@storybook/react';
import {OtpDisableText} from '../../../../components/Settings/AccountSetting/OtpSetting/OtpDisableText';

const meta: Meta<typeof OtpDisableText> = {
  title: 'CateiruSSO/Settings/AccountSetting/OtpSetting/OtpDisableText',
  component: OtpDisableText,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof OtpDisableText>;

export const Default: Story = {};
