import {
  Box,
  FormControl,
  FormLabel,
  Input,
  Flex,
  Button,
  Select,
  FormErrorMessage,
  InputGroup,
  InputRightElement,
  IconButton,
} from '@chakra-ui/react';
import dynamic from 'next/dynamic';
import {useRouter} from 'next/router';
import React from 'react';
import {useForm} from 'react-hook-form';
import type {FieldValues} from 'react-hook-form';
import {IoEyeOutline, IoEyeOffOutline} from 'react-icons/io5';
import PasswordStrengthBar from 'react-password-strength-bar';
import {useSetRecoilState} from 'recoil';
import useCreateInfo from '../../hooks/useCreateInfo';
import {checkUserName} from '../../utils/api/check';
import {CTState, LoadState} from '../../utils/state/atom';

const PasswordChecklist = dynamic(() => import('react-password-checklist'), {
  ssr: false,
});

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

  const [show, setShow] = React.useState(false);
  const [pass, setPass] = React.useState('');
  const [pwOK, setPwOK] = React.useState(false);

  React.useEffect(() => {
    if (user) {
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
            <FormLabel htmlFor="lastName">名字</FormLabel>
            <Input
              id="lastName"
              type="name"
              placeholder="名字"
              {...register('lastName', {
                required: '名字が必要です',
              })}
            />
            <FormErrorMessage>
              {errors.lastName && errors.lastName.message}
            </FormErrorMessage>
          </FormControl>
          <FormControl isInvalid={errors.firstName}>
            <FormLabel htmlFor="firstName">名前</FormLabel>
            <Input
              id="firstName"
              type="name"
              placeholder="名前"
              {...register('firstName', {
                required: '名前が必要です',
              })}
            />
            <FormErrorMessage>
              {errors.firstName && errors.firstName.message}
            </FormErrorMessage>
          </FormControl>
        </Flex>
        <FormControl isInvalid={errors.userName || existUser} mt="1rem">
          <FormLabel htmlFor="userName">ユーザ名（小文字英数字）</FormLabel>
          <Input
            id="userName"
            type="name"
            placeholder="ユーザ名（小文字英数字）"
            {...register('userName', {
              required: 'ユーザ名が必要です',
              pattern: {
                value: /^[a-z0-9]+$/,
                message: 'ユーザ名は小文字英数字で入力してください',
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
        <FormControl isInvalid={errors.password}>
          <FormLabel htmlFor="password" marginTop="1rem">
            パスワード（8文字以上128文字以下）
          </FormLabel>
          <InputGroup>
            <Input
              id="password"
              type={show ? 'text' : 'password'}
              placeholder="パスワード"
              {...register('password', {
                required: true,
                onChange: e => setPass(e.target.value || ''),
              })}
            />
            <InputRightElement>
              <IconButton
                variant="ghost"
                aria-label="show password"
                icon={
                  show ? (
                    <IoEyeOutline size="25px" />
                  ) : (
                    <IoEyeOffOutline size="25px" />
                  )
                }
                size="sm"
                onClick={() => setShow(!show)}
              />
            </InputRightElement>
          </InputGroup>
          <PasswordStrengthBar
            password={pass}
            scoreWords={[
              '短すぎかな',
              '弱いパスワードだと思う',
              '少し弱いパスワードかなと思う',
              'もう少し長くしてみない？',
              '最強!すごく良いよ!',
            ]}
            shortScoreWord="8文字以上にしてみよう"
            minLength={8}
          />
          <FormErrorMessage>{errors.password}</FormErrorMessage>
        </FormControl>
        <Box marginTop=".5rem">
          <PasswordChecklist
            rules={['minLength', 'specialChar', 'number', 'capital']}
            minLength={8}
            value={pass}
            messages={{
              minLength: 'パスワードは8文字以上',
              specialChar: 'パスワードに記号が含まれている',
              number: 'パスワードに数字が含まれている',
              capital: 'パスワードに大文字が含まれている',
            }}
            onChange={isValid => {
              let unmounted = false;
              if (pwOK !== isValid && !unmounted) {
                setPwOK(isValid);
              }

              return () => {
                unmounted = true;
              };
            }}
          />
        </Box>
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
