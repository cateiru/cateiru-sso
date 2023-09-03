import {
  BreadcrumbItem,
  BreadcrumbLink,
  Breadcrumb as ChakraBreadcrumb,
} from '@chakra-ui/react';
import React from 'react';
import {Link} from '../Common/Next/Link';
import {useSecondaryColor} from '../Common/useColor';

export interface Path {
  href?: string;
  pageName: string;
}

interface Props {
  paths: Path[];
}

export const StaffBreadcrumbs: React.FC<Props> = props => {
  const textColor = useSecondaryColor();

  return (
    <ChakraBreadcrumb
      color={textColor}
      overflow="auto"
      whiteSpace="nowrap"
      py=".5rem"
    >
      {props.paths.map((path, index) => {
        const currentPage = props.paths.length - 1 === index;

        return (
          <BreadcrumbItem
            key={`staff-breadcrumb-${index}`}
            isCurrentPage={currentPage}
          >
            <BreadcrumbLink
              as={!currentPage ? Link : undefined}
              href={!currentPage ? path.href : undefined}
            >
              {path.pageName}
            </BreadcrumbLink>
          </BreadcrumbItem>
        );
      })}
    </ChakraBreadcrumb>
  );
};
