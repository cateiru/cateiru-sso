import type {Meta, StoryObj} from '@storybook/react';
import {EmailVerifyForm} from '../../../components/Common/Form/EmailVerifyForm';

const sleep = (msec: number) =>
  new Promise(resolve => setTimeout(resolve, msec));

const meta: Meta<typeof EmailVerifyForm> = {
  title: 'CateiruSSO/Common/Form/EmailVerifyForm',
  component: EmailVerifyForm,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof EmailVerifyForm>;

export const Default: Story = {
  args: {
    onSubmit: async data => {
      console.log('start');
      await sleep(1000);

      window.alert(JSON.stringify(data));
    },
    emailCodeLength: 6,
  },
};

export const NoVerify: Story = {
  args: {
    onSubmit: async () => {
      console.log('start');
      await sleep(1000);

      throw new Error('error');
    },
    emailCodeLength: 6,
  },
};
