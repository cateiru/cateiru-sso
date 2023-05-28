import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {ProfileDatetime} from '../../components/Profile/ProfileDatetime';
import {UserState} from '../../utils/state/atom';
import {RecoilController} from '../RecoilController';

const meta: Meta<typeof ProfileDatetime> = {
  title: 'CateiruSSO/Profile/ProfileDatetime',
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
              is_staff: false,
            },
          },
          {
            key: 'staff',
            value: {
              user: {
                id: faker.string.uuid(),
                user_name: faker.internet.userName(),
                email: faker.internet.email(),
                family_name: faker.person.lastName(),
                middle_name: faker.person.middleName(),
                given_name: faker.person.firstName(),
                gender: '1',
                birthdate: null,
                avatar: null,
                locale_id: 'ja_JP',

                created_at: faker.date.past().toString(),
                updated_at: faker.date.past().toString(),
              },
              is_staff: true,
            },
          },
        ]}
      >
        <ProfileDatetime />
      </RecoilController>
    );
  },
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof ProfileDatetime>;

export const Default: Story = {};
