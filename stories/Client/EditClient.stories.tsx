import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {EditClient} from '../../components/Client/EditClient';
import {api} from '../../utils/api';
import {
  ClientConfig,
  ClientDetail as ClientDetailType,
} from '../../utils/types/client';

const meta: Meta<typeof EditClient> = {
  title: 'CateiruSSO/Client/EditClient',
  component: EditClient,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof EditClient>;

const clientId = faker.string.uuid();
const param = new URLSearchParams({client_id: clientId});

export const Org: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/client', param),
        method: 'GET',
        status: 200,
        delay: 1000,
        response: {
          client_id: faker.string.uuid(),

          name: faker.company.name(),
          description: faker.company.catchPhrase(),
          image: faker.image.url(),

          is_allow: faker.datatype.boolean(),
          prompt: null,

          org_member_only: faker.datatype.boolean(),

          created_at: faker.date.past().toISOString(),
          updated_at: faker.date.past().toISOString(),

          client_secret: faker.string.uuid(),
          redirect_urls: [faker.internet.url()],
          referrer_urls: [new URL(faker.internet.url()).host],
          scopes: ['openid', 'profile', 'email'],
          org_id: faker.string.uuid(),
        } as ClientDetailType,
      },
      {
        url: api('/client/config'),
        method: 'GET',
        status: 200,
        delay: 1000,
        response: {
          redirect_url_max: 10,
          referrer_url_max: 10,
          scopes: ['profile', 'email', 'openid'],
        } as ClientConfig,
      },
    ],
    nextjs: {
      appDirectory: true,
      navigation: {
        segments: [['id', clientId]],
      },
    },
  },
};

export const NoOrg: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/client', param),
        method: 'GET',
        status: 200,
        delay: 1000,
        response: {
          client_id: faker.string.uuid(),

          name: faker.company.name(),
          description: faker.company.catchPhrase(),
          image: faker.image.url(),

          is_allow: faker.datatype.boolean(),
          prompt: null,

          created_at: faker.date.past().toISOString(),
          updated_at: faker.date.past().toISOString(),

          client_secret: faker.string.uuid(),
          redirect_urls: [faker.internet.url()],
          referrer_urls: [new URL(faker.internet.url()).host],
          scopes: ['openid', 'profile', 'email'],
          org_id: null,
        } as ClientDetailType,
      },
      {
        url: api('/client/config'),
        method: 'GET',
        status: 200,
        delay: 1000,
        response: {
          redirect_url_max: 10,
          referrer_url_max: 10,
          scopes: ['profile', 'email', 'openid'],
        } as ClientConfig,
      },
    ],
    nextjs: {
      appDirectory: true,
      navigation: {
        segments: [['id', clientId]],
      },
    },
  },
};

export const Loading: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/client', param),
        method: 'GET',
        status: 200,
        delay: 10000,
        response: {},
      },
      {
        url: api('/client/config'),
        method: 'GET',
        status: 200,
        delay: 1000,
        response: {
          redirect_url_max: 10,
          referrer_url_max: 10,
          scopes: ['profile', 'email', 'openid'],
        } as ClientConfig,
      },
    ],
    nextjs: {
      appDirectory: true,
      navigation: {
        segments: [['id', clientId]],
      },
    },
  },
};
