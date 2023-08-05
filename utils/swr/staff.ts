import {api} from '../api';
import {HTTPError} from '../error';
import {ErrorSchema} from '../types/error';
import {
  Brand,
  Brands,
  BrandsSchema,
  ClientDetailSchema,
  OrganizationDetailSchema,
  OrganizationsSchema,
  StaffClientsSchema,
  StaffUsersSchema,
  UserDetailSchema,
} from '../types/staff';

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

export async function staffClientsFeather() {
  const res = await fetch(api('/v2/admin/clients'), {
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

  const data = StaffClientsSchema.safeParse(await res.json());
  if (data.success) {
    return data.data;
  }
  console.error(data.error);
  throw new HTTPError(data.error.message);
}

export async function staffClientDetailFeather(id: string) {
  const urlSearchParam = new URLSearchParams();
  urlSearchParam.append('client_id', id);

  const res = await fetch(api('/v2/admin/client_detail', urlSearchParam), {
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

  const data = ClientDetailSchema.safeParse(await res.json());
  if (data.success) {
    return data.data;
  }
  console.error(data.error);
  throw new HTTPError(data.error.message);
}
