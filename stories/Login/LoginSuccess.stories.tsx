import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {LoginSuccess} from '../../components/Login/LoginSuccess';

const meta: Meta<typeof LoginSuccess> = {
  title: 'CateiruSSO/Login/LoginSuccess',
  component: LoginSuccess,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof LoginSuccess>;

export const Default: Story = {
  args: {
    loggedInUser: {
      id: faker.datatype.uuid(),
      user_name: faker.internet.userName(),
      email: faker.internet.email(),
      family_name: faker.name.lastName(),
      middle_name: null,
      given_name: faker.name.firstName(),
      gender: '1',
      birthdate: null,
      avatar: faker.image.avatar(),
      locale_id: 'ja_JP',

      created: faker.date.past().toString(),
      modified: faker.date.past().toString(),
    },
  },
};
