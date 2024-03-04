import type {Meta, StoryObj} from '@storybook/react';
import {Spinner} from '../../../components/Common/Icons/Spinner';

const meta: Meta<typeof Spinner> = {
  title: 'CateiruSSO/Common/Icons/Spinner',
  component: Spinner,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof Spinner>;

export const Default: Story = {};
