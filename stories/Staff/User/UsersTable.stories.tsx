import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {UsersTable} from '../../../components/Staff/User/UsersTable';
import {api} from '../../../utils/api';
import {StaffUsers} from '../../../utils/types/staff';

const meta: Meta<typeof UsersTable> = {
  title: 'CateiruSSO/Staff/User/UsersTable',
  component: UsersTable,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof UsersTable>;

export const Default: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/admin/users'),
        method: 'GET',
        status: 200,
        response: [
          {
            id: faker.string.uuid(),
            user_name: faker.internet.userName(),
            email: faker.internet.email(),
            family_name: faker.person.lastName(),
            middle_name: null,
            given_name: faker.person.firstName(),
            gender: '1',
            birthdate: null,
            avatar: faker.image.avatar(),
            locale_id: 'ja_JP',

            created_at: faker.date.past().toISOString(),
            updated_at: faker.date.past().toISOString(),
          },
          {
            id: faker.string.uuid(),
            user_name: faker.internet.userName(),
            email: faker.internet.email(),
            family_name: faker.person.lastName(),
            middle_name: null,
            given_name: faker.person.firstName(),
            gender: '1',
            birthdate: null,
            avatar: faker.image.avatar(),
            locale_id: 'ja_JP',

            created_at: faker.date.past().toISOString(),
            updated_at: faker.date.past().toISOString(),
          },
        ] as StaffUsers,
      },
    ],
  },
};

export const Loading: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/admin/users'),
        method: 'GET',
        status: 200,
        delay: 10000,
        response: [
          {
            id: faker.string.uuid(),
            user_name: faker.internet.userName(),
            email: faker.internet.email(),
            family_name: faker.person.lastName(),
            middle_name: null,
            given_name: faker.person.firstName(),
            gender: '1',
            birthdate: null,
            avatar: faker.image.avatar(),
            locale_id: 'ja_JP',

            created_at: faker.date.past().toISOString(),
            updated_at: faker.date.past().toISOString(),
          },
        ] as StaffUsers,
      },
    ],
  },
};
