import {Center, useColorModeValue} from '@chakra-ui/react';
import React from 'react';
import {TbPlugConnectedX} from 'react-icons/tb';
import {Tooltip} from '../../Common/Chakra/Tooltip';
import {Spinner} from '../../Common/Icons/Spinner';
import {useRequest} from '../../Common/useRequest';

interface Props {
  userId: number;
  userName: string;

  handleSuccess: () => void;
}

export const OrgDeleteUser: React.FC<Props> = props => {
  const defaultTrashColor = useColorModeValue('#CBD5E0', '#4A5568');
  const hoverTrashColor = useColorModeValue('#F56565', '#C53030');

  const [hover, setHover] = React.useState(false);
  const [loading, setLoading] = React.useState(false);

  const {request} = useRequest('/v2/admin/org/member');

  const onDelete = () => {
    const f = async () => {
      setLoading(true);
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
      setLoading(false);
    };
    f();
  };

  return (
    <Tooltip label={`この組織から${props.userName}を削除`} placement="top">
      <Center>
        {loading ? (
          <Spinner />
        ) : (
          <TbPlugConnectedX
            size="25px"
            color={hover ? hoverTrashColor : defaultTrashColor}
            onMouseOver={() => setHover(true)}
            onMouseOut={() => setHover(false)}
            onClick={onDelete}
            style={{cursor: 'pointer'}}
          />
        )}
      </Center>
    </Tooltip>
  );
};
