import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {BrandsTable} from '../../../components/Staff/Brand/BrandsTable';
import {api} from '../../../utils/api';
import {Brands} from '../../../utils/types/staff';

const meta: Meta<typeof BrandsTable> = {
  title: 'CateiruSSO/Staff/Brand/BrandsTable',
  component: BrandsTable,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof BrandsTable>;

export const Default: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/admin/brand'),
        method: 'GET',
        status: 200,
        response: [
          {
            id: faker.string.uuid(),
            name: faker.company.name(),
            description: faker.company.catchPhrase(),

            created_at: faker.date.past().toISOString(),
            updated_at: faker.date.past().toISOString(),
          },
          {
            id: faker.string.uuid(),
            name: faker.company.name(),
            description: null,

            created_at: faker.date.past().toISOString(),
            updated_at: faker.date.past().toISOString(),
          },
        ] as Brands,
      },
    ],
  },
};

export const Loading: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/admin/brand'),
        method: 'GET',
        status: 200,
        delay: 10000,
        response: [] as Brands,
      },
    ],
  },
};
