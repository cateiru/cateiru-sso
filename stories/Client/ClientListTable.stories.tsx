import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {ClientListTable} from '../../components/Client/ClientListTable';

const meta: Meta<typeof ClientListTable> = {
  title: 'CateiruSSO/Client/ClientListTable',
  component: ClientListTable,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof ClientListTable>;

export const Default: Story = {
  args: {
    clients: Array(faker.number.int({min: 1, max: 10}))
      .fill(0)
      .map(() => ({
        client_id: faker.string.uuid(),
        name: faker.company.name(),
        description: faker.lorem.sentence(),
        image: faker.image.url(),
        is_allow: false,
        prompt: 'openid email',
        org_member_only: false,

        created_at: faker.date.past().toISOString(),
        updated_at: faker.date.past().toISOString(),
      })),
  },
};

export const Loading: Story = {
  args: {},
};
