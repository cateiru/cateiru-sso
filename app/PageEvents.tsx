'use client';

import {usePathname, useSearchParams} from 'next/navigation';
import nprogress from 'nprogress';
import React from 'react';
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

export const PageEvents: React.FC = () => {
  const pathname = usePathname();
  const searchParams = useSearchParams();
  // TODO: 一時的対応。router.eventsが復活したら削除
  React.useEffect(() => {
    const url = pathname + searchParams.toString();

    // GA
    pageview(url);

    nprogress.done();
  }, [pathname, searchParams]);

  return null;
};
