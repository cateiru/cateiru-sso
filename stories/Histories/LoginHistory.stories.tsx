import type {Meta, StoryObj} from '@storybook/react';
import {LoginHistory} from '../../components/Histories/LoginHistory';

const meta: Meta<typeof LoginHistory> = {
  title: 'CateiruSSO/Histories/LoginHistory',
  component: LoginHistory,
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<typeof LoginHistory>;

export const Default: Story = {};
