import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {ClientsTable} from '../../../components/Staff/Client/ClientsTable';
import {api} from '../../../utils/api';
import {StaffClients} from '../../../utils/types/staff';

const meta: Meta<typeof ClientsTable> = {
  title: 'CateiruSSO/Staff/Client/ClientsTable',
  component: ClientsTable,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof ClientsTable>;

export const Default: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/admin/clients'),
        method: 'GET',
        status: 200,
        response: [
          {
            client_id: faker.string.uuid(),
            name: faker.company.name(),
            image: faker.image.url(),
          },
          {
            client_id: faker.string.uuid(),
            name: faker.company.name(),
            image: faker.image.url(),
          },
          {
            client_id: faker.string.uuid(),
            name: faker.company.name(),
            image: faker.image.url(),
          },
        ] as StaffClients,
      },
    ],
  },
};

export const Loading: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/admin/clients'),
        method: 'GET',
        status: 200,
        delay: 10000,
        response: [],
      },
    ],
  },
};
