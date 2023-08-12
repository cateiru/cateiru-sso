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
          source: '/(.*)',
          headers: [
            {
              key: 'Cache-Control',
              value:
                'public s-maxage=86400, max-age=0, stale-while-revalidate=86400',
            },
          ],
        },
      ];
    },
  },
  text => (text.includes('Duplicate atom key') ? '' : text)
);
