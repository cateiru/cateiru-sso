import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {SwitchAccount} from '../../components/SwitchAccount/SwitchAccount';
import {api} from '../../utils/api';

const meta: Meta<typeof SwitchAccount> = {
  title: 'CateiruSSO/SwitchAccount/SwitchAccount',
  component: SwitchAccount,
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<typeof SwitchAccount>;

export const Default: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/v2/account/list').toString(),
        method: 'GET',
        status: 200,
        response: [
          {
            id: faker.datatype.uuid(),
            user_name: faker.internet.userName(),
          },
          {
            id: faker.datatype.uuid(),
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
        url: api('/v2/account/list').toString(),
        method: 'GET',
        status: 200,
        response: Array(faker.datatype.number({min: 10, max: 30}))
          .fill(0)
          .map(() => {
            return {
              id: faker.datatype.uuid(),
              user_name: faker.internet.userName(),
              avatar: faker.image.avatar(),
            };
          }),
      },
    ],
  },
};
