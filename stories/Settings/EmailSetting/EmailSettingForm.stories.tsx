import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {EmailSettingForm} from '../../../components/Settings/EmailSetting/EmailSettingForm';

const meta: Meta<typeof EmailSettingForm> = {
  title: 'CateiruSSO/Settings/EmailSettingForm',
  component: EmailSettingForm,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof EmailSettingForm>;

export const Default: Story = {
  args: {
    disabled: false,
    email: faker.internet.email(),
    onSubmit: async data => {
      window.alert(JSON.stringify(data));
    },
  },
};

export const Disabled: Story = {
  args: {
    disabled: true,
    email: faker.internet.email(),
    onSubmit: async data => {
      window.alert(JSON.stringify(data));
    },
  },
};
