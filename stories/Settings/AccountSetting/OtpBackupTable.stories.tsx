import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {OtpBackupTable} from '../../../components/Settings/AccountSetting/OtpBackupTable';

const meta: Meta<typeof OtpBackupTable> = {
  title: 'CateiruSSO/Settings/AccountSetting/OtpBackupTable',
  component: OtpBackupTable,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof OtpBackupTable>;

export const Default: Story = {
  args: {
    backups: Array(10)
      .fill(0)
      .map(() => faker.string.alpha(15)),
  },
};
