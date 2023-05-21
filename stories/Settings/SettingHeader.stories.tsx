import type {Meta, StoryObj} from '@storybook/react';
import {SettingHeader} from '../../components/Settings/SettingHeader';

const meta: Meta<typeof SettingHeader> = {
  title: 'CateiruSSO/Settings/SettingHeader',
  component: SettingHeader,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof SettingHeader>;

export const Default: Story = {};
