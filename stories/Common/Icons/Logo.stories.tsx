import type {Meta, StoryObj} from '@storybook/react';
import {Logo} from '../../../components/Common/Icons/Logo';

const meta: Meta<typeof Logo> = {
  title: 'CateiruSSO/Common/Icons/Logo',
  component: Logo,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof Logo>;

export const Default: Story = {};
