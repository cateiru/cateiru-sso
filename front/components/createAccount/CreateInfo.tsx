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
import React from 'react';
import {useForm} from 'react-hook-form';
import type {FieldValues} from 'react-hook-form';

const CreateInfo = React.memo(() => {
  const {
    handleSubmit,
    register,
    formState: {errors, isSubmitting},
  } = useForm();
  const [user, setUser] = React.useState('');

  // TODO: ユーザ名が使われていないか確認する機能
  React.useEffect(() => {
    if (user) {
      console.log(user);
    }
  }, [user]);

  const submit = (values: FieldValues) => {
    console.log(JSON.stringify(values));
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
        <FormControl isInvalid={errors.userName} mt="1rem">
          <FormLabel htmlFor="userName">ユーザ名</FormLabel>
          <Input
            id="userName"
            type="name"
            placeholder="ユーザ名"
            {...register('userName', {
              required: 'ユーザ名が必要です',
              onBlur: e => setUser(e.target.value),
            })}
          />
          <FormErrorMessage>
            {errors.userName && errors.userName.message}
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
        <Button
          marginTop="1.5rem"
          colorScheme="blue"
          isLoading={isSubmitting}
          type="submit"
        >
          これでOK
        </Button>
      </form>
    </Box>
  );
});

CreateInfo.displayName = 'CreateInfo';

export default CreateInfo;
