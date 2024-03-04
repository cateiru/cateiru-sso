import type {Meta, StoryObj} from '@storybook/react';
import {NotFound} from '../../../components/Common/Error/NotFound';

const meta: Meta<typeof NotFound> = {
  title: 'CateiruSSO/Common/Error/NotFound',
  component: NotFound,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof NotFound>;

export const Default: Story = {};
