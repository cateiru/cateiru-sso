'use client';

import {Box, Heading} from '@chakra-ui/react';
import React from 'react';
import {Margin} from '../Common/Margin';
import {Path, StaffBreadcrumbs} from './StaffBreadcrumbs';

interface Props {
  paths: Path[];
  title: string;
  children: React.ReactNode;
}

export const StaffFrame: React.FC<Props> = props => {
  return (
    <Margin>
      <Heading textAlign="center" mt="3rem" mb="1rem">
        {props.title}
      </Heading>
      <StaffBreadcrumbs paths={props.paths} />
      <Box mt="3rem">{props.children}</Box>
    </Margin>
  );
};
