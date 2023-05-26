import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {WebAuthnDevices} from '../../../../components/Settings/AccountSetting/WebAuthn/WebAuthnDevices';
import {api} from '../../../../utils/api';
import {AccountWebAuthnDevices} from '../../../../utils/types/account';

const meta: Meta<typeof WebAuthnDevices> = {
  title: 'CateiruSSO/Settings/AccountSetting/WebAuthn/WebAuthnDevices',
  component: WebAuthnDevices,
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<typeof WebAuthnDevices>;

export const Default: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/v2/account/webauthn'),
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
          },
          {
            id: 2,
            device: '',
            os: 'Windows',
            browser: 'Google Chrome',
            is_mobile: false,
            ip: faker.internet.ip(),
            created: faker.date.past().toISOString(),
          },
          {
            id: 3,
            device: 'iPhone',
            os: 'iOS',
            browser: 'Safari',
            is_mobile: true,
            ip: faker.internet.ip(),
            created: faker.date.past().toISOString(),
          },
        ] as AccountWebAuthnDevices,
      },
    ],
  },
};

export const TooMany: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/v2/account/webauthn'),
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
          })) as AccountWebAuthnDevices,
      },
    ],
  },
};

export const Loading: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/v2/account/webauthn'),
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
          },
        ] as AccountWebAuthnDevices,
      },
    ],
  },
};
