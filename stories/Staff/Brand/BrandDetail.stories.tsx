import type {Meta, StoryObj} from '@storybook/react';
import {BrandDetail} from '../../../components/Staff/Brand/BrandDetail';

const meta: Meta<typeof BrandDetail> = {
  title: 'CateiruSSO/Staff/Brand/BrandDetail',
  component: BrandDetail,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof BrandDetail>;

export const Default: Story = {};
