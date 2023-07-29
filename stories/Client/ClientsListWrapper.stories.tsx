import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {ClientsListWrapper} from '../../components/Client/ClientsListWrapper';
import {api} from '../../utils/api';
import {UserState} from '../../utils/state/atom';
import {SimpleOrganizationList} from '../../utils/types/organization';
import {RecoilController} from '../RecoilController';

const meta: Meta<typeof ClientsListWrapper> = {
  title: 'CateiruSSO/Client/ClientsListWrapper',
  component: () => {
    return (
      <RecoilController
        recoilState={UserState}
        defaultValue={{
          user: {
            id: faker.string.uuid(),
            user_name: faker.internet.userName(),
            email: faker.internet.email(),
            family_name: faker.person.lastName(),
            middle_name: null,
            given_name: faker.person.firstName(),
            gender: '1',
            birthdate: null,
            avatar: faker.image.avatar(),
            locale_id: 'ja_JP',

            created_at: faker.date.past().toString(),
            updated_at: faker.date.past().toString(),
          },
          is_staff: false,
          joined_organization: true,
        }}
        setValues={[
          {
            key: 'no login',
            value: null,
          },
          {
            key: 'login',
            value: {
              user: {
                id: faker.string.uuid(),
                user_name: faker.internet.userName(),
                email: faker.internet.email(),
                family_name: faker.person.lastName(),
                middle_name: null,
                given_name: faker.person.firstName(),
                gender: '1',
                birthdate: null,
                avatar: faker.image.avatar(),
                locale_id: 'ja_JP',

                created_at: faker.date.past().toString(),
                updated_at: faker.date.past().toString(),
              },
              is_staff: false,
              joined_organization: true,
            },
          },
        ]}
      >
        <ClientsListWrapper>
          <></>
        </ClientsListWrapper>
      </RecoilController>
    );
  },
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof ClientsListWrapper>;

export const Default: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/v2/org/list/simple'),
        method: 'GET',
        status: 200,
        delay: 1000,
        response: [
          {
            id: faker.string.uuid(),
            name: faker.company.name(),
          },
          {
            id: faker.string.uuid(),
            name: faker.company.name(),
          },
          {
            id: faker.string.uuid(),
            name: faker.company.name(),
          },
        ] as SimpleOrganizationList,
      },
    ],
  },
};

export const Loading: Story = {
  parameters: {
    layout: 'fullscreen',
    mockData: [
      {
        url: api('/v2/org/list/simple'),
        method: 'GET',
        status: 200,
        delay: 10000,
        response: [],
      },
    ],
  },
};
