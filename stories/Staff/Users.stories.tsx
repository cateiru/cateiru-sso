import type {Meta, StoryObj} from '@storybook/react';
import {Users} from '../../components/Staff/Users';

const meta: Meta<typeof Users> = {
  title: 'CateiruSSO/Staff/Users',
  component: Users,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof Users>;

export const Default: Story = {};
