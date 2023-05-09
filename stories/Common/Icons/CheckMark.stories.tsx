import type {Meta, StoryObj} from '@storybook/react';
import {CheckMark} from '../../../components/Common/Icons/CheckMark';

const meta: Meta<typeof CheckMark> = {
  title: 'CateiruSSO/Common/Icons/CheckMark',
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
