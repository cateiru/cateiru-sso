import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {Margin} from '../../components/Common/Margin';
import {StaffCard} from '../../components/Staff/StaffCard';

const a: typeof StaffCard = props => {
  return (
    <Margin>
      <StaffCard {...props} />
    </Margin>
  );
};

const meta: Meta<typeof StaffCard> = {
  title: 'CateiruSSO/Staff/StaffCard',
  component: a,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof StaffCard>;

export const Default: Story = {
  args: {
    title: faker.lorem.words(3),
  },
};
