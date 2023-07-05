import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {OrgsTable} from '../../../components/Staff/Org/OrgsTable';
import {api} from '../../../utils/api';
import {Organizations} from '../../../utils/types/staff';

const meta: Meta<typeof OrgsTable> = {
  title: 'CateiruSSO/Staff/Org/OrgsTable',
  component: OrgsTable,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof OrgsTable>;

export const Default: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/v2/admin/orgs'),
        method: 'GET',
        status: 200,
        response: [
          {
            id: faker.string.uuid(),
            name: faker.company.name(),
            image: faker.image.url(),
            link: faker.internet.url(),

            created_at: faker.date.past().toISOString(),
            updated_at: faker.date.past().toISOString(),
          },
          {
            id: faker.string.uuid(),
            name: faker.company.name(),
            image: faker.image.url(),
            link: faker.internet.url(),

            created_at: faker.date.past().toISOString(),
            updated_at: faker.date.past().toISOString(),
          },
          {
            id: faker.string.uuid(),
            name: faker.company.name(),
            image: faker.image.url(),
            link: faker.internet.url(),

            created_at: faker.date.past().toISOString(),
            updated_at: faker.date.past().toISOString(),
          },
          {
            id: faker.string.uuid(),
            name: faker.company.name(),
            image: faker.image.url(),
            link: faker.internet.url(),

            created_at: faker.date.past().toISOString(),
            updated_at: faker.date.past().toISOString(),
          },
        ] as Organizations,
      },
    ],
  },
};

export const Loading: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/v2/admin/users'),
        method: 'GET',
        status: 200,
        delay: 10000,
        response: [] as Organizations,
      },
    ],
  },
};
