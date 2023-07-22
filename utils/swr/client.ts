import {api} from '../api';
import {HTTPError} from '../error';
import {
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
  orgId: string
): Promise<ClientListResponse>;
export async function clientFetcher(
  clientId: string,
  orgId: undefined
): Promise<ClientDetail>;
export async function clientFetcher(
  clientId: string | undefined,
  orgId: string | undefined
): Promise<ClientDetail | ClientListResponse> {
  const param = new URLSearchParams();

  if (clientId) {
    param.append('client_id', clientId);
  }
  if (orgId) {
    param.append('org_id', orgId);
  }

  const res = await fetch(api('/v2/client', param), {
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

  if (clientId) {
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
