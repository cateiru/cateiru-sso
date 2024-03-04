import type {Meta, StoryObj} from '@storybook/react';
import {StaffFrame} from '../../components/Staff/StaffFrame';

const meta: Meta<typeof StaffFrame> = {
  title: 'CateiruSSO/Staff/StaffFrame',
  component: StaffFrame,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof StaffFrame>;

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
    title: 'ユーザー一覧',
  },
};
