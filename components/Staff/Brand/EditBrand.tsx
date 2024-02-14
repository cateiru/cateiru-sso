'use client';

import {
  Button,
  FormControl,
  FormErrorMessage,
  FormLabel,
  Input,
  Textarea,
} from '@chakra-ui/react';
import {useRouter} from 'next/navigation';
import React from 'react';
import {useForm} from 'react-hook-form';
import useSWR from 'swr';
import {brandFeather} from '../../../utils/swr/staff';
import {ErrorType} from '../../../utils/types/error';
import {Brand} from '../../../utils/types/staff';
import {Error} from '../../Common/Error/Error';
import {useRequest} from '../../Common/useRequest';

interface Props {
  id: string;
}

interface EditBrandFromData {
  name: string;
  description?: string;
}

export const EditBrand: React.FC<Props> = ({id}) => {
  const {data, error} = useSWR<Brand, ErrorType>(
    `/v2/admin/brand?brand_id=${id}`,
    () => brandFeather(id)
  );

  const {
    handleSubmit,
    register,
    setValue,
    formState: {errors, isSubmitting},
  } = useForm<EditBrandFromData>();
  const router = useRouter();
  const {request} = useRequest('/admin/brand');

  React.useEffect(() => {
    if (data) {
      setValue('name', data.name);
      setValue('description', data.description ?? '');
    }
  }, [data]);

  const onSubmit = async (data: EditBrandFromData) => {
    const form = new FormData();
    form.append('brand_id', id);
    form.append('name', data.name);

    if (data.description) {
      form.append('description', data.description);
    }

    const res = await request({
      method: 'PUT',
      body: form,
    });

    if (res) {
      router.push(`/staff/brand/${id}`);
    }
  };

  if (error) {
    return <Error {...error} />;
  }

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <FormControl isInvalid={!!errors.name}>
        <FormLabel htmlFor="name">ブランド名</FormLabel>
        <Input
          id="name"
          {...register('name', {
            required: 'ブランド名は必須です',
          })}
        />
        <FormErrorMessage>
          {errors.name && errors.name.message}
        </FormErrorMessage>
      </FormControl>
      <FormControl isInvalid={!!errors.description} mt="1rem">
        <FormLabel htmlFor="description">説明</FormLabel>
        <Textarea id="description" {...register('description')} />
        <FormErrorMessage>
          {errors.description && errors.description.message}
        </FormErrorMessage>
      </FormControl>
      <Button
        mt={4}
        colorScheme="cateiru"
        isLoading={isSubmitting}
        type="submit"
        w="100%"
      >
        ブランドを更新
      </Button>
    </form>
  );
};
