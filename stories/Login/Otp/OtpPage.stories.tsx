import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {OtpPage} from '../../../components/Login/OtpPage';

const meta: Meta<typeof OtpPage> = {
  title: 'CateiruSSO/Login/Otp/Page',
  component: OtpPage,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof OtpPage>;

export const Default: Story = {
  args: {
    loginUser: {
      avatar: null,
    },
    otpToken: '123456',
  },
};

export const Avatar: Story = {
  args: {
    loginUser: {
      avatar: faker.image.avatar(),
    },
    otpToken: '123456',
  },
};
