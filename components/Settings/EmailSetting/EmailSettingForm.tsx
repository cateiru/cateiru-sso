import {
  Button,
  Center,
  FormControl,
  FormErrorMessage,
  FormLabel,
  Input,
  Text,
  useColorModeValue,
} from '@chakra-ui/react';
import React from 'react';
import {useForm} from 'react-hook-form';
import {TbArrowBigDown} from 'react-icons/tb';
import {emailRegex} from '../../../utils/regex';

export interface EmailFormData {
  new_email: string;
}

interface Props {
  disabled: boolean;
  email: string;
  onSubmit: (data: EmailFormData) => Promise<void>;
}

export const EmailSettingForm: React.FC<Props> = props => {
  const textColor = useColorModeValue('gray.500', 'gray.400');
  const {
    handleSubmit,
    register,
    formState: {errors, isSubmitting},
  } = useForm<EmailFormData>();

  return (
    <>
      <Text textAlign="center" color={textColor}>
        現在のメールアドレス
      </Text>
      <Text textAlign="center" fontSize="1.5rem" fontWeight="bold">
        {props.email}
      </Text>
      <Center my=".5rem">
        <TbArrowBigDown size="30px" />
      </Center>
      <form onSubmit={handleSubmit(props.onSubmit)}>
        <FormControl isInvalid={!!errors.new_email} isDisabled={props.disabled}>
          <FormLabel htmlFor="new_email">新しいメールアドレス</FormLabel>
          <Input
            id="new_email"
            type="email"
            autoComplete="email"
            {...register('new_email', {
              required: '新しいメールアドレスを入力してください',
              pattern: {
                value: emailRegex,
                message: '正しいメールアドレスを入力してください',
              },
            })}
          />
          <FormErrorMessage>
            {errors.new_email && errors.new_email.message}
          </FormErrorMessage>
        </FormControl>
        <Button
          mt="1rem"
          isLoading={isSubmitting}
          colorScheme="cateiru"
          type="submit"
          w="100%"
          isDisabled={props.disabled}
        >
          メールアドレスを更新
        </Button>
      </form>
    </>
  );
};
