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

    clientName: faker.internet.displayName(),
    description: faker.lorem.paragraph(),
    clientImage: faker.image.url(),

    orgMemberOnly: false,

    registerUserName: faker.internet.userName(),
    registerUserImage: faker.image.url(),

    scopes: ['openid', 'profile', 'email'],
    redirectUri: faker.internet.url(),

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

    clientName: faker.internet.displayName(),
    clientImage: faker.image.url(),

    orgMemberOnly: false,

    registerUserName: faker.internet.userName(),

    scopes: ['openid', 'profile', 'email'],
    redirectUri: faker.internet.url(),

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

    clientName: faker.internet.displayName(),
    description: Array(100).fill('a').join('\n'),
    clientImage: faker.image.url(),

    orgMemberOnly: false,

    registerUserName: faker.internet.userName(),
    registerUserImage: faker.image.url(),

    scopes: ['openid', 'profile', 'email'],
    redirectUri: faker.internet.url(),

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

    clientName: faker.internet.displayName(),
    description: faker.lorem.paragraph(),

    orgMemberOnly: false,

    registerUserName: faker.internet.userName(),
    registerUserImage: faker.image.url(),

    scopes: ['openid', 'profile'],
    redirectUri: faker.internet.url(),

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

    clientName: faker.internet.displayName(),
    description: faker.lorem.paragraph(),
    clientImage: faker.image.url(),

    registerUserName: faker.internet.userName(),
    registerUserImage: faker.image.url(),

    orgName: faker.company.name(),
    orgImage: faker.image.url(),
    orgMemberOnly: false,

    scopes: ['openid', 'profile', 'email'],
    redirectUri: faker.internet.url(),

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

    clientName: faker.internet.displayName(),
    description: faker.lorem.paragraph(),
    clientImage: faker.image.url(),

    registerUserName: faker.internet.userName(),
    registerUserImage: faker.image.url(),

    orgName: faker.company.name(),
    orgImage: faker.image.url(),
    orgMemberOnly: true,

    scopes: ['openid', 'profile', 'email'],
    redirectUri: faker.internet.url(),

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
