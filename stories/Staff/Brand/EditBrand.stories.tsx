import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {EditBrand} from '../../../components/Staff/Brand/EditBrand';
import {api} from '../../../utils/api';
import {Brands} from '../../../utils/types/staff';

const param = new URLSearchParams();
param.append('brand_id', '1');

const meta: Meta<typeof EditBrand> = {
  title: 'CateiruSSO/Staff/Brand/EditBrand',
  component: EditBrand,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof EditBrand>;

export const Default: Story = {
  args: {
    id: '1',
  },
  parameters: {
    mockData: [
      {
        url: api('/v2/admin/brand', param),
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
        ] as Brands,
      },
    ],
  },
};
