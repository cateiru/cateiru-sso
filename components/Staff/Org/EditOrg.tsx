'use client';

import {
  Button,
  FormControl,
  FormErrorMessage,
  FormLabel,
  Input,
  useToast,
} from '@chakra-ui/react';
import {useRouter} from 'next/navigation';
import React from 'react';
import {FormProvider, useForm} from 'react-hook-form';
import useSWR from 'swr';
import {adminOrgDetailFeather} from '../../../utils/swr/staff';
import {ErrorType} from '../../../utils/types/error';
import {OrganizationDetail} from '../../../utils/types/staff';
import {Error} from '../../Common/Error/Error';
import {ImageForm, ImageFormValue} from '../../Common/Form/ImageForm';
import {useRequest} from '../../Common/useRequest';

interface Props {
  id: string;
}

interface EditOrgFromData extends ImageFormValue {
  name: string;
  link?: string;
}

export const EditOrg: React.FC<Props> = ({id}) => {
  const {data, error} = useSWR<OrganizationDetail, ErrorType>(
    `/v2/admin/org?org_id=${id}`,
    () => adminOrgDetailFeather(id)
  );

  const methods = useForm<EditOrgFromData>();
  const {
    handleSubmit,
    register,
    setValue,
    formState: {errors, isSubmitting},
  } = methods;

  const toast = useToast();
  const router = useRouter();
  const {request} = useRequest('/admin/org');
  const {request: deleteImageRequest} = useRequest('/admin/org/image');

  React.useEffect(() => {
    if (data) {
      setValue('name', data.org.name);
      setValue('link', data.org.link ?? undefined);
    }
  }, [data]);

  const onSubmit = async (data: EditOrgFromData) => {
    const form = new FormData();

    form.append('org_id', id);
    form.append('name', data.name);

    if (data.link) {
      form.append('link', data.link);
    }

    if (data.image) {
      form.append('image', data.image);
    }

    const res = await request({
      method: 'PUT',
      body: form,
    });

    if (res) {
      router.push(`/staff/org/${id}`);
    }
  };

  // 画像を削除する
  const clearImage = () => {
    const f = async () => {
      if (data?.org.image) {
        const params = new URLSearchParams();
        params.append('org_id', id);

        const res = await deleteImageRequest(
          {
            method: 'DELETE',
          },
          params
        );

        if (res) {
          toast({
            title: '組織の画像を削除しました',
            status: 'success',
          });
        }
      }
    };
    f();
  };

  if (error) {
    return <Error {...error} />;
  }

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
          <ImageForm
            imageUrl={data?.org.image ?? undefined}
            clearImage={clearImage}
          />
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
          組織を更新
        </Button>
      </form>
    </FormProvider>
  );
};
