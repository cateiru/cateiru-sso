import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {OrganizationMember} from '../../components/Organization/OrganizationMember';
import {api} from '../../utils/api';
import {
  OrganizationInviteMemberList,
  OrganizationUserList,
} from '../../utils/types/organization';

const meta: Meta<typeof OrganizationMember> = {
  title: 'CateiruSSO/Organization/OrganizationMember',
  component: OrganizationMember,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof OrganizationMember>;

const orgId = faker.string.uuid();
const urlParam = new URLSearchParams({org_id: orgId});

export const Default: Story = {
  args: {
    id: orgId,
  },
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/v2/org/member', urlParam),
        method: 'GET',
        status: 200,
        delay: 1000,
        response: [
          {
            id: 1234,
            user: {
              id: faker.string.uuid(),
              user_name: faker.internet.userName(),
              avatar: faker.image.avatar(),
            },
            role: 'owner',

            created_at: faker.date.past().toISOString(),
            updated_at: faker.date.past().toISOString(),
          },
          ...Array(faker.number.int({min: 1, max: 10}))
            .fill(0)
            .map((_, i) => {
              return {
                id: i,
                user: {
                  id: faker.string.uuid(),
                  user_name: faker.internet.userName(),
                  avatar: faker.image.avatar(),
                },
                role: 'member',

                created_at: faker.date.past().toISOString(),
                updated_at: faker.date.past().toISOString(),
              };
            }),
          {
            id: 1235,
            user: {
              id: faker.string.uuid(),
              user_name: faker.internet.userName(),
              avatar: faker.image.avatar(),
            },
            role: 'guest',

            created_at: faker.date.past().toISOString(),
            updated_at: faker.date.past().toISOString(),
          },
        ] as OrganizationUserList,
      },
      {
        url: api('/v2/org/member/invite', urlParam),
        method: 'GET',
        status: 200,
        delay: 1000,
        response: [
          {
            id: 1234,
            email: faker.internet.email(),
            created_at: faker.date.past().toISOString(),
          },
        ] as OrganizationInviteMemberList,
      },
    ],
  },
};

export const Loading: Story = {
  args: {
    id: orgId,
  },
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/v2/org/member', urlParam),
        method: 'GET',
        status: 200,
        delay: 10000,
        response: [],
      },
      {
        url: api('/v2/org/member/invite', urlParam),
        method: 'GET',
        status: 200,
        delay: 10000,
        response: [],
      },
    ],
  },
};
