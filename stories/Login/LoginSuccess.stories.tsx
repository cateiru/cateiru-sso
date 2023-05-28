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
      id: faker.string.uuid(),
      user_name: faker.internet.userName(),
      email: faker.internet.email(),
      family_name: faker.person.lastName(),
      middle_name: null,
      given_name: faker.person.firstName(),
      gender: '1',
      birthdate: null,
      avatar: faker.image.avatar(),
      locale_id: 'ja_JP',

      created_at: faker.date.past().toString(),
      updated_at: faker.date.past().toString(),
    },
  },
};
