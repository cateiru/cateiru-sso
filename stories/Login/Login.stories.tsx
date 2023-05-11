import type {Meta, StoryObj} from '@storybook/react';
import {Login} from '../../components/Login/Login';

const meta: Meta<typeof Login> = {
  title: 'CateiruSSO/Login/Login',
  component: Login,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof Login>;

export const Default: Story = {};
