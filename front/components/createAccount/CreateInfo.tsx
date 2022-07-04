import {
  Box,
  FormControl,
  FormLabel,
  Input,
  Flex,
  Button,
  Select,
  FormErrorMessage,
} from '@chakra-ui/react';
import {useRouter} from 'next/router';
import React from 'react';
import {useForm, FormProvider, SubmitHandler} from 'react-hook-form';
import {useSetRecoilState} from 'recoil';
import useCreateInfo from '../../hooks/useCreateInfo';
import {checkUserName} from '../../utils/api/check';
import {CTState, LoadState} from '../../utils/state/atom';
import Password, {PasswordForm} from '../common/form/Password';

interface Form extends PasswordForm {
  lastName: string;
  firstName: string;
  userName: string;
  theme: string;
}

const CreateInfo = React.memo(() => {
  const methods = useForm<Form>();
  const {
    handleSubmit,
    register,
    setError,
    clearErrors,
    formState: {errors, isSubmitting},
  } = methods;
  const [user, setUser] = React.useState('');
  const [existUser, setExistUser] = React.useState(false);
  const info = useCreateInfo();
  const router = useRouter();
  const setCT = useSetRecoilState(CTState);
  const setLoad = useSetRecoilState(LoadState);

  const [pwOK, setPwOK] = React.useState(false);

  React.useEffect(() => {
    if (user && /^[a-zA-Z0-9_]{3,15}$/.test(user)) {
      const f = async () => {
        const exist = await checkUserName(user);
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
  }, [user]);

  const submit: SubmitHandler<Form> = values => {
    if (!pwOK) {
      setError('password', {
        type: 'custom',
        message: 'custom message',
      });
      return;
    } else {
      clearErrors('password');
    }

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
      setLoad(true);
      await info(
        values.firstName,
        values.lastName,
        values.userName,
        values.theme,
        values.password
      );
      setCT('');
      router.push('/hello');
    };

    f();

    return () => {};
  };

  return (
    <Box width={{base: '90%', md: '600px'}}>
      <FormProvider {...methods}>
        <form onSubmit={handleSubmit(submit)}>
          <Flex>
            <FormControl isInvalid={Boolean(errors.lastName)} mr=".5rem">
              <FormLabel htmlFor="lastName">姓</FormLabel>
              <Input
                id="lastName"
                type="name"
                placeholder="姓"
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
                placeholder="名"
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
              placeholder="ユーザ名（英数字、アンダースコア）"
              {...register('userName', {
                required: 'ユーザ名が必要です',
                pattern: {
                  value: /^[a-zA-Z0-9_]{3,15}$/,
                  message:
                    'ユーザ名は英数字、アンダースコアで3~15文字で入力してください',
                },
                onBlur: e => setUser(e.target.value),
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
          <Password setOk={setPwOK}>
            パスワード（8文字以上128文字以下）
          </Password>
          <Button
            marginTop="1.5rem"
            colorScheme="blue"
            isLoading={isSubmitting}
            type="submit"
            width={{base: '100%', md: 'auto'}}
          >
            これでOK
          </Button>
        </form>
      </FormProvider>
    </Box>
  );
});

CreateInfo.displayName = 'CreateInfo';

export default CreateInfo;
