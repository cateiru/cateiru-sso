import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {SendMainSuccess} from '../../components/ReregistrationPassword/SendMailSuccess';

const meta: Meta<typeof SendMainSuccess> = {
  title: 'CateiruSSO/ReregistrationPassword/SendMainSuccess',
  component: SendMainSuccess,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof SendMainSuccess>;

export const Default: Story = {
  args: {
    email: faker.internet.email(),
  },
};
