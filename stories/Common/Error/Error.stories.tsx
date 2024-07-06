import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {Error} from '../../../components/Common/Error/Error';

const meta: Meta<typeof Error> = {
  title: 'CateiruSSO/Common/Error/Error',
  component: Error,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof Error>;

export const Default: Story = {
  args: {
    message: faker.lorem.sentence(),
  },
};

export const Unique: Story = {
  args: {
    message: faker.lorem.sentence(),
    unique_code: 4,
  },
};
