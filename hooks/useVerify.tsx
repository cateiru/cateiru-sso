import React from 'react';
import {useSetRecoilState} from 'recoil';
import {createVerify} from '../utils/api/create';
import {CTState} from '../utils/state/atom';

const useVerify = (): [(t: string) => void, boolean, boolean, string] => {
  const [isKeep, setIsKeep] = React.useState(false);
  const [load, setLoad] = React.useState(true);
  const [err, setError] = React.useState('');
  const setCT = useSetRecoilState(CTState);

  const verify = (t: string) => {
    const f = async () => {
      try {
        const resp = await createVerify(t);
        setCT(resp.client_token);
        setIsKeep(resp.keep_this_page);
        setLoad(false);
      } catch (error) {
        if (error instanceof Error) {
          setError(error.message);
        } else {
          setError('エラー');
        }
      }
    };

    f();
  };

  return [verify, isKeep, load, err];
};

export default useVerify;
