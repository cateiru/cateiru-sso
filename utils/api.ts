import {URL} from 'url';
import {config} from './config';

/**
 * @param path - api path
 */
export function api(path: string): URL {
  const url = new URL(config.apiHost);
  url.pathname = path;

  return url;
}
