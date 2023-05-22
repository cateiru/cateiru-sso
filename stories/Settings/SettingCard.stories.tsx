import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {Margin} from '../../components/Common/Margin';
import {SettingCard} from '../../components/Settings/SettingCard';

const a: typeof SettingCard = props => {
  return (
    <Margin>
      <SettingCard {...props} />
    </Margin>
  );
};

const meta: Meta<typeof SettingCard> = {
  title: 'CateiruSSO/Settings/SettingCard',
  component: a,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof SettingCard>;

export const Default: Story = {
  args: {
    title: faker.lorem.words(3),
  },
};
