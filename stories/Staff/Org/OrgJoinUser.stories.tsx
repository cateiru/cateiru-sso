import type {Meta, StoryObj} from '@storybook/react';
import {OrgJoinUser} from '../../../components/Staff/Org/OrgJoinUser';

const meta: Meta<typeof OrgJoinUser> = {
  title: 'CateiruSSO/Staff/Org/OrgJoinUser',
  component: OrgJoinUser,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof OrgJoinUser>;

export const Default: Story = {};
