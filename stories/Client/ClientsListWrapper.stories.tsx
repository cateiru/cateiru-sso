import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {ClientsListWrapper} from '../../components/Client/ClientsListWrapper';
import {api} from '../../utils/api';
import {SimpleOrganizationList} from '../../utils/types/organization';

const meta: Meta<typeof ClientsListWrapper> = {
  title: 'CateiruSSO/Client/ClientsListWrapper',
  component: () => {
    return (
      <ClientsListWrapper>
        <></>
      </ClientsListWrapper>
    );
  },
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof ClientsListWrapper>;

export const Default: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/org/list/simple'),
        method: 'GET',
        status: 200,
        delay: 1000,
        response: [
          {
            id: faker.string.uuid(),
            name: faker.company.name(),
          },
          {
            id: faker.string.uuid(),
            name: faker.company.name(),
          },
          {
            id: faker.string.uuid(),
            name: faker.company.name(),
          },
        ] as SimpleOrganizationList,
      },
    ],
    nextjs: {
      appDirectory: true,
      navigation: {
        pathname: '/',
      },
    },
  },
};

export const Loading: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/org/list/simple'),
        method: 'GET',
        status: 200,
        delay: 10000,
        response: [],
      },
    ],
  },
};
