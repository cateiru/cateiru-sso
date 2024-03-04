import type {Meta, StoryObj} from '@storybook/react';
import {UpdateAvatar} from '../../components/Profile/UpdateAvatar';

const meta: Meta<typeof UpdateAvatar> = {
  title: 'CateiruSSO/Profile/UpdateAvatar',
  component: UpdateAvatar,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof UpdateAvatar>;

export const Default: Story = {};
