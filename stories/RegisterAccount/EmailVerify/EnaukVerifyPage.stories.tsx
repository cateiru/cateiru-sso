import type {Meta, StoryObj} from '@storybook/react';
import {EmailVerifyPage} from '../../../components/RegisterAccount/EmailVerifyPage';

const meta: Meta<typeof EmailVerifyPage> = {
  title: 'CateiruSSO/RegisterAccount/EmailVerify/Page',
  component: EmailVerifyPage,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof EmailVerifyPage>;

export const Default: Story = {
  args: {
    registerToken: 'token',
    setStatus: status => {
      console.log(status);
    },
    reset: () => {},
  },
};

export const NoVerify: Story = {
  args: {
    registerToken: 'token',
    setStatus: status => {
      console.log(status);
    },
    reset: () => {},
  },
};
