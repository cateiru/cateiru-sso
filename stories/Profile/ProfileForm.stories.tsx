import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {ProfileForm} from '../../components/Profile/ProfileForm';
import {UserState} from '../../utils/state/atom';
import {RecoilController} from '../RecoilController';

const meta: Meta<typeof ProfileForm> = {
  title: 'CateiruSSO/Profile/ProfileForm',
  component: () => {
    return (
      <RecoilController
        recoilState={UserState}
        defaultValue={undefined}
        setValues={[
          {
            key: 'no login',
            value: null,
          },
          {
            key: 'login',
            value: {
              user: {
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
          },
          {
            key: 'login no avatar',
            value: {
              user: {
                id: faker.datatype.uuid(),
                user_name: faker.internet.userName(),
                email: faker.internet.email(),
                family_name: faker.name.lastName(),
                middle_name: faker.name.middleName(),
                given_name: faker.name.firstName(),
                gender: '1',
                birthdate: null,
                avatar: null,
                locale_id: 'ja_JP',

                created: faker.date.past().toString(),
                modified: faker.date.past().toString(),
              },
            },
          },
        ]}
      >
        <ProfileForm />
      </RecoilController>
    );
  },
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof ProfileForm>;

export const Default: Story = {};
