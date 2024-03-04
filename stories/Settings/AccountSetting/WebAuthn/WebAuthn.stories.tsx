import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {WebAuthn} from '../../../../components/Settings/AccountSetting/WebAuthn/WebAuthn';
import {api} from '../../../../utils/api';
import {AccountWebAuthnDevices} from '../../../../utils/types/account';

const meta: Meta<typeof WebAuthn> = {
  title: 'CateiruSSO/Settings/AccountSetting/WebAuthn/WebAuthn',
  component: WebAuthn,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof WebAuthn>;

export const Default: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/account/webauthn'),
        method: 'GET',
        status: 200,
        response: Array(100)
          .fill(0)
          .map((_, i) => ({
            id: i,
            device: '',
            os: 'Windows',
            browser: 'Firefox',
            is_mobile: false,
            created_at: faker.date.past().toISOString(),
          })) as AccountWebAuthnDevices,
      },
    ],
  },
};
