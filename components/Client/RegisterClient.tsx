'use client';

import {
  Button,
  Divider,
  FormControl,
  FormErrorMessage,
  FormHelperText,
  FormLabel,
  Heading,
  Input,
  Select,
  Spacer,
  Switch,
  useColorModeValue,
} from '@chakra-ui/react';
import {useSearchParams} from 'next/navigation';
import React from 'react';
import {FormProvider, useForm} from 'react-hook-form';
import {ClientConfig, ClientConfigSchema} from '../../utils/types/client';
import {ImageForm, ImageFormValue} from '../Common/Form/ImageForm';
import {
  RedirectUrlsForm,
  RedirectUrlsFormValue,
} from '../Common/Form/RedirectUrlsForm';
import {
  ReferrerUrlsForm,
  ReferrerUrlsFormValue,
} from '../Common/Form/ReferrerUrlsForm';
import {ScopesForm, ScopesFormValue} from '../Common/Form/ScopesFrom';
import {Margin} from '../Common/Margin';
import {useRequest} from '../Common/useRequest';

interface RegisterClientForm
  extends ImageFormValue,
    ScopesFormValue,
    RedirectUrlsFormValue,
    ReferrerUrlsFormValue {
  name: string;
  description?: string;
  isAllow: boolean;
  prompt: 'login' | '2fa_login' | '';
  orgMemberOnly?: boolean;
}

export const RegisterClient = () => {
  const param = useSearchParams();
  const orgId = param.get('org_id');

  const textColor = useColorModeValue('gray.500', 'gray.400');

  const {request} = useRequest('/v2/client/config');
  const [config, setConfig] = React.useState<ClientConfig | undefined>();

  const methods = useForm<RegisterClientForm>({
    defaultValues: {
      redirectUrls: [{value: ''}],
      referrerUrls: [{value: ''}],
    },
  });
  const {
    handleSubmit,
    register,
    setValue,
    formState: {errors, isSubmitting},
  } = methods;

  // 設定（リダイレクトURLの作成最大数など）を取得するやつ
  // SWR使ってもいいが、初回にしか使わないので愚直に書いている
  React.useEffect(() => {
    const f = async () => {
      const res = await request({
        method: 'GET',
        mode: 'cors',
        credentials: 'include',
      });

      if (res) {
        const data = ClientConfigSchema.safeParse(await res.json());
        if (data.success) {
          setConfig(data.data);
          setValue(
            'scopes',
            data.data.scopes.map(v => ({
              value: v,
              isRequired: v === 'openid', // openid は必須
            }))
          );
          return;
        }
        console.error(data.error);
      }
    };
    f();
  }, []);

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
            <ScopesForm scopes={config?.scopes ?? []} />
            <FormErrorMessage>
              {errors.scopes &&
                errors.scopes.root &&
                errors.scopes.root?.message}
            </FormErrorMessage>
          </FormControl>

          <FormControl isInvalid={Boolean(errors.redirectUrls)} mt="1rem">
            <FormLabel htmlFor="redirectUrls">リダイレクトURL</FormLabel>
            <FormHelperText color={textColor}>
              リダイレクトURLは最大{config?.redirect_url_max ?? '-'}
              件まで作成することができます。
            </FormHelperText>
            <RedirectUrlsForm maxCreatedCount={config?.redirect_url_max ?? 1} />
            <FormErrorMessage>
              {errors.redirectUrls &&
                errors.redirectUrls.root &&
                errors.redirectUrls.root?.message}
            </FormErrorMessage>
          </FormControl>

          <FormControl isInvalid={Boolean(errors.referrerUrls)} mt="1rem">
            <FormLabel htmlFor="referrerUrls">リファラーURL</FormLabel>
            <FormHelperText color={textColor}>
              リファラーURLは最大{config?.referrer_url_max ?? '-'}
              件まで作成することができます。
            </FormHelperText>
            <ReferrerUrlsForm maxCreatedCount={config?.referrer_url_max ?? 1} />
            <FormErrorMessage>
              {errors.referrerUrls &&
                errors.referrerUrls.root &&
                errors.referrerUrls.root?.message}
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
