import {
  Box,
  FormControl,
  FormLabel,
  Input,
  Flex,
  Button,
} from '@chakra-ui/react';
import React from 'react';
import {useForm} from 'react-hook-form';
import type {FieldValues} from 'react-hook-form';

const CreateInfo = () => {
  const {
    handleSubmit,
    register,
    formState: {errors, isSubmitting},
  } = useForm();

  const submit = (values: FieldValues) => {
    console.log(JSON.stringify(values));
  };

  return (
    <Box width={{base: '90%', md: '600px'}}>
      <form onSubmit={handleSubmit(submit)}>
        <Flex>
          <FormControl isInvalid={errors.firstName}>
            <FormLabel htmlFor="firstName">名前</FormLabel>
            <Input
              id="firstName"
              type="name"
              placeholder="名字"
              {...register('firstName', {
                required: '名字が必要です',
              })}
            />
          </FormControl>
          <FormControl isInvalid={errors.lastName}>
            <FormLabel htmlFor="lastName">名前</FormLabel>
            <Input
              id="lastName"
              type="name"
              placeholder="名前"
              {...register('lastName', {
                required: '名前が必要です',
              })}
            />
          </FormControl>
        </Flex>
        <FormControl isInvalid={errors.userName}>
          <FormLabel htmlFor="userName">ユーザ名</FormLabel>
          <Input
            id="userName"
            type="name"
            placeholder="ユーザ名"
            {...register('userName', {
              required: 'ユーザ名が必要です',
            })}
          />
        </FormControl>
      </form>
      <Button
        marginTop="1rem"
        colorScheme="blue"
        isLoading={isSubmitting}
        type="submit"
      >
        これでOK
      </Button>
    </Box>
  );
};

export default CreateInfo;
