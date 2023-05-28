import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {SwitchAccount} from '../../components/SwitchAccount/SwitchAccount';
import {api} from '../../utils/api';
import {UserState} from '../../utils/state/atom';
import {RecoilController} from '../RecoilController';

const user = {
  id: '1234abc',
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
  modified_at: faker.date.past().toString(),
};

const meta: Meta<typeof SwitchAccount> = {
  title: 'CateiruSSO/SwitchAccount/SwitchAccount',
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
              user: user,
              is_staff: false,
            },
          },
        ]}
      >
        <SwitchAccount />
      </RecoilController>
    );
  },
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<typeof SwitchAccount>;

export const Default: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/v2/account/list'),
        method: 'GET',
        status: 200,
        response: [
          {
            id: user.id,
            user_name: user.user_name,
          },
          {
            id: faker.string.uuid(),
            user_name: faker.internet.userName(),
          },
          {
            id: faker.string.uuid(),
            user_name: faker.internet.userName(),
            avatar: faker.image.avatar(),
          },
        ],
      },
    ],
  },
};

export const ManyUser: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/v2/account/list'),
        method: 'GET',
        status: 200,
        response: [
          {
            id: user.id,
            user_name: user.user_name,
          },
          ...Array(faker.datatype.number({min: 10, max: 30}))
            .fill(0)
            .map(() => {
              return {
                id: faker.string.uuid(),
                user_name: faker.internet.userName(),
                avatar: faker.image.avatar(),
              };
            }),
        ],
      },
    ],
  },
};

export const Loading: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/v2/account/list'),
        method: 'GET',
        status: 200,
        delay: 10000,
        response: [
          {
            id: user.id,
            user_name: user.user_name,
          },
          {
            id: faker.string.uuid(),
            user_name: faker.internet.userName(),
          },
          {
            id: faker.string.uuid(),
            user_name: faker.internet.userName(),
            avatar: faker.image.avatar(),
          },
        ],
      },
    ],
  },
};
