import type {Meta, StoryObj} from '@storybook/react';
import {AccountSetting} from '../../../components/Settings/AccountSetting/AccountSetting';

const meta: Meta<typeof AccountSetting> = {
  title: 'CateiruSSO/Settings/AccountSetting/AccountSetting',
  component: AccountSetting,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof AccountSetting>;

export const Default: Story = {};
