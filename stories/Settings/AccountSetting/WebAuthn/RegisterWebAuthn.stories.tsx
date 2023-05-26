import type {Meta, StoryObj} from '@storybook/react';
import {RegisterWebAuthn} from '../../../../components/Settings/AccountSetting/WebAuthn/RegisterWebAuthn';

const meta: Meta<typeof RegisterWebAuthn> = {
  title: 'CateiruSSO/Settings/AccountSetting/WebAuthn/RegisterWebAuthn',
  component: RegisterWebAuthn,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof RegisterWebAuthn>;

export const Default: Story = {};
