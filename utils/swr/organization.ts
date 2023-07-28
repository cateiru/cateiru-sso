import {api} from '../api';
import {HTTPError} from '../error';
import {ErrorSchema} from '../types/error';
import {
  OrganizationInviteMemberListSchema,
  OrganizationUserListSchema,
  PublicOrganizationDetailSchema,
  PublicOrganizationListSchema,
  SimpleOrganizationListSchema,
} from '../types/organization';

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

export async function orgUsersFeather(id: string) {
  const urlSearchParam = new URLSearchParams();
  urlSearchParam.append('org_id', id);

  const res = await fetch(api('/v2/org/member', urlSearchParam), {
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

  const data = OrganizationUserListSchema.safeParse(await res.json());
  if (data.success) {
    return data.data;
  }
  console.error(data.error);
  throw new HTTPError(data.error.message);
}

export async function orgSimpleListFeather(id?: string, isJoined?: boolean) {
  if (!isJoined) return [];

  const urlSearchParam = new URLSearchParams();
  if (id) {
    urlSearchParam.append('org_id', id);
  }

  const res = await fetch(api('/v2/org/list/simple', urlSearchParam), {
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

  const data = SimpleOrganizationListSchema.safeParse(await res.json());
  if (data.success) {
    return data.data;
  }
  console.error(data.error);
  throw new HTTPError(data.error.message);
}

export async function orgInviteMemberListFeather(id: string) {
  const urlSearchParam = new URLSearchParams();
  urlSearchParam.append('org_id', id);

  const res = await fetch(api('/v2/org/member/invite', urlSearchParam), {
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

  const data = OrganizationInviteMemberListSchema.safeParse(await res.json());
  if (data.success) {
    return data.data;
  }
  console.error(data.error);
  throw new HTTPError(data.error.message);
}
