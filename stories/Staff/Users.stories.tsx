import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {Users} from '../../components/Staff/Users';
import {api} from '../../utils/api';
import {StaffUsers} from '../../utils/types/staff';

const meta: Meta<typeof Users> = {
  title: 'CateiruSSO/Staff/Users',
  component: Users,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof Users>;

export const Default: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/v2/admin/users'),
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
        url: api('/v2/admin/users'),
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
