import type {Meta, StoryObj} from '@storybook/react';
import {Margin} from '../../components/Common/Margin';
import {SettingCardSkelton} from '../../components/Settings/SettingCardSkelton';

const a = () => {
  return (
    <Margin>
      <SettingCardSkelton />
    </Margin>
  );
};

const meta: Meta<typeof SettingCardSkelton> = {
  title: 'CateiruSSO/Settings/SettingCardSkelton',
  component: a,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof SettingCardSkelton>;

export const Default: Story = {};
