import type {Meta, StoryObj} from '@storybook/react';
import {UserNameTable} from '../../../components/Staff/User/UserNameTable';

const meta: Meta<typeof UserNameTable> = {
  title: 'CateiruSSO/Staff/User/UserNameTable',
  component: UserNameTable,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof UserNameTable>;

export const Default: Story = {};
