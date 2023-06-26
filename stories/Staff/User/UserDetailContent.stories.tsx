import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {UserDetailContent} from '../../../components/Staff/User/UserDetailContent';

const meta: Meta<typeof UserDetailContent> = {
  title: 'CateiruSSO/Staff/User/UserDetailContent',
  component: UserDetailContent,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof UserDetailContent>;

const userId = faker.string.uuid();

export const Default: Story = {
  args: {
    data: {
      user: {
        id: userId,
        user_name: faker.internet.userName(),
        email: faker.internet.email(),
        family_name: faker.person.lastName(),
        middle_name: null,
        given_name: faker.person.firstName(),
        gender: '1',
        birthdate: null,
        avatar: faker.image.avatar(),
        locale_id: 'ja_JP',

        created_at: faker.date.past().toString(),
        updated_at: faker.date.past().toString(),
      },

      staff: {
        user_id: userId,
        memo: faker.lorem.paragraph(),

        created_at: faker.date.past().toString(),
        updated_at: faker.date.past().toString(),
      },

      user_brands: Array(faker.number.int({min: 1, max: 5}))
        .fill(0)
        .map(() => ({
          id: faker.string.uuid(),
          brand_id: faker.string.uuid(),
          brand_name: faker.internet.displayName(),

          created_at: faker.date.past().toString(),
        })),

      clients: Array(faker.number.int({min: 1, max: 5}))
        .fill(0)
        .map(() => ({
          client_id: faker.string.uuid(),
          name: faker.internet.domainName(),
          image: faker.internet.avatar(),
        })),
    },
  },
};

export const Loading: Story = {};
