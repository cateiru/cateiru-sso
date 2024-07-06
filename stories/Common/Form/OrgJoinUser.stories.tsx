import type {Meta, StoryObj} from '@storybook/react';
import {OrgJoinUser} from '../../../components/Common/Form/OrgJoinUser';

const meta: Meta<typeof OrgJoinUser> = {
  title: 'CateiruSSO/Common/Form/OrgJoinUser',
  component: OrgJoinUser,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof OrgJoinUser>;

export const Default: Story = {};
