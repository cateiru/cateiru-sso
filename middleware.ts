import {NextResponse} from 'next/server';
import type {NextRequest} from 'next/server';
import {config} from './utils/config';

export const middleware = (request: NextRequest) => {
  const requestHeaders = new Headers(request.headers);
  const url = new URL(request.nextUrl);

  let cache = 's-maxage=31536000, max-age=0, stale-while-revalidate=31536000';

  // 画像などの静的ファイルはパスにハッシュが含まれているので、
  // ブラウザにキャッシュを持ってもOK
  if (/\.(js|png|jpeg|jpg|css|svg)$/.test(url.pathname)) {
    cache = 'public, max-age=31536000, immutable';
  }

  requestHeaders.set('Cache-Control', cache);
  requestHeaders.set('X-Revision', config.revision);

  return NextResponse.next({
    request: {
      // New request headers
      headers: requestHeaders,
    },
  });
};
