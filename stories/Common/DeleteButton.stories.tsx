import type {Meta, StoryObj} from '@storybook/react';
import {TbTrashX} from 'react-icons/tb';
import {DeleteButton} from '../../components/Common/DeleteButton';

const meta: Meta<typeof DeleteButton> = {
  title: 'CateiruSSO/Common/DeleteButton',
  component: DeleteButton,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof DeleteButton>;

export const Default: Story = {
  args: {
    tooltipLabel: '削除',
    onSubmit: async () => {
      await new Promise(resolve => setTimeout(resolve, 1000));
    },
    text: {
      confirmHeader: '削除しますか？',
      confirmOkText: '削除',
      confirmCancelText: 'キャンセル',
      confirmOkTextColor: 'red',
    },
    icon: <TbTrashX size="25px" />,
  },
};
