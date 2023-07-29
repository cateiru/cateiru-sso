import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {ClientDetail} from '../../components/Client/ClientDetail';
import {api} from '../../utils/api';
import {ClientDetail as ClientDetailType} from '../../utils/types/client';

const meta: Meta<typeof ClientDetail> = {
  title: 'CateiruSSO/Client/ClientDetail',
  component: ClientDetail,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof ClientDetail>;

const clientId = faker.string.uuid();
const param = new URLSearchParams({client_id: clientId});

export const Default: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/v2/client', param),
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
        } as ClientDetailType,
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
        url: api('/v2/client', param),
        method: 'GET',
        status: 200,
        delay: 10000,
        response: {},
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
