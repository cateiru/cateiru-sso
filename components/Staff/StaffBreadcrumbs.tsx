import {
  BreadcrumbItem,
  BreadcrumbLink,
  Breadcrumb as ChakraBreadcrumb,
  useColorModeValue,
} from '@chakra-ui/react';
import React from 'react';
import {Link} from '../Common/Next/Link';

export interface Path {
  href?: string;
  pageName: string;
}

interface Props {
  paths: Path[];
}

export const StaffBreadcrumbs: React.FC<Props> = props => {
  const textColor = useColorModeValue('gray.500', 'gray.400');

  return (
    <ChakraBreadcrumb color={textColor}>
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
