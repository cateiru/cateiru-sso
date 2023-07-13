import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {OrgList} from '../../components/Organization/OrgList';
import {api} from '../../utils/api';
import {type PublicOrganizationList} from '../../utils/types/organization';

const meta: Meta<typeof OrgList> = {
  title: 'CateiruSSO/Organization/OrgList',
  component: OrgList,
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<typeof OrgList>;

export const Default: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/v2/org/list'),
        method: 'GET',
        status: 200,
        response: [
          {
            id: faker.string.uuid(),
            name: faker.company.name(),
            image: faker.image.url(),
            link: faker.internet.url(),
            role: 'owner',
            join_date: faker.date.past().toISOString(),
          },
          {
            id: faker.string.uuid(),
            name: faker.company.name(),
            image: faker.image.url(),
            link: faker.internet.url(),
            role: 'member',
            join_date: faker.date.past().toISOString(),
          },
          {
            id: faker.string.uuid(),
            name: faker.company.name(),
            image: null,
            link: faker.internet.url(),
            role: 'guest',
            join_date: faker.date.past().toISOString(),
          },
          {
            id: faker.string.uuid(),
            name: faker.company.name(),
            image: faker.image.url(),
            link: null,
            role: 'owner',
            join_date: faker.date.past().toISOString(),
          },
          {
            id: faker.string.uuid(),
            name: faker.company.name(),
            image: null,
            link: null,
            role: 'owner',
            join_date: faker.date.past().toISOString(),
          },
        ] as PublicOrganizationList,
      },
    ],
  },
};

export const Loading: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/v2/org/list'),
        method: 'GET',
        status: 200,
        delay: 10000,
        response: [],
      },
    ],
  },
};
