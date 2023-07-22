import type {Meta, StoryObj} from '@storybook/react';
import {RegisterClient} from '../../components/Client/RegisterClient';

const meta: Meta<typeof RegisterClient> = {
  title: 'CateiruSSO/Client/RegisterClient',
  component: RegisterClient,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof RegisterClient>;

export const Default: Story = {};
