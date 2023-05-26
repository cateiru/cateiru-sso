import type {Meta, StoryObj} from '@storybook/react';
import {WebAuthn} from '../../../../components/Settings/AccountSetting/WebAuthn/WebAuthn';

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

export const Default: Story = {};
