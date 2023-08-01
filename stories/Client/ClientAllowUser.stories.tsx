import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {ClientAllowUser} from '../../components/Client/ClientAllowUser';
import {api} from '../../utils/api';
import {ClientAllowUserList} from '../../utils/types/client';

const meta: Meta<typeof ClientAllowUser> = {
  title: 'CateiruSSO/Client/ClientAllowUser',
  component: ClientAllowUser,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof ClientAllowUser>;

const clientId = faker.string.uuid();
const param = new URLSearchParams({client_id: clientId});

export const Default: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/v2/client/allow_user', param),
        method: 'GET',
        status: 200,
        delay: 1000,
        response: [
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
        url: api('/v2/client/allow_user', param),
        method: 'GET',
        status: 200,
        delay: 10000,
        response: [],
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
