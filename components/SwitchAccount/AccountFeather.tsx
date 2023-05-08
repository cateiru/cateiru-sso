import useSWR from 'swr';
import {accountUserFeather} from '../../utils/swr/featcher';
import {AccountUserList} from '../../utils/types/account';
import {ErrorType} from '../../utils/types/error';
import {Error} from '../Common/Error/Error';
import {AccountList} from './AccountList';

export const AccountFeather = () => {
  const {data, error} = useSWR<AccountUserList, ErrorType>(
    '/',
    accountUserFeather
  );

  if (error) {
    return <Error {...error} />;
  }

  if (!data) {
    return <>aaa</>;
  }

  return <AccountList data={data} />;
};
