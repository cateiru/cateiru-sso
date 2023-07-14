import type {Meta, StoryObj} from '@storybook/react';
import {OrganizationDetail} from '../../components/Organization/OrganizationDetail';

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

export const Default: Story = {};
