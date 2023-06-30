'use client';

import {usePathname, useSearchParams} from 'next/navigation';
import React, {Suspense} from 'react';
import {
  routeChangeEnd,
  routeChangeError,
  routeChangeStart,
} from '../utils/event';
import {pageview} from '../utils/ga/gtag';

export const PageEventsImplementation: React.FC = () => {
  const pathname = usePathname();
  const searchParams = useSearchParams();

  // Chromiumのみ対応
  React.useEffect(() => {
    // まだ試験的機能なので型情報がない。そのため、一旦anyにしてしまう
    // mdn: https://developer.mozilla.org/en-US/docs/Web/API/Window/navigation
    const {navigation} = window as any;
    if (!navigation) {
      return;
    }

    navigation.addEventListener('navigate', routeChangeStart);
    navigation.addEventListener('navigateerror', routeChangeError);
    navigation.addEventListener('navigatesuccess', routeChangeEnd);
    return () => {
      navigation.removeEventListener('navigate', routeChangeStart);
      navigation.removeEventListener('navigateerror', routeChangeError);
      navigation.removeEventListener('navigatesuccess', routeChangeEnd);
    };
  }, []);

  // TODO: 一時的対応。router.eventsが復活したら削除
  React.useEffect(() => {
    const url = pathname + searchParams.toString();

    // GA
    pageview(url);

    routeChangeEnd();
  }, [pathname, searchParams]);

  return null;
};

export const PageEvents = () => (
  <Suspense fallback={null}>
    <PageEventsImplementation />
  </Suspense>
);
