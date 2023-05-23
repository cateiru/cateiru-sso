import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {OtpSetting} from '../../../components/Settings/AccountSetting/OtpSetting';
import {api} from '../../../utils/api';
import {AccountCertificates} from '../../../utils/types/account';

const meta: Meta<typeof OtpSetting> = {
  title: 'CateiruSSO/Settings/AccountSetting/OtpSetting',
  component: OtpSetting,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof OtpSetting>;

export const Enable: Story = {
  parameters: {
    mockData: [
      {
        url: api('/v2/account/certificates'),
        method: 'GET',
        status: 200,
        response: {
          password: true,
          otp: true,
          otp_modified: faker.date.recent().toISOString(),
        } as AccountCertificates,
      },
    ],
  },
};

export const Disable: Story = {
  parameters: {
    mockData: [
      {
        url: api('/v2/user/otp'),
        method: 'GET',
        status: 200,
        response: {
          password: true,
          otp: false,
          otp_modified: null,
        } as AccountCertificates,
      },
    ],
  },
};

export const Impossible: Story = {
  parameters: {
    mockData: [
      {
        url: api('/v2/user/otp'),
        method: 'GET',
        status: 200,
        response: {
          password: false,
          otp: false,
          otp_modified: null,
        } as AccountCertificates,
      },
    ],
  },
};

export const Loading: Story = {
  parameters: {
    mockData: [
      {
        url: api('/v2/user/otp'),
        method: 'GET',
        status: 200,
        delay: 10000,
        response: {
          password: true,
          otp: true,
          otp_modified: faker.date.recent().toISOString(),
        } as AccountCertificates,
      },
    ],
  },
};
