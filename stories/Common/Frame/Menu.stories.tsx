import {Button, MenuButton} from '@chakra-ui/react';
import type {Meta, StoryObj} from '@storybook/react';
import {Menu} from '../../../components/Common/Frame/Menu';

const meta: Meta<typeof Menu> = {
  title: 'CateiruSSO/Common/Frame/Menu',
  component: () => {
    return (
      <Menu>
        <MenuButton as={Button}>open</MenuButton>
      </Menu>
    );
  },
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof Menu>;

export const Default: Story = {};
