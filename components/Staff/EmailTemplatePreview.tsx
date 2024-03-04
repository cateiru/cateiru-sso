'use client';

import {Link, ListItem, Text, UnorderedList} from '@chakra-ui/react';
import {api} from '../../utils/api';
import {useSecondaryColor} from '../Common/useColor';

export const EmailTemplatePreview = () => {
  const textColor = useSecondaryColor();

  return (
    <>
      <Text mb="1rem" color={textColor}>
        送信するメール本文のテンプレートをプレビューします。
      </Text>
      <UnorderedList w="90%" mx="auto">
        <ListItem>
          <Link href={api('/admin/template/register')}>アカウント登録</Link>
        </ListItem>
        <ListItem>
          <Link href={api('/admin/template/register_resend')}>
            アカウント登録（再送）
          </Link>
        </ListItem>
        <ListItem>
          <Link href={api('/admin/template/update_email')}>
            メールアドレス更新
          </Link>
        </ListItem>
        <ListItem>
          <Link href={api('/admin/template/update_password')}>
            パスワード再設定
          </Link>
        </ListItem>
        <ListItem>
          <Link href={api('/admin/template/invite_org')}>組織招待</Link>
        </ListItem>
        <ListItem>
          <Link href={api('/admin/template/test')}>テスト</Link>
        </ListItem>
      </UnorderedList>
    </>
  );
};
