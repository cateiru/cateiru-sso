import {config} from './config';

/**
 * APIのURLを生成する
 *
 * @param path - API path
 * @returns APIのURL
 */
export function api(path: string, searchParams?: URLSearchParams): string {
  if (config.apiHost) {
    const url = new URL(
      config.apiPathPrefix + path,
      config.apiHost ?? location.href
    );

    if (searchParams) {
      searchParams.forEach((value, key) => {
        url.searchParams.append(key, value);
      });
    }

    return url.toString();
  }

  const mergedPath = `${config.apiPathPrefix}${path}`;

  if (searchParams) {
    return `${mergedPath}?${searchParams.toString()}`;
  }
  return mergedPath;
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
