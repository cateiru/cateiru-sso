import {config} from './config';

/**
 * APIのURLを生成する
 *
 * @param path - API path
 * @returns APIのURL
 */
export function api(path: string, searchParams?: URLSearchParams): string {
  const url = new URL(
    config.apiHost ?? '',
    typeof config.apiHost === 'undefined' ? location.href : undefined
  );
  url.pathname = config.apiPathPrefix + path;

  if (searchParams) {
    searchParams.forEach((value, key) => {
      url.searchParams.append(key, value);
    });
  }

  return url.toString();
}

// fetch APIのラッパー
export async function fetch(
  input: RequestInfo,
  init?: RequestInit
): Promise<Response> {
  let c: RequestInit = {};
  if (config.apiCors) {
    c = {
      credentials: 'include',
      mode: 'cors',
    };
  }
  c = {...c, ...init};

  return await window.fetch(input, c);
}
