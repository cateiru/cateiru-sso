import {api} from '../api';
import {HTTPError} from '../error';
import {ErrorSchema} from '../types/error';
import {
  LoginDeviceListScheme,
  LoginTryHistoryListScheme,
  OperationHistoryListScheme,
} from '../types/history';

export async function loginDeviceFeather(path: string) {
  const res = await fetch(api(path), {
    credentials: 'include',
    mode: 'cors',
  });

  if (!res.ok) {
    const data = ErrorSchema.safeParse(await res.json());
    if (data.success) {
      throw data.data;
    }
    console.error(data.error.message);
    throw new HTTPError(data.error.message);
  }

  const data = LoginDeviceListScheme.safeParse(await res.json());
  if (data.success) {
    return data.data;
  }
  console.error(data.error.message);
  throw new HTTPError(data.error.message);
}

export async function loginTryHistoryFeather() {
  const res = await fetch(api('/v2/history/try_login'), {
    credentials: 'include',
    mode: 'cors',
  });

  if (!res.ok) {
    const data = ErrorSchema.safeParse(await res.json());
    if (data.success) {
      throw data.data;
    }
    console.error(data.error.message);
    throw new HTTPError(data.error.message);
  }

  const data = LoginTryHistoryListScheme.safeParse(await res.json());
  if (data.success) {
    return data.data;
  }
  console.error(data.error.message);
  throw new HTTPError(data.error.message);
}

export async function operationHistoryFeather() {
  const res = await fetch(api('/v2/history/operation'), {
    credentials: 'include',
    mode: 'cors',
  });

  if (!res.ok) {
    const data = ErrorSchema.safeParse(await res.json());
    if (data.success) {
      throw data.data;
    }
    console.error(data.error.message);
    throw new HTTPError(data.error.message);
  }

  const data = OperationHistoryListScheme.safeParse(await res.json());
  if (data.success) {
    return data.data;
  }
  console.error(data.error.message);
  throw new HTTPError(data.error.message);
}
