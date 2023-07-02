import type {Meta, StoryObj} from '@storybook/react';
import {AddUser} from '../../../components/Staff/Brand/AddUser';

const meta: Meta<typeof AddUser> = {
  title: 'CateiruSSO/Staff/Brand/AddUser',
  component: AddUser,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof AddUser>;

export const Default: Story = {};
