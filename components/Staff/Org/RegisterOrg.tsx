'use client';

import {
  Button,
  FormControl,
  FormErrorMessage,
  FormLabel,
  Input,
} from '@chakra-ui/react';
import {useRouter} from 'next/navigation';
import {FormProvider, useForm} from 'react-hook-form';
import {OrganizationSchema} from '../../../utils/types/staff';
import {ImageForm, ImageFormValue} from '../../Common/Form/ImageForm';
import {useRequest} from '../../Common/useRequest';

interface RegisterOrgFromData extends ImageFormValue {
  name: string;
  link?: string;
}

export const RegisterOrg = () => {
  const methods = useForm<RegisterOrgFromData>();
  const {
    handleSubmit,
    register,
    formState: {errors, isSubmitting},
  } = methods;
  const router = useRouter();
  const {request} = useRequest('/admin/org');

  const onSubmit = async (data: RegisterOrgFromData) => {
    const form = new FormData();
    form.append('name', data.name);

    if (data.link) {
      form.append('link', data.link);
    }

    if (data.image) {
      form.append('image', data.image);
    }

    const res = await request({
      method: 'POST',
      body: form,
    });

    if (res) {
      const data = OrganizationSchema.safeParse(await res.json());
      if (data.success) {
        router.push(`/staff/org/${data.data.id}`);
        return;
      }
      console.error(data.error);
    }
  };

  return (
    <FormProvider {...methods}>
      <form onSubmit={handleSubmit(onSubmit)}>
        <FormControl isInvalid={!!errors.name}>
          <FormLabel htmlFor="name">組織名</FormLabel>
          <Input
            id="name"
            {...register('name', {
              required: '組織名は必須です',
            })}
          />
          <FormErrorMessage>
            {errors.name && errors.name.message}
          </FormErrorMessage>
        </FormControl>
        <FormControl isInvalid={!!errors.link} mt="1rem">
          <FormLabel htmlFor="description">組織のURL（オプション）</FormLabel>
          <Input
            id="link"
            type="url"
            placeholder="https://"
            {...register('link')}
          />
          <FormErrorMessage>
            {errors.link && errors.link.message}
          </FormErrorMessage>
        </FormControl>
        <FormControl mt="1rem" isInvalid={Boolean(errors.image)}>
          <FormLabel htmlFor="image">アイコン画像（オプション）</FormLabel>
          <ImageForm />
          <FormErrorMessage>
            {errors.image && errors.image.message}
          </FormErrorMessage>
        </FormControl>
        <Button
          mt={4}
          colorScheme="cateiru"
          isLoading={isSubmitting}
          type="submit"
          w="100%"
        >
          組織を新規作成
        </Button>
      </form>
    </FormProvider>
  );
};
