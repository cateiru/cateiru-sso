import type {Meta, StoryObj} from '@storybook/react';
import {Footer} from '../../../components/Common/Frame/Footer';

const meta: Meta<typeof Footer> = {
  title: 'CateiruSSO/Common/Frame/Footer',
  component: Footer,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof Footer>;

export const Default: Story = {};
