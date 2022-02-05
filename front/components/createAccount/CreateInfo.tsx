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
import {useForm} from 'react-hook-form';
import type {FieldValues} from 'react-hook-form';
import {useSetRecoilState} from 'recoil';
import useCreateInfo from '../../hooks/useCreateInfo';
import {checkUserName} from '../../utils/api/check';
import {CTState, LoadState} from '../../utils/state/atom';
import Password from '../common/form/Password';

const CreateInfo = React.memo(() => {
  const {
    handleSubmit,
    register,
    formState: {errors, isSubmitting},
  } = useForm();
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
      };

      f();
    }
  }, [user]);

  const submit = (values: FieldValues) => {
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
      <form onSubmit={handleSubmit(submit)}>
        <Flex>
          <FormControl isInvalid={errors.lastName} mr=".5rem">
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
          <FormControl isInvalid={errors.firstName}>
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
        <FormControl isInvalid={errors.userName || existUser} mt="1rem">
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
            {existUser && 'このユーザ名はすでに使用されています'}
          </FormErrorMessage>
        </FormControl>
        <FormControl isInvalid={errors.theme} mt="1rem">
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
        <Password
          errors={errors}
          register={register}
          onChange={status => setPwOK(status)}
        >
          パスワード（8文字以上128文字以下）
        </Password>
        <Button
          marginTop="1.5rem"
          colorScheme="blue"
          isLoading={isSubmitting}
          type="submit"
          disabled={existUser || !pwOK}
        >
          これでOK
        </Button>
      </form>
    </Box>
  );
});

CreateInfo.displayName = 'CreateInfo';

export default CreateInfo;
