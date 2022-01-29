import {API} from './api';

interface SetResponse {
  url: string;
}

export const setAvatar = async (image: File): Promise<string> => {
  const api = new API();

  const form = new FormData();

  form.append('upload', image);

  api.postForm(form);

  const resp = await api.connect('/user/avatar');
  const url = ((await resp.json()) as SetResponse).url;
  return url;
};
