import type {Meta, StoryObj} from '@storybook/react';
import {CompleteRegisterPage} from '../../components/RegisterAccount/CompleteRegisterPage';

const meta: Meta<typeof CompleteRegisterPage> = {
  title: 'CateiruSSO/RegisterAccount/CompleteRegisterPage',
  component: CompleteRegisterPage,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof CompleteRegisterPage>;

export const Default: Story = {};
