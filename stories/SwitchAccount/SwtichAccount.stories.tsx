import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {SwitchAccount} from '../../components/SwitchAccount/SwitchAccount';
import {api} from '../../utils/api';

const user = {
  id: '123',
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
};

const meta: Meta<typeof SwitchAccount> = {
  title: 'CateiruSSO/SwitchAccount/SwitchAccount',
  component: SwitchAccount,
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<typeof SwitchAccount>;

export const Default: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/v2/account/list'),
        method: 'GET',
        status: 200,
        response: [
          {
            id: user.id,
            user_name: user.user_name,
          },
          {
            id: faker.string.uuid(),
            user_name: faker.internet.userName(),
          },
          {
            id: faker.string.uuid(),
            user_name: faker.internet.userName(),
            avatar: faker.image.avatar(),
          },
        ],
      },
    ],
  },
};

export const Oauth: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/v2/account/list'),
        method: 'GET',
        status: 200,
        response: [
          {
            id: user.id,
            user_name: user.user_name,
          },
          {
            id: faker.string.uuid(),
            user_name: faker.internet.userName(),
          },
          {
            id: faker.string.uuid(),
            user_name: faker.internet.userName(),
            avatar: faker.image.avatar(),
          },
        ],
      },
    ],
    nextjs: {
      appDirectory: true,
      query: {
        oauth: '1',
      },
      navigation: {
        query: {
          oauth: '1',
        },
      },
    },
  },
};

export const ManyUser: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/v2/account/list'),
        method: 'GET',
        status: 200,
        response: [
          {
            id: user.id,
            user_name: user.user_name,
          },
          ...Array(faker.datatype.number({min: 10, max: 30}))
            .fill(0)
            .map(() => {
              return {
                id: faker.string.uuid(),
                user_name: faker.internet.userName(),
                avatar: faker.image.avatar(),
              };
            }),
        ],
      },
    ],
  },
};

export const Loading: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/v2/account/list'),
        method: 'GET',
        status: 200,
        delay: 10000,
        response: [
          {
            id: user.id,
            user_name: user.user_name,
          },
          {
            id: faker.string.uuid(),
            user_name: faker.internet.userName(),
          },
          {
            id: faker.string.uuid(),
            user_name: faker.internet.userName(),
            avatar: faker.image.avatar(),
          },
        ],
      },
    ],
  },
};
