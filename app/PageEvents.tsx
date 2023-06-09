'use client';

import {usePathname, useSearchParams} from 'next/navigation';
import nprogress from 'nprogress';
import React, {Suspense} from 'react';
import {pageview} from '../utils/ga/gtag';

import '../public/nprogress.css';

nprogress.configure({showSpinner: false, speed: 400, minimum: 0.25});

// export const usePageEvent = () => {
//   const pathname = usePathname();
//   const searchParams = useSearchParams();
//   // TODO: 一時的対応。router.eventsが復活したら削除
//   React.useEffect(() => {
//     const url = pathname + searchParams.toString();

//     // GA
//     pageview(url);

//     nprogress.done();
//   }, [pathname, searchParams]);

//   // TODO
//   // const router = useRouter();
//   // React.useEffect(() => {
//   //   const handleRouteChange = (url: string) => {
//   //     pageview(url);
//   //   };
//   //   router.events.on('routeChangeComplete', handleRouteChange);
//   //   router.events.on('routeChangeStart', () => {
//   //     nprogress.start();
//   //   });
//   //   router.events.on('routeChangeComplete', () => {
//   //     nprogress.done();
//   //   });
//   //   router.events.on('routeChangeError', () => {
//   //     nprogress.done();
//   //   });
//   //   return () => {
//   //     router.events.off('routeChangeComplete', handleRouteChange);
//   //   };
//   // }, [router.events]);
// };

const routeChangeStart = () => {
  console.log('routeChangeStart');
  nprogress.start();
};
const routeChangeError = () => {
  nprogress.done();
};
const routeChangeEnd = () => {
  console.log('routeChangeEnd');
  nprogress.done();
};

export const PageEventsImplementation: React.FC = () => {
  const pathname = usePathname();
  const searchParams = useSearchParams();

  // まだChromeにしか対応していないが無いよりまし
  // ref. https://developer.mozilla.org/en-US/docs/Web/API/Navigation/navigate_event
  React.useEffect(() => {
    window.addEventListener('navigate', routeChangeStart);
    window.addEventListener('navigateerror', routeChangeError);
    window.addEventListener('navigatesuccess', routeChangeEnd);
    return () => {
      window.removeEventListener('navigate', routeChangeStart);
      window.removeEventListener('navigateerror', routeChangeError);
      window.removeEventListener('navigatesuccess', routeChangeEnd);
    };
  }, []);

  // TODO: 一時的対応。router.eventsが復活したら削除
  React.useEffect(() => {
    const url = pathname + searchParams.toString();

    // GA
    pageview(url);

    nprogress.done();
  }, [pathname, searchParams]);

  return null;
};

export const PageEvents = () => (
  <Suspense fallback={null}>
    <PageEventsImplementation />
  </Suspense>
);
