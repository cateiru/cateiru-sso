import {api, fetch} from '../api';
import {HTTPError} from '../error';
import {
  ClientAllowUserListSchema,
  ClientDetail,
  ClientDetailSchema,
  ClientListResponse,
  ClientListResponseSchema,
} from '../types/client';
import {ErrorSchema} from '../types/error';

export async function clientFetcher(
  clientId: undefined,
  orgId: undefined
): Promise<ClientListResponse>;
export async function clientFetcher(
  clientId: undefined,
  orgId: string | string[]
): Promise<ClientListResponse>;
export async function clientFetcher(
  clientId: string | string[],
  orgId: undefined
): Promise<ClientDetail>;
export async function clientFetcher(
  clientId: string | string[] | undefined,
  orgId: string | string[] | undefined
): Promise<ClientDetail | ClientListResponse> {
  const param = new URLSearchParams();

  if (typeof clientId === 'string') {
    param.append('client_id', clientId);
  }
  if (typeof orgId === 'string') {
    param.append('org_id', orgId);
  }

  const res = await fetch(api('/client', param));

  if (!res.ok) {
    const data = ErrorSchema.safeParse(await res.json());
    if (data.success) {
      throw data.data;
    }
    console.error(data.error.message);
    throw new HTTPError(data.error.message);
  }

  if (typeof clientId === 'string') {
    const data = ClientDetailSchema.safeParse(await res.json());
    if (data.success) {
      return data.data;
    }
    console.error(data.error.message);
    throw new HTTPError(data.error.message);
  }

  const data = ClientListResponseSchema.safeParse(await res.json());
  if (data.success) {
    return data.data;
  }
  console.error(data.error);
  throw new HTTPError(data.error.message);
}

export async function allowUserFetcher(clientId: string | string[]) {
  if (typeof clientId !== 'string') return;

  const param = new URLSearchParams();
  param.append('client_id', clientId);

  const res = await fetch(api('/client/allow_user', param));

  if (!res.ok) {
    const data = ErrorSchema.safeParse(await res.json());
    if (data.success) {
      throw data.data;
    }
    console.error(data.error.message);
    throw new HTTPError(data.error.message);
  }

  const data = ClientAllowUserListSchema.safeParse(await res.json());
  if (data.success) {
    return data.data;
  }
  console.error(data.error.message);
  throw new HTTPError(data.error.message);
}
