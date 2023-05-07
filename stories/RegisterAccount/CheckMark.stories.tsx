import type {Meta, StoryObj} from '@storybook/react';
import {CheckMark} from '../../components/RegisterAccount/CheckMark';

const meta: Meta<typeof CheckMark> = {
  title: 'CateiruSSO/RegisterAccount/CheckMark',
  component: CheckMark,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof CheckMark>;

export const Default: Story = {
  args: {
    size: 100,
    bgColor: '#7ac142',
    color: '#fff',
  },
};
