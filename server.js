/**
 * Based on https://zenn.dev/sotszk/articles/b4e6a4e19d2e35
 */

const fs = require('fs');
const https = require('https');
const express = require('express');
const next = require('next');

const port = 3000;
const host = '0.0.0.0';

const app = next({
  dev: process.env.NODE_ENV !== 'production',
});
const handle = app.getRequestHandler();

(async () => {
  await app.prepare();
  const expressApp = express();

  expressApp.get('*', (req, res) => handle(req, res));

  // Use HTTPS if HTTPS option enabled
  const hasCertificates =
    fs.existsSync('./certificates/tls.key') &&
    fs.existsSync('./certificates/tls.crt');
  const useHttps = process.env.HTTPS === 'true' && hasCertificates;

  if (useHttps) {
    const options = {
      key: fs.readFileSync('./certificates/tls.key'),
      cert: fs.readFileSync('./certificates/tls.crt'),
    };
    const server = https.createServer(options, expressApp);
    server.listen(port, host);
    console.log(`> Ready on https://sso.cateiru.test`);
  } else {
    expressApp.listen(port, host);
    console.log(`> Ready on http://sso.cateiru.test`);
  }
})();
