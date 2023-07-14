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
import {
  PublicOrganizationDetailSchema,
  PublicOrganizationListSchema,
} from '../types/organization';
import {
  Brand,
  Brands,
  BrandsSchema,
  OrganizationDetailSchema,
  OrganizationsSchema,
  StaffUsersSchema,
  UserDetailSchema,
} from '../types/staff';

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

  const data = StaffUsersSchema.safeParse(await res.json());
  if (data.success) {
    return data.data;
  }
  console.error(data.error);
  throw new HTTPError(data.error.message);
}

export async function staffUserDetailFeather(id: string) {
  const urlSearchParam = new URLSearchParams();
  urlSearchParam.append('user_id', id);

  const res = await fetch(api('/v2/admin/user_detail', urlSearchParam), {
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

  const data = UserDetailSchema.safeParse(await res.json());
  if (data.success) {
    return data.data;
  }
  console.error(data.error);
  throw new HTTPError(data.error.message);
}

export async function brandFeather(id: string): Promise<Brand>;
export async function brandFeather(): Promise<Brands>;
export async function brandFeather(id?: string): Promise<Brand | Brands> {
  const urlSearchParam = new URLSearchParams();

  if (typeof id !== 'undefined') {
    urlSearchParam.append('brand_id', id);
  }

  const res = await fetch(api('/v2/admin/brand', urlSearchParam), {
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

  let error;
  if (typeof id !== 'undefined') {
    const data = BrandsSchema.safeParse(await res.json());
    if (data.success) {
      return data.data[0];
    }
    error = data.error;
  } else {
    const data = BrandsSchema.safeParse(await res.json());
    if (data.success) {
      return data.data;
    }
    error = data.error;
  }
  console.error(error);
  throw new HTTPError(error.message);
}

export async function orgsFeather() {
  const res = await fetch(api('/v2/admin/orgs'), {
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

  const data = OrganizationsSchema.safeParse(await res.json());
  if (data.success) {
    return data.data;
  }
  console.error(data.error);
  throw new HTTPError(data.error.message);
}

export async function adminOrgDetailFeather(id: string) {
  const urlSearchParam = new URLSearchParams();
  urlSearchParam.append('org_id', id);

  const res = await fetch(api('/v2/admin/org', urlSearchParam), {
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

  const data = OrganizationDetailSchema.safeParse(await res.json());
  if (data.success) {
    return data.data;
  }
  console.error(data.error);
  throw new HTTPError(data.error.message);
}

export async function orgListFeather() {
  const res = await fetch(api('/v2/org/list'), {
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

  const data = PublicOrganizationListSchema.safeParse(await res.json());
  if (data.success) {
    return data.data;
  }
  console.error(data.error);
  throw new HTTPError(data.error.message);
}

export async function orgDetailFeather(id: string) {
  const urlSearchParam = new URLSearchParams();
  urlSearchParam.append('org_id', id);

  const res = await fetch(api('/v2/org/detail', urlSearchParam), {
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

  const data = PublicOrganizationDetailSchema.safeParse(await res.json());
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
