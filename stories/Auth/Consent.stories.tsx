import {faker} from '@faker-js/faker';
import type {Meta, StoryObj} from '@storybook/react';
import {Consent} from '../../components/Auth/Consent';

const meta: Meta<typeof Consent> = {
  title: 'CateiruSSO/Auth/Consent',
  component: Consent,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof Consent>;

export const Default: Story = {
  args: {
    userName: faker.internet.userName(),
    userImage: faker.image.url(),

    data: {
      client_name: faker.internet.displayName(),
      client_id: faker.string.uuid(),
      client_description: faker.lorem.paragraph(),
      image: faker.image.url(),

      org_name: null,
      org_image: null,
      org_member_only: false,

      redirect_uri: faker.internet.url(),
      response_type: 'code',
      scopes: ['openid', 'profile', 'email'],
      register_user_name: faker.internet.userName(),
      register_user_image: faker.image.url(),

      prompts: ['consent'],
    },

    onSubmit: () => {
      return new Promise(resolve => {
        setTimeout(() => {
          resolve();
          window.alert('submit');
        }, 1000);
      });
    },
    onCancel: () => {
      return new Promise(resolve => {
        setTimeout(() => {
          resolve();
          window.alert('cancel');
        }, 1000);
      });
    },
  },
};

export const NoDescription: Story = {
  args: {
    userName: faker.internet.userName(),

    data: {
      client_name: faker.internet.displayName(),
      client_id: faker.string.uuid(),
      client_description: null,
      image: faker.image.url(),

      org_name: null,
      org_image: null,
      org_member_only: false,

      redirect_uri: faker.internet.url(),
      response_type: 'code',
      scopes: ['openid', 'profile', 'email'],
      register_user_name: faker.internet.userName(),
      register_user_image: null,

      prompts: ['consent'],
    },

    onSubmit: () => {
      return new Promise(resolve => {
        setTimeout(() => {
          resolve();
          window.alert('submit');
        }, 1000);
      });
    },
    onCancel: () => {
      return new Promise(resolve => {
        setTimeout(() => {
          resolve();
          window.alert('cancel');
        }, 1000);
      });
    },
  },
};

export const TooLongDescription: Story = {
  args: {
    userName: faker.internet.userName(),
    userImage: faker.image.url(),

    data: {
      client_name: faker.internet.displayName(),
      client_id: faker.string.uuid(),
      client_description: Array(100).fill('a').join('\n'),
      image: faker.image.url(),

      org_name: faker.company.name(),
      org_image: faker.image.url(),
      org_member_only: false,

      redirect_uri: faker.internet.url(),
      response_type: 'code',
      scopes: ['openid', 'profile', 'email'],
      register_user_name: faker.internet.userName(),
      register_user_image: faker.image.url(),

      prompts: ['consent'],
    },

    onSubmit: () => {
      return new Promise(resolve => {
        setTimeout(() => {
          resolve();
          window.alert('submit');
        }, 1000);
      });
    },
    onCancel: () => {
      return new Promise(resolve => {
        setTimeout(() => {
          resolve();
          window.alert('cancel');
        }, 1000);
      });
    },
  },
};

export const NoImage: Story = {
  args: {
    userName: faker.internet.userName(),

    data: {
      client_name: faker.internet.displayName(),
      client_id: faker.string.uuid(),
      client_description: faker.lorem.paragraph(),
      image: null,

      org_name: faker.company.name(),
      org_image: faker.image.url(),
      org_member_only: false,

      redirect_uri: faker.internet.url(),
      response_type: 'code',
      scopes: ['openid', 'profile'],
      register_user_name: faker.internet.userName(),
      register_user_image: faker.image.url(),

      prompts: ['consent'],
    },

    onSubmit: () => {
      return new Promise(resolve => {
        setTimeout(() => {
          resolve();
          window.alert('submit');
        }, 1000);
      });
    },
    onCancel: () => {
      return new Promise(resolve => {
        setTimeout(() => {
          resolve();
          window.alert('cancel');
        }, 1000);
      });
    },
  },
};

export const Org: Story = {
  args: {
    userName: faker.internet.userName(),

    data: {
      client_name: faker.internet.displayName(),
      client_id: faker.string.uuid(),
      client_description: faker.lorem.paragraph(),
      image: null,

      org_name: faker.company.name(),
      org_image: faker.image.url(),
      org_member_only: false,

      redirect_uri: faker.internet.url(),
      response_type: 'code',
      scopes: ['openid', 'profile'],
      register_user_name: faker.internet.userName(),
      register_user_image: faker.image.url(),

      prompts: ['consent'],
    },

    onSubmit: () => {
      return new Promise(resolve => {
        setTimeout(() => {
          resolve();
          window.alert('submit');
        }, 1000);
      });
    },
    onCancel: () => {
      return new Promise(resolve => {
        setTimeout(() => {
          resolve();
          window.alert('cancel');
        }, 1000);
      });
    },
  },
};

export const OrgMemberOnly: Story = {
  args: {
    userName: faker.internet.userName(),

    data: {
      client_name: faker.internet.displayName(),
      client_id: faker.string.uuid(),
      client_description: faker.lorem.paragraph(),
      image: faker.image.url(),

      org_name: faker.company.name(),
      org_image: faker.image.url(),
      org_member_only: true,

      redirect_uri: faker.internet.url(),
      response_type: 'code',
      scopes: ['openid', 'profile'],
      register_user_name: faker.internet.userName(),
      register_user_image: faker.image.url(),

      prompts: ['consent'],
    },

    onSubmit: () => {
      return new Promise(resolve => {
        setTimeout(() => {
          resolve();
          window.alert('submit');
        }, 1000);
      });
    },
    onCancel: () => {
      return new Promise(resolve => {
        setTimeout(() => {
          resolve();
          window.alert('cancel');
        }, 1000);
      });
    },
  },
};
