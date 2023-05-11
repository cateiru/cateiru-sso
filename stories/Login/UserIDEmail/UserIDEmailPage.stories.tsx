import type {Meta, StoryObj} from '@storybook/react';
import {UserIDEmailPage} from '../../../components/Login/UserIDEmailPage';

const meta: Meta<typeof UserIDEmailPage> = {
  title: 'CateiruSSO/Login/UserIDEmail/Page',
  component: UserIDEmailPage,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof UserIDEmailPage>;

export const Default: Story = {};
