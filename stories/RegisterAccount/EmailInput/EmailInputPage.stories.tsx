import type {Meta, StoryObj} from '@storybook/react';
import {EmailInputPage} from '../../../components/RegisterAccount/EmailInputPage';
import type {StepStatus} from '../../../components/RegisterAccount/Steps';

const meta: Meta<typeof EmailInputPage> = {
  title: 'CateiruSSO/RegisterAccount/EmailInput/Page',
  component: EmailInputPage,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};
export default meta;

type Story = StoryObj<typeof EmailInputPage>;

export const Default: Story = {
  args: {
    nextStep: () => {
      window.alert('next step');
    },
    prevStep: () => {
      window.alert('prev step');
    },
    setStatus: (status: StepStatus) => {
      window.alert(status);
    },
    setRegisterToken: (token: string) => {
      console.log(token);
    },
  },
};
