import {Center} from '@chakra-ui/react';
import React from 'react';
import {TbPlugConnectedX} from 'react-icons/tb';
import {DeleteButton} from '../../Common/DeleteButton';
import {useRequest} from '../../Common/useRequest';

interface Props {
  userId: number;
  userName: string;

  handleSuccess: () => void;
}

export const OrgDeleteUser: React.FC<Props> = props => {
  const {request} = useRequest('/admin/org/member');

  const onDelete = () => {
    const f = async () => {
      const param = new URLSearchParams();
      param.append('org_user_id', props.userId.toString());

      const res = await request(
        {
          method: 'DELETE',
          mode: 'cors',
          credentials: 'include',
        },
        param
      );

      if (res) {
        props.handleSuccess();
      }
    };
    f();
  };

  return (
    <Center>
      <DeleteButton
        onSubmit={onDelete}
        tooltipLabel={`この組織から${props.userName}を削除`}
        text={{
          confirmHeader: `この組織から${props.userName}を削除しますか？`,
          confirmOkTextColor: 'red',
          confirmOkText: '削除',
        }}
        icon={<TbPlugConnectedX size="25px" />}
      />
    </Center>
  );
};
