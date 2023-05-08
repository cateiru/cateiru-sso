import {Box, Divider, Heading} from '@chakra-ui/react';
import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {AccountList} from '../../../components/SwitchAccount/AccountList';

const meta: Meta<typeof AccountList> = {
  title: 'CateiruSSO/SwitchAccount/AccountList',
  component: ({data}) => {
    return (
      <Box
        w={{base: '96%', sm: '450px'}}
        h={{base: '600px', sm: '700px'}}
        borderWidth={{base: 'none', sm: '2px'}}
        margin="auto"
        mt="3rem"
        borderRadius="10px"
        borderColor="gray.300"
        mb={{base: '0', sm: '3rem'}}
      >
        <Box h="100px">
          <Heading textAlign="center" pt="2rem">
            アカウントを選択
          </Heading>
          <Divider mt=".5rem" w="90%" mx="auto" />
        </Box>
        <AccountList data={data} />
      </Box>
    );
  },
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof AccountList>;

export const Default: Story = {
  args: {
    data: [
      {
        id: faker.datatype.uuid(),
        user_name: faker.internet.userName(),
      },
      {
        id: faker.datatype.uuid(),
        user_name: faker.internet.userName(),
        avatar: faker.image.avatar(),
      },
    ],
  },
};

export const ManyUser: Story = {
  args: {
    data: Array(faker.datatype.number({min: 10, max: 30}))
      .fill(0)
      .map(() => {
        return {
          id: faker.datatype.uuid(),
          user_name: faker.internet.userName(),
          avatar: faker.image.avatar(),
        };
      }),
  },
};
