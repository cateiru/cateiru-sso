import type {Meta, StoryObj} from '@storybook/react';
import {PasswordPage} from '../../../components/Login/PasswordPage';

const meta: Meta<typeof PasswordPage> = {
  title: 'CateiruSSO/Login/Password/Page',
  component: PasswordPage,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof PasswordPage>;

export const Default: Story = {};
