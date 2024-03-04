import type {Meta, StoryObj} from '@storybook/react';
import {KeyGen} from '../../../components/Common/Animation/KeyGen';

const meta: Meta<typeof KeyGen> = {
  title: 'CateiruSSO/Common/Animation/KeyGen',
  component: KeyGen,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof Error>;

export const Default: Story = {};
