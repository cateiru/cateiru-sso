import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {ClientDetailContent} from '../../../components/Staff/Client/ClientDetailContent';
import {ClientAllowUserList} from '../../../utils/types/client';

const meta: Meta<typeof ClientDetailContent> = {
  title: 'CateiruSSO/Staff/Client/ClientDetailContent',
  component: ClientDetailContent,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof ClientDetailContent>;

const clientId = faker.string.uuid();

export const Default: Story = {
  args: {
    data: {
      client: {
        client_id: clientId,
        name: faker.company.name(),
        description: faker.lorem.paragraph(),
        image: faker.image.url(),
        org_id: faker.string.uuid(),
        org_member_only: faker.datatype.boolean(),
        is_allow: faker.datatype.boolean(),
        prompt: faker.helpers.arrayElement(['login', '2fa_login']),
        owner_user_id: faker.string.uuid(),
        client_secret: faker.string.alpha(),

        created_at: faker.date.past().toString(),
        updated_at: faker.date.past().toString(),
      },
      redirect_urls: Array(faker.number.int({min: 1, max: 5}))
        .fill('')
        .map(() => faker.internet.url()),
      referrer_urls: Array(faker.number.int({min: 1, max: 5}))
        .fill('')
        .map(() => faker.internet.domainName()),

      scopes: ['openid', 'profile', 'email'],

      allow_rules: [
        {
          id: 123,
          user: {
            id: faker.string.uuid(),
            user_name: faker.internet.userName(),
            avatar: faker.image.avatar(),
          },
          email_domain: null,
        },
        {
          id: 124,
          user: {
            id: faker.string.uuid(),
            user_name: faker.internet.userName(),
            avatar: faker.image.avatar(),
          },
          email_domain: null,
        },
        {
          id: 125,
          user: null,
          email_domain: faker.internet.domainName(),
        },
        {
          id: 126,
          user: null,
          email_domain: faker.internet.domainName(),
        },
        {
          id: 127,
          user: null,
          email_domain: faker.internet.domainName(),
        },
      ] as ClientAllowUserList,
    },
  },
};

export const Loading: Story = {};
