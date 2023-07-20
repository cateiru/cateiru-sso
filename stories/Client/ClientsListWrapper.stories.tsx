import type {Meta, StoryObj} from '@storybook/react';
import {ClientsListWrapper} from '../../components/Client/ClientsListWrapper';

const meta: Meta<typeof ClientsListWrapper> = {
  title: 'CateiruSSO/Client/ClientsListWrapper',
  component: ClientsListWrapper,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof ClientsListWrapper>;

export const Default: Story = {};
