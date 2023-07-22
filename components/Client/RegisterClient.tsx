'use client';

import {
  Button,
  Divider,
  FormControl,
  FormErrorMessage,
  FormLabel,
  Heading,
  Input,
  Select,
  Spacer,
  Switch,
} from '@chakra-ui/react';
import {useSearchParams} from 'next/navigation';
import {FormProvider, useForm} from 'react-hook-form';
import {ImageForm, ImageFormValue} from '../Common/Form/ImageForm';
import {ListForm} from '../Common/Form/ListForm';
import {Margin} from '../Common/Margin';

interface RegisterClientForm extends ImageFormValue {
  name: string;
  description?: string;
  isAllow: boolean;
  prompt: 'login' | '2fa_login' | '';
  scopes: string[];
  orgMemberOnly?: boolean;
  redirectUrls: string[];
  referrerUrls: string[];
}

export const RegisterClient = () => {
  const param = useSearchParams();
  const orgId = param.get('org_id');

  const methods = useForm<RegisterClientForm>({
    defaultValues: {
      scopes: ['openid'],
      redirectUrls: [],
      referrerUrls: [],
    },
  });
  const {
    handleSubmit,
    register,
    formState: {errors, isSubmitting},
  } = methods;

  const onSubmit = async (data: RegisterClientForm) => {
    console.log(data);
  };

  return (
    <Margin>
      <Heading textAlign="center" mt="3rem" mb="1rem">
        {orgId && '組織'}クライアント新規作成
      </Heading>
      <FormProvider {...methods}>
        <form onSubmit={handleSubmit(onSubmit)}>
          <FormControl isInvalid={!!errors.name}>
            <FormLabel htmlFor="name">クライアント名</FormLabel>
            <Input
              id="name"
              {...register('name', {
                required: 'クライアント名は必須です',
              })}
            />
            <FormErrorMessage>
              {errors.name && errors.name.message}
            </FormErrorMessage>
          </FormControl>

          <FormControl isInvalid={!!errors.description} mt="1rem">
            <FormLabel htmlFor="description">
              クライアントの説明（オプション）
            </FormLabel>
            <Input id="description" {...register('description')} />
            <FormErrorMessage>
              {errors.description && errors.description.message}
            </FormErrorMessage>
          </FormControl>

          <FormControl isInvalid={Boolean(errors.image)} mt="1rem">
            <FormLabel htmlFor="image">アイコン画像（オプション）</FormLabel>
            <ImageForm />
            <FormErrorMessage>
              {errors.image && errors.image.message}
            </FormErrorMessage>
          </FormControl>

          <Divider my="1rem" />

          <FormControl isInvalid={Boolean(errors.scopes)} mt="1rem">
            <FormLabel htmlFor="scopes">スコープ</FormLabel>
            <ListForm name="scopes" />
            <FormErrorMessage>
              {errors.scopes && errors.scopes.message}
            </FormErrorMessage>
          </FormControl>

          <FormControl isInvalid={Boolean(errors.redirectUrls)} mt="1rem">
            <FormLabel htmlFor="redirectUrls">リダイレクトURL</FormLabel>
            <ListForm name="redirectUrls" placeholder="https://" />
            <FormErrorMessage>
              {errors.redirectUrls && errors.redirectUrls.message}
            </FormErrorMessage>
          </FormControl>

          <FormControl isInvalid={Boolean(errors.referrerUrls)} mt="1rem">
            <FormLabel htmlFor="referrerUrls">リファラーURL</FormLabel>
            <ListForm name="referrerUrls" placeholder="https://" />
            <FormErrorMessage>
              {errors.referrerUrls && errors.referrerUrls.message}
            </FormErrorMessage>
          </FormControl>

          <Divider my="1rem" />

          <FormControl
            isInvalid={!!errors.isAllow}
            mt="1rem"
            display="flex"
            alignItems="center"
          >
            <FormLabel htmlFor="isAllow" mb="0">
              使用できるユーザーを制限する
            </FormLabel>
            <Spacer />
            <Switch
              id="isAllow"
              colorScheme="cateiru"
              {...register('isAllow')}
            />
          </FormControl>

          <FormControl isInvalid={!!errors.prompt} mt="1rem">
            <FormLabel htmlFor="prompt">使用する際の認証</FormLabel>
            <Select placeholder="認証しない" {...register('prompt')}>
              <option value="login">認証を求める</option>
              <option value="2fa_login">二段階認証のみを求める</option>
            </Select>
            <FormErrorMessage>
              {errors.prompt && errors.prompt.message}
            </FormErrorMessage>
          </FormControl>

          {orgId && (
            <FormControl
              isInvalid={!!errors.orgMemberOnly}
              mt="1rem"
              display="flex"
              alignItems="center"
            >
              <FormLabel htmlFor="orgMemberOnly" mb="0">
                使用するユーザーを組織のメンバーのみに限定する
              </FormLabel>
              <Spacer />
              <Switch
                id="orgMemberOnly"
                colorScheme="cateiru"
                {...register('orgMemberOnly')}
              />
            </FormControl>
          )}

          <Button
            mt={4}
            colorScheme="cateiru"
            isLoading={isSubmitting}
            type="submit"
            w="100%"
          >
            クライアントを新規作成
          </Button>
        </form>
      </FormProvider>
    </Margin>
  );
};
