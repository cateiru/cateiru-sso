import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {BrandDetailContent} from '../../../components/Staff/Brand/BrandDetailContent';

const meta: Meta<typeof BrandDetailContent> = {
  title: 'CateiruSSO/Staff/Brand/BrandDetailContent',
  component: BrandDetailContent,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof BrandDetailContent>;

export const Default: Story = {
  args: {
    brand: {
      id: faker.string.uuid(),
      name: faker.internet.domainName(),
      description: faker.lorem.paragraph(),

      created_at: faker.date.past().toString(),
      updated_at: faker.date.past().toString(),
    },
  },
};
