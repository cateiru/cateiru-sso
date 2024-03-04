import type {Meta, StoryObj} from '@storybook/react';
import {RegisterPassword} from '../../../components/Settings/PasswordSetting/RegisterPassword';

const meta: Meta<typeof RegisterPassword> = {
  title: 'CateiruSSO/Settings/PasswordSetting/RegisterPassword',
  component: RegisterPassword,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof RegisterPassword>;

export const Default: Story = {};
