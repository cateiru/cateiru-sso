import {Center, ListItem, UnorderedList, useToast} from '@chakra-ui/react';
import React from 'react';
import {TbTrashX} from 'react-icons/tb';
import {useSWRConfig} from 'swr';
import {DeleteButton} from '../../../Common/DeleteButton';
import {useRequest} from '../../../Common/useRequest';

export const DeleteWebAuthn: React.FC<{id: number}> = ({id}) => {
  const toast = useToast();

  const {request} = useRequest('/account/webauthn');
  const {mutate} = useSWRConfig();

  const onDeleteWebAuthn = async () => {
    const param = new URLSearchParams();
    param.append('webauthn_id', id.toString());

    const res = await request(
      {
        method: 'DELETE',
        mode: 'cors',
        credentials: 'include',
      },
      param
    );

    if (res) {
      toast({
        title: 'WebAuthnを削除しました',
        status: 'success',
      });

      mutate(
        key => typeof key === 'string' && key.startsWith('/account/webauthn'),
        undefined,
        {revalidate: true}
      );
    }
  };

  return (
    <Center>
      <DeleteButton
        onSubmit={onDeleteWebAuthn}
        tooltipLabel="この生体認証（パスキー）を削除"
        text={{
          confirmHeader: 'この生体認証（パスキー）を削除しますか？',
          confirmOkTextColor: 'red',
          confirmOkText: '削除する',
        }}
        icon={<TbTrashX size="25px" />}
      >
        <UnorderedList>
          <ListItem>一度削除すると、操作を取り消すことはできません。</ListItem>
          <ListItem>
            ブラウザのパスキーは削除しないため、手動で削除してください。
          </ListItem>
        </UnorderedList>
      </DeleteButton>
    </Center>
  );
};
