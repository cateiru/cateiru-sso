import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {LoginTryTable} from '../../components/Histories/LoginTryTable';
import {api} from '../../utils/api';
import {LoginTryHistoryList} from '../../utils/types/history';

const meta: Meta<typeof LoginTryTable> = {
  title: 'CateiruSSO/Histories/LoginTryTable',
  component: LoginTryTable,
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<typeof LoginTryTable>;

export const Default: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/history/try_login'),
        method: 'GET',
        status: 200,
        response: [
          {
            id: 1,
            device: '',
            os: 'Windows',
            browser: 'Firefox',
            is_mobile: false,
            ip: faker.internet.ip(),
            created_at: faker.date.past().toISOString(),

            identifier: 0,
          },
          {
            id: 2,
            device: '',
            os: 'Windows',
            browser: 'Google Chrome',
            is_mobile: false,
            ip: faker.internet.ip(),
            created_at: faker.date.past().toISOString(),

            identifier: 0,
          },
          {
            id: 3,
            device: 'iPhone',
            os: 'iOS',
            browser: 'Safari',
            is_mobile: true,
            ip: faker.internet.ip(),
            created_at: faker.date.past().toISOString(),

            identifier: 1,
          },
        ] as LoginTryHistoryList,
      },
    ],
  },
};

export const TooMany: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/history/try_login'),
        method: 'GET',
        status: 200,
        response: Array(100)
          .fill(0)
          .map((_, i) => ({
            id: i,
            device: '',
            os: 'Windows',
            browser: 'Firefox',
            is_mobile: false,
            ip: faker.internet.ip(),
            created_at: faker.date.past().toISOString(),

            identifier: 0,
          })) as LoginTryHistoryList,
      },
    ],
  },
};

export const Loading: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/history/try_login'),
        method: 'GET',
        status: 200,
        delay: 10000,
        response: [
          {
            id: 1,
            device: '',
            os: 'Windows',
            browser: 'Firefox',
            is_mobile: false,
            ip: faker.internet.ip(),
            created_at: faker.date.past().toISOString(),

            identifier: 0,
          },
        ] as LoginTryHistoryList,
      },
    ],
  },
};
