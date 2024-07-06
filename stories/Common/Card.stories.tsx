import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {Card} from '../../components/Common/Card';
import {Margin} from '../../components/Common/Margin';

const a: typeof Card = props => {
  return (
    <Margin>
      <Card {...props} />
    </Margin>
  );
};

const meta: Meta<typeof Card> = {
  title: 'CateiruSSO/Common/Card',
  component: a,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof Card>;

export const Default: Story = {
  args: {
    title: faker.lorem.words(3),
  },
};

export const Desctiption: Story = {
  args: {
    title: faker.lorem.words(3),
    description: faker.lorem.words(10),
  },
};
