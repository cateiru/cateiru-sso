import type {Meta, StoryObj} from '@storybook/react';
import {AgoTime} from '../../../components/Common/Time';

const meta: Meta<typeof AgoTime> = {
  title: 'CateiruSSO/Common/Time/AgoTime',
  component: AgoTime,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof AgoTime>;

export const Default: Story = {
  args: {
    time: new Date(),
  },
};

export const String: Story = {
  args: {
    time: new Date().toISOString(),
  },
};

export const InvalidTimeString: Story = {
  args: {
    time: 'invalid time string',
  },
};
