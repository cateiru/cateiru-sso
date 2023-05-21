import type {Meta, StoryObj} from '@storybook/react';
import {EmailSettingVerifyForm} from '../../components/Settings/EmailSettingVerifyForm';

const meta: Meta<typeof EmailSettingVerifyForm> = {
  title: 'CateiruSSO/Settings/EmailSettingVerifyForm',
  component: EmailSettingVerifyForm,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof EmailSettingVerifyForm>;

export const Default: Story = {};
