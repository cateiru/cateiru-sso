import type {Meta, StoryObj} from '@storybook/react';
import {ProfileDatetime} from '../../components/Profile/ProfileDatetime';

const meta: Meta<typeof ProfileDatetime> = {
  title: 'CateiruSSO/Profile/ProfileDatetime',
  component: ProfileDatetime,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof ProfileDatetime>;

export const Default: Story = {};
