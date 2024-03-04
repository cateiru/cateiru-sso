import {Button, useDisclosure} from '@chakra-ui/react';
import type {Meta, StoryObj} from '@storybook/react';
import {Confirm} from '../../../components/Common/Confirm/Confirm';
import {CheckMark} from '../../../components/Common/Icons/CheckMark';

const C = () => {
  const disclosure = useDisclosure();

  const onSubmit = () => {
    console.log('submit');
  };

  return (
    <>
      <Button onClick={disclosure.onOpen}>Open</Button>
      <Confirm
        onClose={disclosure.onClose}
        isOpen={disclosure.isOpen}
        onSubmit={onSubmit}
        text={{
          confirmHeader: 'ヘッダー',
          confirmOkText: 'OK',
          confirmCancelText: 'キャンセル',
        }}
      >
        メッセージ
      </Confirm>
    </>
  );
};

const meta: Meta<typeof CheckMark> = {
  title: 'CateiruSSO/Common/Confirm/Confirm',
  component: C,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof C>;

export const Default: Story = {};
