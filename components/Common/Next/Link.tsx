import NextLink, {LinkProps} from 'next/link';
import React, {forwardRef} from 'react';
import {routeChangeStart} from '../../../utils/event';

type Props = Omit<
  React.AnchorHTMLAttributes<HTMLAnchorElement>,
  keyof LinkProps
> &
  LinkProps & {
    children?: React.ReactNode;
  } & React.RefAttributes<HTMLAnchorElement>;

export const Link = forwardRef<HTMLAnchorElement, Props>((props, ref) => {
  const onClick = (e: React.MouseEvent<HTMLAnchorElement, MouseEvent>) => {
    routeChangeStart();
    if (props.onClick) {
      props.onClick(e);
    }
  };

  return <NextLink {...props} ref={ref} onClick={onClick} />;
});

Link.displayName = 'Link';
