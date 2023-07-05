import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {OrgDetailContent} from '../../../components/Staff/Org/OrgDetailContent';

const meta: Meta<typeof OrgDetailContent> = {
  title: 'CateiruSSO/Staff/Org/OrgDetailContent',
  component: OrgDetailContent,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof OrgDetailContent>;

export const Default: Story = {
  args: {
    org: {
      id: faker.string.uuid(),
      name: faker.company.name(),
      image: faker.image.url(),
      link: faker.internet.url(),

      created_at: faker.date.past().toISOString(),
      updated_at: faker.date.past().toISOString(),
    },
    users: [
      {
        id: 1234,
        user: {
          id: faker.string.uuid(),
          user_name: faker.internet.userName(),
          avatar: faker.image.avatar(),
        },
        role: 'admin',

        created_at: faker.date.past().toISOString(),
        updated_at: faker.date.past().toISOString(),
      },
      ...Array(faker.number.int({min: 1, max: 10}))
        .fill(0)
        .map((_, i) => {
          return {
            id: i,
            user: {
              id: faker.string.uuid(),
              user_name: faker.internet.userName(),
              avatar: faker.image.avatar(),
            },
            role: 'member',

            created_at: faker.date.past().toISOString(),
            updated_at: faker.date.past().toISOString(),
          };
        }),
      {
        id: 1235,
        user: {
          id: faker.string.uuid(),
          user_name: faker.internet.userName(),
          avatar: faker.image.avatar(),
        },
        role: 'guest',

        created_at: faker.date.past().toISOString(),
        updated_at: faker.date.past().toISOString(),
      },
    ],
  },
};
