import {Button, ListItem, UnorderedList, useDisclosure} from '@chakra-ui/react';
import {useRouter} from 'next/navigation';
import React from 'react';
import {Confirm} from '../Common/Confirm/Confirm';
import {useRequest} from '../Common/useRequest';

interface Props {
  clientId: string;
  orgId?: string;
}

export const ClientDeleteModal: React.FC<Props> = ({clientId, orgId}) => {
  const deleteModal = useDisclosure();
  const {request} = useRequest('/v2/client');
  const router = useRouter();

  const onDelete = async () => {
    const params = new URLSearchParams();
    params.append('client_id', clientId);

    const res = await request(
      {
        method: 'DELETE',
        mode: 'cors',
        credentials: 'include',
      },
      params
    );

    if (res) {
      // orgのクライアントだったら、orgのクライアント一覧にリダイレクトする
      let redirectUrl = '/clients';
      if (orgId) {
        redirectUrl = `/clients/org/${orgId}`;
      }
      router.replace(redirectUrl);
    }
  };

  return (
    <>
      <Button w="100%" onClick={deleteModal.onOpen}>
        クライアントを削除
      </Button>
      <Confirm
        onSubmit={onDelete}
        isOpen={deleteModal.isOpen}
        onClose={deleteModal.onClose}
        text={{
          confirmHeader: 'クライアントを削除しますか？',
          confirmOkText: '削除',
          confirmOkTextColor: 'red',
        }}
      >
        <UnorderedList>
          <ListItem>この操作は元に戻せません。</ListItem>
          <ListItem>
            削除すると、クライアントの利用履歴もすべて削除されます。
          </ListItem>
        </UnorderedList>
      </Confirm>
    </>
  );
};
