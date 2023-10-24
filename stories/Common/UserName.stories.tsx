import type {Meta, StoryObj} from '@storybook/react';
import {UserName} from '../../components/Common/UserName';

const meta: Meta<typeof UserName> = {
  title: 'CateiruSSO/Common/UserName',
  component: UserName,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof UserName>;

export const Default: Story = {};
