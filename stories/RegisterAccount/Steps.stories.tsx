import type {Meta, StoryObj} from '@storybook/react';
import {Steps} from '../../components/RegisterAccount/Steps';

const meta: Meta<typeof Steps> = {
  title: 'CateiruSSO/RegisterAccount/Steps',
  component: Steps,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof Steps>;

export const Default: Story = {
  args: {
    activeStep: 0,
  },
};
