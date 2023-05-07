import type {Meta, StoryObj} from '@storybook/react';
import {RegisterPasswordForm} from '../../../components/RegisterAccount/RegisterPasswordForm';

const sleep = (msec: number) =>
  new Promise(resolve => setTimeout(resolve, msec));

const meta: Meta<typeof RegisterPasswordForm> = {
  title: 'CateiruSSO/RegisterAccount/RegisterCertificate/RegisterPasswordForm',
  component: RegisterPasswordForm,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof RegisterPasswordForm>;

export const Default: Story = {
  args: {
    onSubmit: async () => {
      await sleep(1000);
      console.log('submit');
    },
  },
};
