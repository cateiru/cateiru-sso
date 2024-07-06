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
import {useForm} from 'react-hook-form';
import {BrandSchema} from '../../../utils/types/staff';
import {useRequest} from '../../Common/useRequest';

interface RegisterBrandFromData {
  name: string;
  description?: string;
}

export const RegisterBrand = () => {
  const {
    handleSubmit,
    register,
    formState: {errors, isSubmitting},
  } = useForm<RegisterBrandFromData>();
  const router = useRouter();
  const {request} = useRequest('/admin/brand');

  const onSubmit = async (data: RegisterBrandFromData) => {
    const form = new FormData();
    form.append('name', data.name);

    if (data.description) {
      form.append('description', data.description);
    }

    const res = await request({
      method: 'POST',
      body: form,
    });

    if (res) {
      const data = BrandSchema.safeParse(await res.json());
      if (data.success) {
        router.push(`/staff/brand/${data.data.id}`);
        return;
      }
      console.error(data.error);
    }
  };

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
        ブランドを新規作成
      </Button>
    </form>
  );
};
