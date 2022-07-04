import {
  Stack,
  Box,
  Center,
  Flex,
  FormControl,
  FormLabel,
  FormErrorMessage,
  Input,
  Select,
  Button,
  useToast,
} from '@chakra-ui/react';
import React from 'react';
import {SubmitHandler, useForm} from 'react-hook-form';
import {useRecoilState} from 'recoil';
import {changeUser} from '../../utils/api/change';
import {checkUserName} from '../../utils/api/check';
import {UserState} from '../../utils/state/atom';
import {UserInfo} from '../../utils/state/types';
import AvatarSetting from './AvatarSetting';

interface Form {
  lastName: string;
  firstName: string;
  userName: string;
  theme: string;
}

const UserSetting = React.memo(() => {
  const {
    handleSubmit,
    register,
    setError,
    clearErrors,
    formState: {errors, isSubmitting},
  } = useForm<Form>();
  const [existUser, setExistUser] = React.useState(false);
  const [userName, setUserName] = React.useState('');
  const [user, setUser] = useRecoilState(UserState);
  const toast = useToast();

  // ユーザ名が存在するかチェックする
  React.useEffect(() => {
    if (
      userName &&
      /^[a-zA-Z0-9_]{3,15}$/.test(userName) &&
      userName !== user?.user_name
    ) {
      const f = async () => {
        const exist = await checkUserName(userName);
        setExistUser(exist);
        if (exist) {
          setError('userName', {
            type: 'custom',
            message: 'このユーザ名はすでに使用されています',
          });
        } else {
          clearErrors('userName');
        }
      };

      f();
    }
  }, [userName]);

  const submit: SubmitHandler<Form> = values => {
    if (existUser) {
      setError('userName', {
        type: 'custom',
        message: 'このユーザ名はすでに使用されています',
      });
      return;
    } else {
      clearErrors('userName');
    }

    const f = async () => {
      let changeFirstName: string | undefined = undefined;
      let changeLastName: string | undefined = undefined;
      let changeUserName: string | undefined = undefined;
      let changeTheme: string | undefined = undefined;
      let isChanged = false;

      if (values.lastName !== user?.last_name) {
        changeLastName = values.lastName;
        isChanged = true;
      }
      if (values.firstName !== user?.first_name) {
        changeFirstName = values.firstName;
        isChanged = true;
      }
      if (values.userName !== user?.user_name) {
        changeUserName = values.userName;
        isChanged = true;
      }
      if (values.theme !== user?.theme) {
        changeTheme = values.theme;
        isChanged = true;
      }

      // なにも変更されていない場合はなにもしない
      if (!isChanged) {
        return;
      }

      let newUserInfo: UserInfo;
      try {
        newUserInfo = await changeUser(
          changeFirstName,
          changeLastName,
          changeUserName,
          changeTheme
        );
        toast({
          title: '変更しました',
          status: 'info',
          isClosable: true,
        });
      } catch (error) {
        if (error instanceof Error) {
          toast({
            title: error.message,
            status: 'error',
            isClosable: true,
            duration: 9000,
          });
        }
        return;
      }

      setUser(newUserInfo);
    };

    f();

    return () => {};
  };

  return (
    <Stack
      direction={{base: 'column', md: 'row'}}
      spacing="20px"
      mt={{base: '3rem', md: '5rem'}}
    >
      <Center width={{base: '100%', md: '80%'}} mt="2.3rem" mb="1rem">
        <AvatarSetting />
      </Center>
      <Box width="100%">
        <Center mx=".5rem" height="100%">
          <form onSubmit={handleSubmit(submit)}>
            <Flex>
              <FormControl isInvalid={Boolean(errors.lastName)} mr=".5rem">
                <FormLabel htmlFor="lastName">姓</FormLabel>
                <Input
                  id="lastName"
                  type="name"
                  defaultValue={user?.last_name}
                  placeholder={user?.last_name}
                  {...register('lastName', {
                    required: '姓の入力が必要です',
                  })}
                />
                <FormErrorMessage>
                  {errors.lastName && errors.lastName.message}
                </FormErrorMessage>
              </FormControl>
              <FormControl isInvalid={Boolean(errors.firstName)}>
                <FormLabel htmlFor="firstName">名</FormLabel>
                <Input
                  id="firstName"
                  type="name"
                  defaultValue={user?.first_name}
                  placeholder={user?.first_name}
                  {...register('firstName', {
                    required: '名が必要です',
                  })}
                />
                <FormErrorMessage>
                  {errors.firstName && errors.firstName.message}
                </FormErrorMessage>
              </FormControl>
            </Flex>
            <FormControl
              isInvalid={Boolean(errors.userName || existUser)}
              mt="1rem"
            >
              <FormLabel htmlFor="userName">ユーザ名</FormLabel>
              <Input
                id="userName"
                type="name"
                defaultValue={user?.user_name}
                placeholder={user?.user_name}
                {...register('userName', {
                  required: 'ユーザ名が必要です',
                  pattern: {
                    value: /^[a-zA-Z0-9_]{3,15}$/,
                    message:
                      'ユーザ名は英数字、アンダースコアで3~15文字で入力してください',
                  },
                  onBlur: e => setUserName(e.target.value),
                })}
              />
              <FormErrorMessage>
                {errors.userName && errors.userName.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl isInvalid={Boolean(errors.theme)} mt="1rem">
              <FormLabel htmlFor="theme">テーマ</FormLabel>
              <Select
                id="theme"
                defaultValue={user?.theme}
                placeholder="テーマを選択"
                {...register('theme', {
                  required: 'テーマの選択が必要です',
                })}
              >
                <option value="dark">ダーク</option>
                <option value="light">ライト</option>
              </Select>
              <FormErrorMessage>
                {errors.theme && errors.theme.message}
              </FormErrorMessage>
            </FormControl>
            <Button
              marginTop="1.5rem"
              colorScheme="blue"
              isLoading={isSubmitting}
              type="submit"
              width={{base: '100%', md: 'auto'}}
            >
              変える
            </Button>
          </form>
        </Center>
      </Box>
    </Stack>
  );
});

UserSetting.displayName = 'userSetting';

export default UserSetting;
