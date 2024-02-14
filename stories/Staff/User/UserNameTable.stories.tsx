import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {UserNameTable} from '../../../components/Staff/User/UserNameTable';
import {api} from '../../../utils/api';
import {UserNames} from '../../../utils/types/staff';

const meta: Meta<typeof UserNameTable> = {
  title: 'CateiruSSO/Staff/User/UserNameTable',
  component: UserNameTable,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof UserNameTable>;

export const Default: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/admin/user_name'),
        method: 'GET',
        status: 200,
        response: [
          {
            id: 1,
            user_name: faker.internet.userName(),
            user_id: faker.string.uuid(),

            period: faker.date.future().toISOString(),

            created_at: faker.date.past().toISOString(),
            updated_at: faker.date.past().toISOString(),
          },
          {
            id: 2,
            user_name: faker.internet.userName(),
            user_id: faker.string.uuid(),

            period: faker.date.future().toISOString(),

            created_at: faker.date.past().toISOString(),
            updated_at: faker.date.past().toISOString(),
          },
        ] as UserNames,
      },
    ],
  },
};
