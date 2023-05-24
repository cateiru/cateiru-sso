import {Center, useColorModeValue} from '@chakra-ui/react';
import React from 'react';
import {TbTrashX} from 'react-icons/tb';
import {useSWRConfig} from 'swr';
import {Tooltip} from '../Common/Chakra/Tooltip';
import {useRequest} from '../Common/useRequest';

interface Props {
  loginHistoryId: number;
}

export const LogoutDevice: React.FC<Props> = props => {
  const defaultTrashColor = useColorModeValue('#CBD5E0', '#4A5568');
  const hoverTrashColor = useColorModeValue('#F56565', '#C53030');

  const {request} = useRequest('/v2/account/logout');
  const {mutate} = useSWRConfig();

  const [hover, setHover] = React.useState(false);

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
          typeof key === 'string' &&
          key.startsWith('/v2/history/login_devices'),
        undefined,
        {revalidate: true}
      );
    }
  };

  return (
    <Tooltip label="このデバイスからログアウト" placement="top">
      <Center>
        <TbTrashX
          size="25px"
          color={hover ? hoverTrashColor : defaultTrashColor}
          onMouseOver={() => setHover(true)}
          onMouseOut={() => setHover(false)}
          onClick={onLogout}
          style={{cursor: 'pointer'}}
        />
      </Center>
    </Tooltip>
  );
};
