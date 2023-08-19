import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {RegisterSessionTable} from '../../../components/Staff/Session/RegisterSessionTable';
import {api} from '../../../utils/api';
import {RegisterSessions} from '../../../utils/types/staff';

const meta: Meta<typeof RegisterSessionTable> = {
  title: 'CateiruSSO/Staff/Session/RegisterSessionTable',
  component: RegisterSessionTable,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof RegisterSessionTable>;

export const Default: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/v2/admin/register_session'),
        method: 'GET',
        status: 200,
        response: [
          {
            email: faker.internet.email(),
            email_verified: faker.datatype.boolean(),
            send_count: faker.number.int({min: 0, max: 10}),
            retry_count: faker.number.int({min: 0, max: 10}),
            org_id: null,

            period: faker.date.future().toISOString(),

            created_at: faker.date.past().toISOString(),
            updated_at: faker.date.past().toISOString(),
          },
          {
            email: faker.internet.email(),
            email_verified: faker.datatype.boolean(),
            send_count: faker.number.int({min: 0, max: 10}),
            retry_count: faker.number.int({min: 0, max: 10}),
            org_id: null,

            period: faker.date.future().toISOString(),

            created_at: faker.date.past().toISOString(),
            updated_at: faker.date.past().toISOString(),
          },
          {
            email: faker.internet.email(),
            email_verified: faker.datatype.boolean(),
            send_count: faker.number.int({min: 0, max: 10}),
            retry_count: faker.number.int({min: 0, max: 10}),
            org_id: null,

            period: faker.date.future().toISOString(),

            created_at: faker.date.past().toISOString(),
            updated_at: faker.date.past().toISOString(),
          },
          {
            email: faker.internet.email(),
            email_verified: faker.datatype.boolean(),
            send_count: faker.number.int({min: 0, max: 10}),
            retry_count: faker.number.int({min: 0, max: 10}),
            org_id: null,

            period: faker.date.future().toISOString(),

            created_at: faker.date.past().toISOString(),
            updated_at: faker.date.past().toISOString(),
          },
          {
            email: faker.internet.email(),
            email_verified: faker.datatype.boolean(),
            send_count: faker.number.int({min: 0, max: 10}),
            retry_count: faker.number.int({min: 0, max: 10}),
            org_id: faker.string.uuid(),

            period: faker.date.future().toISOString(),

            created_at: faker.date.past().toISOString(),
            updated_at: faker.date.past().toISOString(),
          },
        ] as RegisterSessions,
      },
    ],
  },
};

export const Loading: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/v2/admin/register_session'),
        method: 'GET',
        status: 200,
        delay: 10000,
        response: [],
      },
    ],
  },
};
