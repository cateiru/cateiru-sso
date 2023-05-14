import type {Meta, StoryObj} from '@storybook/react';
import {Device} from '../../components/Histories/Device';

const meta: Meta<typeof Device> = {
  title: 'CateiruSSO/Histories/Device',
  component: Device,
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<typeof Device>;

export const DesktopMacChrome: Story = {
  args: {
    device: '',
    os: 'macOS',
    browser: 'Google Chrome',
    isMobile: false,
  },
};

export const DesktopWindowsChrome: Story = {
  args: {
    device: '',
    os: 'Windows',
    browser: 'Google Chrome',
    isMobile: false,
  },
};

export const Mobile: Story = {
  args: {
    device: 'iPhone',
    os: 'iOS',
    browser: 'Safari',
    isMobile: true,
  },
};
