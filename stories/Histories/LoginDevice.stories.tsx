import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {LoginDevice} from '../../components/Histories/LoginDevice';
import {api} from '../../utils/api';
import {type LoginDeviceList} from '../../utils/types/history';

const meta: Meta<typeof LoginDevice> = {
  title: 'CateiruSSO/Histories/LoginDevice',
  component: LoginDevice,
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<typeof LoginDevice>;

export const Default: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/history/login_devices'),
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
            is_current: false,
          },
          {
            id: 2,
            device: '',
            os: 'Windows',
            browser: 'Google Chrome',
            is_mobile: false,
            ip: faker.internet.ip(),
            created_at: faker.date.past().toISOString(),
            is_current: false,
          },
          {
            id: 3,
            device: 'iPhone',
            os: 'iOS',
            browser: 'Safari',
            is_mobile: true,
            ip: faker.internet.ip(),
            created_at: faker.date.past().toISOString(),
            is_current: true,
          },
        ] as LoginDeviceList,
      },
    ],
  },
};

export const Loading: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/history/login_devices'),
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
            is_current: true,
          },
        ] as LoginDeviceList,
      },
    ],
  },
};
