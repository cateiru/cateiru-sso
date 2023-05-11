import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {SelectCertification} from '../../components/Login/SelectCertification';

const meta: Meta<typeof SelectCertification> = {
  title: 'CateiruSSO/Login/SelectCertification',
  component: SelectCertification,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof SelectCertification>;

export const Default: Story = {
  args: {
    loginUser: {
      avatar: null,
      user_name: faker.internet.userName(),

      // このコンポーネントはどちらもtrueになる
      available_password: true,
      available_passkey: true,
    },
    setStep: step => window.alert(step),
  },
};

export const Avatar: Story = {
  args: {
    loginUser: {
      avatar: faker.image.avatar(),
      user_name: faker.internet.userName(),

      // このコンポーネントはどちらもtrueになる
      available_password: true,
      available_passkey: true,
    },
    setStep: step => window.alert(step),
  },
};
