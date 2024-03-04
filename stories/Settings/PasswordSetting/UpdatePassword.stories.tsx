import type {Meta, StoryObj} from '@storybook/react';
import {UpdatePassword} from '../../../components/Settings/PasswordSetting/UpdatePassword';

const meta: Meta<typeof UpdatePassword> = {
  title: 'CateiruSSO/Settings/PasswordSetting/UpdatePassword',
  component: UpdatePassword,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof UpdatePassword>;

export const Default: Story = {};
