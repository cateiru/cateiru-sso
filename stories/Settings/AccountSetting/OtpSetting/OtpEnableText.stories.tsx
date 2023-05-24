import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {OtpEnableText} from '../../../../components/Settings/AccountSetting/OtpSetting/OtpEnableText';

const meta: Meta<typeof OtpEnableText> = {
  title: 'CateiruSSO/Settings/AccountSetting/OtpSetting/OtpEnableText',
  component: OtpEnableText,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof OtpEnableText>;

export const Default: Story = {
  args: {
    modified: faker.date.recent(),
  },
};
