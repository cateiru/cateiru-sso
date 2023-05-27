import {api} from '../api';
import {HTTPError} from '../error';
import {
  AccountCertificatesSchema,
  AccountUserListSchema,
  AccountWebAuthnDevicesSchema,
} from '../types/account';
import {ErrorSchema} from '../types/error';
import {
  LoginDeviceListScheme,
  LoginTryHistoryListScheme,
} from '../types/history';
import {StaffUsers} from '../types/staff';

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
    console.error(data.error.message);
    throw new HTTPError(data.error.message);
  }

  const data = AccountUserListSchema.safeParse(await res.json());
  if (data.success) {
    return data.data;
  }
  console.error(data.error.message);
  throw new HTTPError(data.error.message);
}

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

export async function userAccountCertificatesFeather() {
  const res = await fetch(api('/v2/account/certificates'), {
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

  const data = AccountCertificatesSchema.safeParse(await res.json());
  if (data.success) {
    return data.data;
  }
  console.error(data.error.message);
  throw new HTTPError(data.error.message);
}

export async function webAuthnDevicesFeather() {
  const res = await fetch(api('/v2/account/webauthn'), {
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

  const data = AccountWebAuthnDevicesSchema.safeParse(await res.json());
  if (data.success) {
    return data.data;
  }
  console.error(data.error.message);
  throw new HTTPError(data.error.message);
}

export async function staffUsersFeather() {
  const res = await fetch(api('/v2/admin/users'), {
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

  const data = StaffUsers.safeParse(await res.json());
  if (data.success) {
    return data.data;
  }
  console.error(data.error);
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
    console.error(data.error.message);
    throw new HTTPError(data.error.message);
  }
  return await res.json();
}
