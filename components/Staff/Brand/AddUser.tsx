import {
  Button,
  Flex,
  FormControl,
  FormErrorMessage,
  Input,
} from '@chakra-ui/react';
import React from 'react';
import {useForm} from 'react-hook-form';
import {useRequest} from '../../Common/useRequest';

interface Props {
  userId?: string;
  brandId?: string;

  handleSuccess: () => void;
}

interface AddUserData {
  user_or_brand_id: string;
}

export const AddUser: React.FC<Props> = props => {
  const {
    handleSubmit,
    register,
    reset,
    formState: {isSubmitting, errors},
  } = useForm<AddUserData>();
  const {request} = useRequest('/admin/user/brand');

  const onSubmit = async (data: AddUserData) => {
    const form = new FormData();

    form.append('user_id', props.userId ?? data.user_or_brand_id);
    form.append('brand_id', props.brandId ?? data.user_or_brand_id);

    const res = await request({
      method: 'POST',
      body: form,
    });

    if (res) {
      reset();
      props.handleSuccess();
    }
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <FormControl isInvalid={!!errors.user_or_brand_id}>
        <Flex>
          <Input
            type="text"
            placeholder={props.userId ? 'ブランドID' : 'ユーザーID'}
            {...register('user_or_brand_id', {
              required: 'IDは必須です',
            })}
          />
          <Button
            ml=".5rem"
            colorScheme="cateiru"
            type="submit"
            isLoading={isSubmitting}
          >
            追加
          </Button>
        </Flex>
        <FormErrorMessage>
          {errors.user_or_brand_id && errors.user_or_brand_id.message}
        </FormErrorMessage>
      </FormControl>
    </form>
  );
};
