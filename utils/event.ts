import nprogress from 'nprogress';

import '../app/nprogress.css';

nprogress.configure({showSpinner: false, speed: 400, minimum: 0.25});

export const routeChangeStart = () => {
  nprogress.start();
};
export const routeChangeError = () => {
  nprogress.done();
};
export const routeChangeEnd = () => {
  nprogress.done();
};
