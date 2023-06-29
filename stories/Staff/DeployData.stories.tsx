import type {Meta, StoryObj} from '@storybook/react';
import {DeployData} from '../../components/Staff/DeployData';

const meta: Meta<typeof DeployData> = {
  title: 'CateiruSSO/Staff/DeployData',
  component: DeployData,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof DeployData>;

export const Default: Story = {};
