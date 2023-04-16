import {config} from './config';

/**
 * APIのURLを生成する
 *
 * @param path - API path
 * @returns APIのURL
 */
export function api(path: string): URL {
  const url = new URL(config.apiHost);
  url.pathname = path;

  return url;
}
