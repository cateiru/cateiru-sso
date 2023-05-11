import type {Meta, StoryObj} from '@storybook/react';
import {UserIDEmailForm} from '../../../components/Login/UserIDEmailForm';

const meta: Meta<typeof UserIDEmailForm> = {
  title: 'CateiruSSO/Login/UserIDEmail/Form',
  component: UserIDEmailForm,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof UserIDEmailForm>;

export const Default: Story = {
  args: {
    onSubmit: async data => {
      window.alert(JSON.stringify(data));
    },
  },
};
