import type {Meta, StoryObj} from '@storybook/react';
import {Staff} from '../../components/Staff/Staff';

const meta: Meta<typeof Staff> = {
  title: 'CateiruSSO/Staff/Staff',
  component: Staff,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof Staff>;

export const Default: Story = {};
