'use client';

import {
  Box,
  Button,
  Center,
  Divider,
  FormControl,
  FormErrorMessage,
  FormHelperText,
  FormLabel,
  Heading,
  Input,
  Link,
  Select,
  Spacer,
  Switch,
  Textarea,
  useToast,
} from '@chakra-ui/react';
import {useParams, useRouter} from 'next/navigation';
import React from 'react';
import {FormProvider, useForm} from 'react-hook-form';
import {
  ClientConfig,
  ClientConfigSchema,
  ClientDetailSchema,
} from '../../utils/types/client';
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
import {Spinner} from '../Common/Icons/Spinner';
import {Margin} from '../Common/Margin';
import {Link as NextLink} from '../Common/Next/Link';
import {useSecondaryColor} from '../Common/useColor';
import {useRequest} from '../Common/useRequest';

interface EditClientForm
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

export const EditClient = () => {
  const router = useRouter();
  const {id} = useParams();

  const textColor = useSecondaryColor();
  const toast = useToast();

  const {request: requestConfig} = useRequest('/client/config');
  const {request: requestDeleteImage} = useRequest('/client/image');
  const {request} = useRequest('/client');

  const [config, setConfig] = React.useState<ClientConfig | undefined>();
  const [orgId, setOrgId] = React.useState<string | null | undefined>(
    undefined
  );
  const [imageUrl, setImageUrl] = React.useState<string | undefined>(undefined);
  const [loading, setLoading] = React.useState<boolean | null>(true);
  const [isAllow, setIsAllow] = React.useState<boolean>(false);

  const methods = useForm<EditClientForm>({
    defaultValues: {
      redirectUrls: [{value: ''}],
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
    const c = async () => {
      const res = await requestConfig({
        method: 'GET',
      });

      if (res) {
        const data = ClientConfigSchema.safeParse(await res.json());
        if (data.success) {
          setConfig(data.data);
          return;
        }
        console.error(data.error);
      }
    };
    const i = async () => {
      if (typeof id !== 'string') return;

      const param = new URLSearchParams();
      param.append('client_id', id);

      const res = await request(
        {
          method: 'GET',
        },
        param
      );

      if (res) {
        const data = ClientDetailSchema.safeParse(await res.json());
        if (data.success) {
          setOrgId(data.data.org_id);
          setImageUrl(data.data.image ?? undefined);

          setValue('name', data.data.name);
          setValue('description', data.data.description ?? undefined);
          setValue('isAllow', data.data.is_allow);
          setIsAllow(data.data.is_allow);
          setValue(
            'prompt',
            (data.data.prompt ?? '') as 'login' | '2fa_login' | ''
          );
          setValue('orgMemberOnly', data.data.org_member_only);
          setValue(
            'redirectUrls',
            data.data.redirect_urls.map(v => ({value: v}))
          );
          setValue(
            'referrerUrls',
            data.data.referrer_urls.map(v => ({value: `https://${v}`}))
          );
          setValue(
            'scopes',
            data.data.scopes.map(v => ({
              value: v,
              isRequired: v === 'openid', // openid は必須
            }))
          );

          // ロード終わり
          setLoading(false);

          return;
        }
        console.error(data.error);
      }

      // エラー時はnullにする
      setLoading(null);
    };
    c();
    i();
  }, []);

  const onSubmit = async (data: EditClientForm) => {
    if (typeof id !== 'string') return;

    const form = new FormData();
    form.append('client_id', id);
    form.append('name', data.name);

    if (data.description) {
      form.append('description', data.description);
    }

    if (data.image) {
      form.append('image', data.image);
    }

    form.append('is_allow', data.isAllow ? 'true' : 'false');

    if (data.prompt !== '') {
      form.append('prompt', data.prompt);
    }

    form.append('scopes', data.scopes.map(v => v.value).join(' '));

    if (orgId) {
      form.append('org_member_only', data.orgMemberOnly ? 'true' : 'false');
    }

    form.append('redirect_url_count', data.redirectUrls.length.toString());
    for (let i = 0; i < data.redirectUrls.length; i++) {
      form.append(`redirect_url_${i}`, data.redirectUrls[i].value);
    }

    if (data.referrerUrls.length > 0) {
      form.append('referrer_url_count', data.referrerUrls.length.toString());
      for (let i = 0; i < data.referrerUrls.length; i++) {
        form.append(`referrer_url_${i}`, data.referrerUrls[i].value);
      }
    }

    const res = await request({
      method: 'PUT',
      body: form,
    });

    if (res) {
      const data = ClientDetailSchema.safeParse(await res.json());
      if (data.success) {
        router.replace(`/client/${data.data.client_id}`);
        return;
      }
      console.error(data.error);
    }
  };

  const onDeleteImage = async () => {
    if (typeof id !== 'string') return;
    if (imageUrl === undefined) return;

    const param = new URLSearchParams();
    param.append('client_id', id);

    const res = await requestDeleteImage(
      {
        method: 'DELETE',
      },
      param
    );

    if (res) {
      toast({
        status: 'success',
        title: '画像を削除しました',
      });
    }
  };

  return (
    <Margin>
      <Heading textAlign="center" mt="3rem" mb="1rem">
        {orgId && '組織'}クライアントの編集
      </Heading>
      {loading === null ? (
        <></>
      ) : loading ? (
        <Center mt="1rem">
          <Spinner />
        </Center>
      ) : (
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
              <Textarea id="description" {...register('description')} />
              <FormErrorMessage>
                {errors.description && errors.description.message}
              </FormErrorMessage>
            </FormControl>

            <FormControl isInvalid={Boolean(errors.image)} mt="1rem">
              <FormLabel htmlFor="image">アイコン画像（オプション）</FormLabel>
              <ImageForm imageUrl={imageUrl} clearImage={onDeleteImage} />
              <FormErrorMessage>
                {errors.image && errors.image.message}
              </FormErrorMessage>
            </FormControl>

            <Divider my="1rem" />

            <FormControl isInvalid={Boolean(errors.scopes)} mt="1rem">
              <FormLabel htmlFor="scopes">スコープ</FormLabel>
              <FormHelperText color={textColor}>
                OpenIdConnectのスコープを設定します。
                <br />
                このスコープに設定した上限が取得可能になります。
              </FormHelperText>
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
                <br />
                最低でも1つは設定する必要があります。
              </FormHelperText>
              <RedirectUrlsForm
                maxCreatedCount={config?.redirect_url_max ?? 1}
              />
              <FormErrorMessage>
                {errors.redirectUrls &&
                  errors.redirectUrls.root &&
                  errors.redirectUrls.root?.message}
              </FormErrorMessage>
            </FormControl>

            <FormControl isInvalid={Boolean(errors.referrerUrls)} mt="1rem">
              <FormLabel htmlFor="referrerUrls">リファラーURL</FormLabel>
              <FormHelperText color={textColor}>
                アクセス元のURLです。設定をするとこれ以外のアクセス元からのURLは拒否されます。
                <br />
                リファラーはHostのみが使用されます。（例: https://example.test
                の場合 example.test からのアクセスを通過させます。）
                <br />
                リファラーURLは最大{config?.referrer_url_max ?? '-'}
                件まで作成することができます。
              </FormHelperText>
              <ReferrerUrlsForm
                maxCreatedCount={config?.referrer_url_max ?? 1}
              />
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
              <Box>
                <FormLabel htmlFor="isAllow">
                  使用できるユーザーを制限する
                </FormLabel>
                <FormHelperText color={textColor} maxW="90%">
                  この設定をONにすると、ユーザーを直接指定または、メールアドレスのドメインを指定してユーザーを制限することができます。
                  <br />
                  {isAllow && (
                    <>
                      ユーザー追加は{' '}
                      <Link
                        as={NextLink}
                        href={`/client/edit/user/${id}`}
                        fontWeight="bold"
                      >
                        許可ユーザー編集ページ
                      </Link>{' '}
                      にて行うことができます。
                    </>
                  )}
                </FormHelperText>
              </Box>

              <Spacer />
              <Switch
                id="isAllow"
                colorScheme="cateiru"
                {...register('isAllow', {
                  onChange: v => setIsAllow(v.target.checked),
                })}
              />
            </FormControl>

            <FormControl isInvalid={!!errors.prompt} mt="1rem">
              <FormLabel htmlFor="prompt">認証</FormLabel>
              <FormHelperText color={textColor} mb=".5rem">
                認証を有効化すると、アクセスを許可する前に認証を求めるようになります。
              </FormHelperText>
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
                <Box>
                  <FormLabel htmlFor="orgMemberOnly" mb="0">
                    使用するユーザーを組織のメンバーのみに限定する
                  </FormLabel>
                  <FormHelperText color={textColor} mb=".5rem">
                    この設定をONにすると組織のメンバーのみが使用できるようになります。
                    <br />
                    組織のクライアントのみの設定です。
                  </FormHelperText>
                </Box>

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
              クライアントを編集
            </Button>
          </form>
        </FormProvider>
      )}
    </Margin>
  );
};
