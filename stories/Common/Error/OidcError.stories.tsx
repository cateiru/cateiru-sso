import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {OidcError} from '../../../components/Common/Error/Error';

const meta: Meta<typeof OidcError> = {
  title: 'CateiruSSO/Common/Error/OidcError',
  component: OidcError,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof OidcError>;

export const Default: Story = {
  args: {
    error: faker.lorem.sentence(),
    error_description: faker.lorem.sentence(),

    error_uri: faker.internet.url(),
    state: 'openid profile',
  },
};

export const NoDescription: Story = {
  args: {
    error: faker.lorem.sentence(),
  },
};

export const NoUri: Story = {
  args: {
    error: faker.lorem.sentence(),
    error_description: faker.lorem.sentence() + faker.internet.url(),

    state: 'openid profile',
  },
};

export const NoState: Story = {
  args: {
    error: faker.lorem.sentence(),
    error_description: faker.lorem.sentence(),

    error_uri: faker.internet.url(),
  },
};
