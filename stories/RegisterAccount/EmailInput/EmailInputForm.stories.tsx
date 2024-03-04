import type {Meta, StoryObj} from '@storybook/react';
import {EmailInputForm} from '../../../components/RegisterAccount/EmailInputForm';

const meta: Meta<typeof EmailInputForm> = {
  title: 'CateiruSSO/RegisterAccount/EmailInput/Form',
  component: EmailInputForm,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof EmailInputForm>;

export const Default: Story = {
  args: {
    onSubmit: async data => {
      window.alert(JSON.stringify(data));
    },
  },
};
