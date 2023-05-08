import {api} from '../api';
import {HTTPError} from '../error';
import {AccountUserListSchema} from '../types/account';
import {ErrorSchema} from '../types/error';

export async function accountUserFeather() {
  const res = await fetch(api('/v2/account/list'), {
    credentials: 'include',
    mode: 'cors',
  });

  if (!res.ok) {
    const data = ErrorSchema.safeParse(await res.json());
    if (data.success) {
      throw data.data;
    }
    throw new HTTPError(data.error.message);
  }

  const data = AccountUserListSchema.safeParse(await res.json());
  if (data.success) {
    return data.data;
  }
  throw new HTTPError(data.error.message);
}

export async function fetcher(url: string) {
  const res = await fetch(api(url), {
    credentials: 'include',
    mode: 'cors',
  });

  if (!res.ok) {
    const data = ErrorSchema.safeParse(await res.json());
    if (data.success) {
      throw data.data;
    }
    throw new HTTPError(data.error.message);
  }
  return await res.json();
}
