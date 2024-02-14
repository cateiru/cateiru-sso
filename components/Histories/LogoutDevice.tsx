import {Center, ListItem, UnorderedList} from '@chakra-ui/react';
import React from 'react';
import {TbPlugConnectedX} from 'react-icons/tb';
import {useSWRConfig} from 'swr';
import {DeleteButton} from '../Common/DeleteButton';
import {useRequest} from '../Common/useRequest';

interface Props {
  loginHistoryId: number;
}

export const LogoutDevice: React.FC<Props> = props => {
  const {request} = useRequest('/account/logout');
  const {mutate} = useSWRConfig();

  const onLogout = async () => {
    const form = new FormData();
    form.append('login_history_id', String(props.loginHistoryId));

    const res = await request({
      method: 'POST',
      mode: 'cors',
      credentials: 'include',
      body: form,
    });

    if (res) {
      mutate(
        key =>
          typeof key === 'string' && key.startsWith('/history/login_devices'),
        undefined,
        {revalidate: true}
      );
    }
  };

  return (
    <Center>
      <DeleteButton
        onSubmit={onLogout}
        tooltipLabel="このデバイスからログアウト"
        text={{
          confirmHeader: 'このデバイスからログアウトしますか？',
          confirmOkTextColor: 'red',
          confirmOkText: 'ログアウト',
        }}
        icon={<TbPlugConnectedX size="25px" />}
      >
        <UnorderedList>
          <ListItem>
            一度ログアウトすると、操作を取り消すことはできません。
          </ListItem>
          <ListItem>ログイン履歴は保持されます。</ListItem>
        </UnorderedList>
      </DeleteButton>
    </Center>
  );
};
