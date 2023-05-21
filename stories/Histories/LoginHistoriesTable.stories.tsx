import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {LoginHistoriesTable} from '../../components/Histories/LoginHistoriesTable';
import {api} from '../../utils/api';
import {type LoginDeviceList} from '../../utils/types/history';

const meta: Meta<typeof LoginHistoriesTable> = {
  title: 'CateiruSSO/Histories/LoginHistoriesTable',
  component: LoginHistoriesTable,
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<typeof LoginHistoriesTable>;

export const Default: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/v2/history/login'),
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
            created: faker.date.past().toISOString(),
            is_current: false,
          },
          {
            id: 2,
            device: '',
            os: 'Windows',
            browser: 'Google Chrome',
            is_mobile: false,
            ip: faker.internet.ip(),
            created: faker.date.past().toISOString(),
            is_current: false,
          },
          {
            id: 3,
            device: 'iPhone',
            os: 'iOS',
            browser: 'Safari',
            is_mobile: true,
            ip: faker.internet.ip(),
            created: faker.date.past().toISOString(),
            is_current: false,
          },
        ] as LoginDeviceList,
      },
    ],
  },
};

export const TooMany: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/v2/history/login'),
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
            created: faker.date.past().toISOString(),
            is_current: false,
          })) as LoginDeviceList,
      },
    ],
  },
};

export const Loading: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/v2/history/login'),
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
            created: faker.date.past().toISOString(),
            is_current: true,
          },
        ] as LoginDeviceList,
      },
    ],
  },
};
