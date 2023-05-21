import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {Header} from '../../../components/Common/Frame/Header';
import {UserState} from '../../../utils/state/atom';
import {RecoilController} from '../../RecoilController';

const meta: Meta<typeof Header> = {
  title: 'CateiruSSO/Common/Frame/Header',
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

                created: faker.date.past().toString(),
                modified: faker.date.past().toString(),
              },
            },
          },
          {
            key: 'login no avatar',
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
                avatar: null,
                locale_id: 'ja_JP',

                created: faker.date.past().toString(),
                modified: faker.date.past().toString(),
              },
            },
          },
        ]}
      >
        <Header />
      </RecoilController>
    );
  },
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof Header>;

export const Default: Story = {};
