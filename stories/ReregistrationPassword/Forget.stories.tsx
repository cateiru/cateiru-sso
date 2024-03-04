import type {Meta, StoryObj} from '@storybook/react';
import {Forget} from '../../components/ReregistrationPassword/Forget';

const meta: Meta<typeof Forget> = {
  title: 'CateiruSSO/ReregistrationPassword/Forget',
  component: Forget,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof Forget>;

export const Default: Story = {};
