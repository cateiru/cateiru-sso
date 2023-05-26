import {Center, useColorModeValue, useToast} from '@chakra-ui/react';
import React from 'react';
import {TbTrashX} from 'react-icons/tb';
import {useSWRConfig} from 'swr';
import {Tooltip} from '../../../Common/Chakra/Tooltip';
import {useRequest} from '../../../Common/useRequest';

export const DeleteWebAuthn: React.FC<{id: number}> = ({id}) => {
  const defaultTrashColor = useColorModeValue('#CBD5E0', '#4A5568');
  const hoverTrashColor = useColorModeValue('#F56565', '#C53030');

  const toast = useToast();

  const {request} = useRequest('/v2/account/webauthn');
  const {mutate} = useSWRConfig();

  const [hover, setHover] = React.useState(false);

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
        key =>
          typeof key === 'string' && key.startsWith('/v2/account/webauthn'),
        undefined,
        {revalidate: true}
      );
    }
  };

  return (
    <Tooltip label="この生体認証を削除" placement="top">
      <Center>
        <TbTrashX
          size="25px"
          color={hover ? hoverTrashColor : defaultTrashColor}
          onMouseOver={() => setHover(true)}
          onMouseOut={() => setHover(false)}
          onClick={onDeleteWebAuthn}
          style={{cursor: 'pointer'}}
        />
      </Center>
    </Tooltip>
  );
};
