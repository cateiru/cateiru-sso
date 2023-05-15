const withInterceptStdout = require('next-intercept-stdout');

/** @type {import('next').NextConfig} */
module.exports = withInterceptStdout(
  {
    reactStrictMode: true,
    output: 'standalone',
  },
  text => (text.includes('Duplicate atom key') ? '' : text)
);
