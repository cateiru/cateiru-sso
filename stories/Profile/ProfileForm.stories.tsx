import type {Meta, StoryObj} from '@storybook/react';
import {ProfileForm} from '../../components/Profile/ProfileForm';

const meta: Meta<typeof ProfileForm> = {
  title: 'CateiruSSO/Profile/ProfileForm',
  component: ProfileForm,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof ProfileForm>;

export const Default: Story = {};
