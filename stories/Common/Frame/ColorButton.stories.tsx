import type {Meta, StoryObj} from '@storybook/react';
import {ColorButton} from '../../../components/Common/Frame/ColorButton';

const meta: Meta<typeof ColorButton> = {
  title: 'CateiruSSO/Common/Frame/ColorButton',
  component: ColorButton,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof ColorButton>;

export const Default: Story = {};
