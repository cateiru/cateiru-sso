import type {Meta, StoryObj} from '@storybook/react';
import {RegisterBrand} from '../../../components/Staff/Brand/RegisterBrand';

const meta: Meta<typeof RegisterBrand> = {
  title: 'CateiruSSO/Staff/Brand/RegisterBrand',
  component: RegisterBrand,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof RegisterBrand>;

export const Default: Story = {};
