import {faker} from '@faker-js/faker';
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

export const Default: Story = {
  args: {
    loginUser: {
      user_name: 'test',
      avatar: null,
      available_passkey: false,
      available_password: true,
    },
  },
};

export const Avatar: Story = {
  args: {
    loginUser: {
      user_name: 'test',
      avatar: faker.image.avatar(),
      available_passkey: false,
      available_password: true,
    },
  },
};