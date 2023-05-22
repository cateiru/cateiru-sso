import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {OtpSetting} from '../../../components/Settings/AccountSetting/OtpSetting';
import {api} from '../../../utils/api';
import {UserOtp} from '../../../utils/types/user';

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
        url: api('/v2/user/otp'),
        method: 'GET',
        status: 200,
        response: {
          enable: true,
          modified: faker.date.recent().toISOString(),
        } as UserOtp,
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
          enable: false,
          modified: null,
        } as UserOtp,
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
          enable: true,
          modified: faker.date.recent().toISOString(),
        } as UserOtp,
      },
    ],
  },
};
