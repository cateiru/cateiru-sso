import {config} from './config';

/**
 * APIのURLを生成する
 *
 * @param path - API path
 * @returns APIのURL
 */
export function api(path: string, searchParams?: URLSearchParams): string {
  const url = new URL(config.apiHost);
  url.pathname = path;

  if (searchParams) {
    searchParams.forEach((value, key) => {
      url.searchParams.append(key, value);
    });
  }

  return url.toString();
}
