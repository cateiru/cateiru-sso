import {API} from './api';

export interface Service {
  login_count: number;
  client_id: string;
  token_secret: string;
  name: string;
  service_icon: string;
  from_url: string[];
  to_url: string[];
  user_id: string;
}

export const getSSOs = async (): Promise<Service[]> => {
  const api = new API();

  api.get();

  const resp = await api.connect('/pro/sso');

  return (await resp.json()) as Service[];
};

export const setSSOs = async (
  name: string,
  fromURLs: string[],
  toURLs: string[]
): Promise<Service> => {
  const api = new API();

  api.post(JSON.stringify({name: name, from_url: fromURLs, to_url: toURLs}));

  const resp = await api.connect('/pro/sso');

  return (await resp.json()) as Service;
};

export const editSSO = async (
  clientId: string,
  name: string,
  fromURLs: string[],
  toURLs: string[],
  changeTokenSecret: boolean
) => {
  const api = new API();

  api.post(
    JSON.stringify({
      client_id: clientId,
      name: name,
      from_url: fromURLs,
      to_url: toURLs,
      change_token_secret: changeTokenSecret,
    })
  );

  const resp = await api.connect('/pro/sso');

  return (await resp.json()) as Service;
};

export const editImage = async (
  image: File,
  clientID: string
): Promise<Service> => {
  const api = new API();

  const form = new FormData();

  form.append('client_id', clientID);
  form.append('image', image);

  api.postForm(form);

  const resp = await api.connect('/pro/sso/image');

  return (await resp.json()) as Service;
};

export const deleteService = async (clientId: string) => {
  const api = new API();

  api.delete();

  await api.connect(`/pro/sso?id=${clientId}`);
};
