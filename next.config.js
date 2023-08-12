const withInterceptStdout = require('next-intercept-stdout');

/** @type {import('next').NextConfig} */
module.exports = withInterceptStdout(
  {
    reactStrictMode: true,
    output: 'standalone',
    experimental: {
      scrollRestoration: true,
    },
    async headers() {
      return [
        {
          source: '*',
          headers: [
            {
              key: 'cache-control',
              value: 'max-age=600, stale-while-revalidate=1800', // 10分キャッシュし、30分間は古いキャッシュを返す
            },
          ],
        },
      ];
    },
  },
  text => (text.includes('Duplicate atom key') ? '' : text)
);
