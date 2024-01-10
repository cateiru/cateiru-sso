import {serverApi} from '../../../../utils/api';

export const GET = async () => {
  const response = await fetch(serverApi('/v2/fedcm/client_metadata'), {
    cache: 'no-store',
    headers: {
      'Sec-Fetch-Dest': 'webidentity',
    },
  });

  const data = await response.json();

  return Response.json(data, {
    headers: {
      // CDNでキャッシュするため、通常のNext.jsのキャッシュと同じにする
      'Cache-Control': 's-maxage=31536000, stale-while-revalidate',
    },
  });
};
