import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {Header} from '../../../components/Common/Frame/Header';
import {UserState} from '../../../utils/state/atom';

const meta: Meta<typeof Header> = {
  title: 'CateiruSSO/Common/Frame/Header',
  component: Header,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof Header>;

export const Default: Story = {};
