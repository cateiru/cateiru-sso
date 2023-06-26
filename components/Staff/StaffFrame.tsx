'use client';

import {Heading, Text, useColorModeValue} from '@chakra-ui/react';
import React from 'react';
import {config} from '../../utils/config';
import {Margin} from '../Common/Margin';
import {Path, StaffBreadcrumbs} from './StaffBreadcrumbs';

interface Props {
  paths: Path[];
  title: string;
  children: React.ReactNode;
}

export const StaffFrame: React.FC<Props> = props => {
  const textColor = useColorModeValue('gray.500', 'gray.400');

  return (
    <Margin>
      <Heading textAlign="center" mt="3rem">
        {props.title}
      </Heading>
      <Text color={textColor} textAlign="center" mt=".5rem" mb="1rem">
        Revision: {config.revision}{' '}
        {config.branchName && `/ Branch: ${config.branchName}`}
      </Text>
      <StaffBreadcrumbs paths={props.paths} />
      {props.children}
    </Margin>
  );
};
