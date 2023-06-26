import type {Meta, StoryObj} from '@storybook/react';
import {StaffBreadcrumbs} from '../../components/Staff/StaffBreadcrumbs';

const meta: Meta<typeof StaffBreadcrumbs> = {
  title: 'CateiruSSO/Staff/StaffBreadcrumbs',
  component: StaffBreadcrumbs,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof StaffBreadcrumbs>;

export const Default: Story = {
  args: {
    paths: [
      {
        pageName: 'スタッフ管理画面',
        href: '/staff',
      },
      {
        pageName: 'ユーザー一覧',
        href: '/staff/users',
      },
    ],
  },
};
