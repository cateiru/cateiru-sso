import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {OtpBackups} from '../../../components/Settings/AccountSetting/OtpBackups';

const meta: Meta<typeof OtpBackups> = {
  title: 'CateiruSSO/Settings/AccountSetting/OtpBackups',
  component: OtpBackups,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof OtpBackups>;

export const Default: Story = {
  args: {
    backups: Array(10)
      .fill(0)
      .map(() => faker.string.alpha(15)),
    title: 'バックアップ',
  },
};
