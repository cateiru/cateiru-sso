import {Box} from '@chakra-ui/react';
import {useRouter} from 'next/router';
import React from 'react';
import {useRecoilValue} from 'recoil';
import {UserState} from '../../utils/state/atom';

const Require: React.FC<{
  isLogin: boolean;
  path: string;
  role?: string;
  loading?: JSX.Element;
}> = ({isLogin, path, role, loading, children}) => {
  const user = useRecoilValue(UserState);
  const router = useRouter();
  const [show, setShow] = React.useState(false);

  React.useEffect(() => {
    if (typeof user !== 'undefined') {
      // ログインしていてログインしていないことを求めているの場合
      // ログインしていなくてログインしていることを求めている場合
      if ((user === null) === isLogin) {
        router.replace(path);
      } else {
        // ロールが設定されている場合でそのロールのユーザではない場合は表示させない
        if (role && !user?.role.includes(role)) {
          router.replace(path);
        } else {
          setShow(true);
        }
      }
    }
  }, [user]);

  const LoadPage = () => {
    return loading || <Box h="80vh"></Box>;
  };

  return <>{show ? children : <LoadPage />}</>;
};

export default Require;
