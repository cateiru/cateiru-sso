import type {Meta, StoryObj} from '@storybook/react';
import {RegisterCertificatePage} from '../../../components/RegisterAccount/RegisterCertificatePage';

const meta: Meta<typeof RegisterCertificatePage> = {
  title: 'CateiruSSO/RegisterAccount/RegisterCertificate/Page',
  component: RegisterCertificatePage,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof RegisterCertificatePage>;

export const Default: Story = {
  args: {},
};
