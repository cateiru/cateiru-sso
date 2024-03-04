import type {Meta, StoryObj} from '@storybook/react';
import {LoginHistoryHeader} from '../../components/Histories/LoginHistoryHeader';

const meta: Meta<typeof LoginHistoryHeader> = {
  title: 'CateiruSSO/Histories/LoginHistoryHeader',
  component: LoginHistoryHeader,
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<typeof LoginHistoryHeader>;

export const Default: Story = {};
