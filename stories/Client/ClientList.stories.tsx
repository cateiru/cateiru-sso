import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {ClientList} from '../../components/Client/ClientList';

const meta: Meta<typeof ClientList> = {
  title: 'CateiruSSO/Client/ClientList',
  component: ClientList,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof ClientList>;

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
