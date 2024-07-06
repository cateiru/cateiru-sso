import type {Meta, StoryObj} from '@storybook/react';
import {Top} from '../../components/Top/Top';

const meta: Meta<typeof Top> = {
  title: 'CateiruSSO/Top/Top',
  component: Top,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof Top>;

export const Default: Story = {};
