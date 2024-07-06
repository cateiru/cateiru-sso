import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {JoinOrganization} from '../../components/Organization/JoinOrganization';
import {api} from '../../utils/api';
import {ErrorType} from '../../utils/types/error';

const meta: Meta<typeof JoinOrganization> = {
  title: 'CateiruSSO/Organization/JoinOrganization',
  component: JoinOrganization,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof JoinOrganization>;

export const Default: Story = {
  args: {
    orgId: faker.string.uuid(),
    handleSuccess: () => {},
  },
  parameters: {
    mockData: [
      {
        url: api('/org/member'),
        method: 'POST',
        status: 200,
        delay: 1000,
        response: {},
      },
    ],
  },
};

export const NotFoundUser: Story = {
  args: {
    orgId: faker.string.uuid(),
    handleSuccess: () => {},
  },
  parameters: {
    mockData: [
      {
        url: api('/org/member'),
        method: 'POST',
        status: 400,
        delay: 1000,
        response: {
          message: 'Not found user',
          unique_code: 10,
        } as ErrorType,
      },
      {
        url: api('/org/member/invite'),
        method: 'POST',
        status: 200,
        delay: 1000,
        response: {},
      },
    ],
  },
};
