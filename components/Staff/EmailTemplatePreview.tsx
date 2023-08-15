'use client';

import {Link, ListItem, Text, UnorderedList} from '@chakra-ui/react';
import {api} from '../../utils/api';
import {useSecondaryColor} from '../Common/useColor';

export const EmailTemplatePreview = () => {
  const textColor = useSecondaryColor();

  return (
    <>
      <Text mb="1rem" textAlign="center" color={textColor}>
        送信するメール本文のテンプレートをプレビューします。
      </Text>
      <UnorderedList>
        <ListItem>
          <Link href={api('/v2/admin/template/register')} isExternal>
            アカウント登録
          </Link>
        </ListItem>
        <ListItem>
          <Link href={api('/v2/admin/template/register_resend')} isExternal>
            アカウント登録（再送）
          </Link>
        </ListItem>
        <ListItem>
          <Link href={api('/v2/admin/template/update_email')} isExternal>
            メールアドレス更新
          </Link>
        </ListItem>
        <ListItem>
          <Link href={api('/v2/admin/template/update_password')} isExternal>
            パスワード再設定
          </Link>
        </ListItem>
        <ListItem>
          <Link href={api('/v2/admin/template/invite_org')} isExternal>
            組織招待
          </Link>
        </ListItem>
        <ListItem>
          <Link href={api('/v2/admin/template/test')} isExternal>
            テスト
          </Link>
        </ListItem>
      </UnorderedList>
    </>
  );
};
