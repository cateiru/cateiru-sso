import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {OtpRegisterReadQR} from '../../../../components/Settings/AccountSetting/OtpSetting/OtpRegisterReadQR';

const meta: Meta<typeof OtpRegisterReadQR> = {
  title: 'CateiruSSO/Settings/AccountSetting/OtpSetting/OtpRegisterReadQR',
  component: OtpRegisterReadQR,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof OtpRegisterReadQR>;

export const Default: Story = {
  args: {
    token: faker.internet.url(),
  },
};
