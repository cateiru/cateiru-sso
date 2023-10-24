import type {Meta, StoryObj} from '@storybook/react';
import {Profile} from '../../components/Profile/Profile';

const meta: Meta<typeof Profile> = {
  title: 'CateiruSSO/Profile/Profile',
  component: Profile,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof Profile>;

export const Default: Story = {};
