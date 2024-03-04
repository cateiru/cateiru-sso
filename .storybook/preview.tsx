import {ChakraProvider, useColorMode} from '@chakra-ui/react';
import React from 'react';
import {theme} from '../utils/theme';
import {GoogleReCaptchaProvider} from 'react-google-recaptcha-v3';
import {UserState} from '../utils/state/atom';
import {faker} from '@faker-js/faker';
import {useSetAtom} from 'jotai';

interface ProviderProps<T> {
  value: T;
  children: JSX.Element;
}

function ColorMode(props: ProviderProps<'light' | 'dark'>) {
  const {setColorMode} = useColorMode();

  React.useEffect(() => {
    setColorMode(props.value);
  }, [props.value]);

  return props.children;
}

function JotaiUser(
  props: ProviderProps<
    | 'noLogin'
    | 'login'
    | 'loginAndJoinedOrg'
    | 'loading'
    | 'loginNoAvatar'
    | 'loginAdmin'
  >
) {
  const setUser = useSetAtom(UserState);

  React.useEffect(() => {
    switch (props.value) {
      case 'noLogin':
        setUser(null);
        break;
      case 'login':
        setUser({
          user: {
            id: '123',
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
          joined_organization: false,
        });
        break;
      case 'loginAndJoinedOrg':
        setUser({
          user: {
            id: '123',
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
        });
        break;
      case 'loginNoAvatar':
        setUser({
          user: {
            id: '123',
            user_name: faker.internet.userName(),
            email: faker.internet.email(),
            family_name: faker.person.lastName(),
            middle_name: null,
            given_name: faker.person.firstName(),
            gender: '1',
            birthdate: null,
            avatar: null,
            locale_id: 'ja_JP',

            created_at: faker.date.past().toString(),
            updated_at: faker.date.past().toString(),
          },
          is_staff: false,
          joined_organization: false,
        });
        break;
      case 'loginAdmin':
        setUser({
          user: {
            id: '123',
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
          is_staff: true,
          joined_organization: false,
        });
        break;
      case 'loading':
        setUser(undefined);
    }
  }, [props.value]);

  return props.children;
}

export const decorators = [
  (Story, context) => {
    return (
      <ChakraProvider theme={theme}>
        <GoogleReCaptchaProvider reCaptchaKey="empty_recaptcha_key">
          <ColorMode value={context.globals.colorMode}>
            <JotaiUser value={context.globals.user}>
              <Story />
            </JotaiUser>
          </ColorMode>
        </GoogleReCaptchaProvider>
      </ChakraProvider>
    );
  },
];

export const globalTypes = {
  colorMode: {
    name: 'Color Mode',
    defaultValue: 'light',
    toolbar: {
      items: [
        {title: 'Light', value: 'light'},
        {title: 'Dark', value: 'dark'},
      ],
      dynamicTitle: true,
    },
  },
  user: {
    name: 'User',
    defaultValue: 'noLogin',
    toolbar: {
      items: [
        {title: 'NoLogin', value: 'noLogin'},
        {title: 'Login', value: 'login'},
        {title: 'LoginAndJoinedOrg', value: 'loginAndJoinedOrg'},
        {title: 'LoginNoAvatar', value: 'loginNoAvatar'},
        {title: 'LoginWithAdmin', value: 'loginAdmin'},
        {title: 'Loading', value: 'loading'},
      ],
      dynamicTitle: true,
    },
  },
};

export const parameters = {
  nextjs: {
    appDirectory: true,
  },
};
