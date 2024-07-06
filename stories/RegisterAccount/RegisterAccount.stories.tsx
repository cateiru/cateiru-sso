import type {Meta, StoryObj} from '@storybook/react';
import {RegisterAccount} from '../../components/RegisterAccount/RegisterAccount';

const meta: Meta<typeof RegisterAccount> = {
  title: 'CateiruSSO/RegisterAccount/RegisterAccount',
  component: RegisterAccount,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof RegisterAccount>;

export const Default: Story = {};
