import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {LoginSuccess} from '../../components/Login/LoginSuccess';
import {UserState} from '../../utils/state/atom';

const meta: Meta<typeof LoginSuccess> = {
  title: 'CateiruSSO/Login/LoginSuccess',
  component: LoginSuccess,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof LoginSuccess>;

export const Default: Story = {};
