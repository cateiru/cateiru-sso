import type {Meta, StoryObj} from '@storybook/react';
import {RegisterAccountPage} from '../../components/RegisterAccount/RegisterAccountPage';

const meta: Meta<typeof RegisterAccountPage> = {
  title: 'CateiruSSO/RegisterAccount/RegisterAccountPage',
  component: RegisterAccountPage,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof RegisterAccountPage>;

export const Default: Story = {};
