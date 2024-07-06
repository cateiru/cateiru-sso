import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {EmailView} from '../../components/Profile/EmailView';

const meta: Meta<typeof EmailView> = {
  title: 'CateiruSSO/Profile/EmailView',
  component: EmailView,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof EmailView>;

export const Default: Story = {
  args: {
    email: faker.internet.email(),
  },
};
