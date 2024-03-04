import type {Meta, StoryObj} from '@storybook/react';
import {OtpForm} from '../../../components/Login/OtpForm';

const meta: Meta<typeof OtpForm> = {
  title: 'CateiruSSO/Login/Otp/Form',
  component: OtpForm,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof OtpForm>;

export const Default: Story = {
  args: {
    onSubmit: async data => {
      window.alert(data.otp);
    },
  },
};
