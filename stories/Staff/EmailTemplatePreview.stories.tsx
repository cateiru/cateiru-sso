import type {Meta, StoryObj} from '@storybook/react';
import {EmailTemplatePreview} from '../../components/Staff/EmailTemplatePreview';

const meta: Meta<typeof EmailTemplatePreview> = {
  title: 'CateiruSSO/Staff/EmailTemplatePreview',
  component: EmailTemplatePreview,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof EmailTemplatePreview>;

export const Default: Story = {};
