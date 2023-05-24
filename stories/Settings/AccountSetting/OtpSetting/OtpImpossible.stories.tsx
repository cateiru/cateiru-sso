import type {Meta, StoryObj} from '@storybook/react';
import {OtpImpossible} from '../../../../components/Settings/AccountSetting/OtpSetting/OtpImpossible';

const meta: Meta<typeof OtpImpossible> = {
  title: 'CateiruSSO/Settings/AccountSetting/OtpSetting/OtpImpossible',
  component: OtpImpossible,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof OtpImpossible>;

export const Default: Story = {};
