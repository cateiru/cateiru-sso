import {useRouter} from 'next/navigation';
import React from 'react';
import {pageview} from '../utils/ga/gtag';

import '../public/nprogress.css';

export const usePageEvent = () => {
  // TODO
  // const router = useRouter();
  // React.useEffect(() => {
  //   const handleRouteChange = (url: string) => {
  //     pageview(url);
  //   };
  //   router.events.on('routeChangeComplete', handleRouteChange);
  //   router.events.on('routeChangeStart', () => {
  //     nprogress.start();
  //   });
  //   router.events.on('routeChangeComplete', () => {
  //     nprogress.done();
  //   });
  //   router.events.on('routeChangeError', () => {
  //     nprogress.done();
  //   });
  //   return () => {
  //     router.events.off('routeChangeComplete', handleRouteChange);
  //   };
  // }, [router.events]);
};
