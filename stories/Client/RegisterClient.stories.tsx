import type {Meta, StoryObj} from '@storybook/react';
import {RegisterClient} from '../../components/Client/RegisterClient';
import {api} from '../../utils/api';
import {ClientConfig} from '../../utils/types/client';

const meta: Meta<typeof RegisterClient> = {
  title: 'CateiruSSO/Client/RegisterClient',
  component: RegisterClient,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof RegisterClient>;

export const Default: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/v2/client/config'),
        method: 'GET',
        status: 200,
        delay: 1000,
        response: {
          redirect_url_max: 10,
          referrer_url_max: 10,
          scopes: ['profile', 'email', 'openid'],
        } as ClientConfig,
      },
    ],
  },
};

export const Loading: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/v2/client/config'),
        method: 'GET',
        status: 200,
        delay: 10000,
        response: {
          redirect_url_max: 10,
          referrer_url_max: 10,
          scopes: ['profile', 'email', 'openid'],
        } as ClientConfig,
      },
    ],
  },
};
