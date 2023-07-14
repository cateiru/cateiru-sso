import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {OrganizationDetail} from '../../components/Organization/OrganizationDetail';
import {api} from '../../utils/api';
import {PublicOrganizationDetail} from '../../utils/types/organization';

const meta: Meta<typeof OrganizationDetail> = {
  title: 'CateiruSSO/Organization/OrganizationDetail',
  component: OrganizationDetail,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof OrganizationDetail>;

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
        url: api('/v2/org/detail', urlParam),
        method: 'GET',
        status: 200,
        delay: 1000,
        response: {
          id: orgId,
          name: faker.company.name(),
          image: faker.image.url(),
          link: faker.internet.url(),
          role: 'owner',
          join_date: faker.date.past().toISOString(),
          created_at: faker.date.past().toISOString(),
        } as PublicOrganizationDetail,
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
        url: api('/v2/org/detail', urlParam),
        method: 'GET',
        status: 200,
        delay: 10000,
        response: [],
      },
    ],
  },
};
