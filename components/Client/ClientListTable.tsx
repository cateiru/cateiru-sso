import {Box, Flex, Text, useColorModeValue} from '@chakra-ui/react';
import React from 'react';
import type {ClientList as ClientListType} from '../../utils/types/client';
import {Avatar} from '../Common/Chakra/Avatar';
import {Link} from '../Common/Next/Link';
import {useSecondaryColor} from '../Common/useColor';

interface Props {
  clients?: ClientListType;
}

export const ClientListTable: React.FC<Props> = ({clients}) => {
  const shadow = useColorModeValue(
    '0px 0px 5px -2px #242838',
    '0px 0px 10px -2px #000000'
  );
  const textColor = useSecondaryColor();

  if (typeof clients === 'undefined') {
    return <></>;
  }

  return (
    <Box>
      {clients.map(v => {
        return (
          <Link
            href={`/client/${v.client_id}`}
            key={`client-list-${v.client_id}`}
          >
            <Box
              w="100%"
              minH="100px"
              boxShadow={shadow}
              borderRadius="10px"
              p="1rem"
              my="1rem"
            >
              <Flex>
                <Avatar src={v.image ?? ''} size="sm" />
                <Text ml=".5rem">
                  <Text fontSize="1.2rem" fontWeight="bold" as="span">
                    {v.name}
                  </Text>
                  <Text as="span" pl=".2rem" color={textColor}>
                    - {new Date(v.created_at).toLocaleString()}
                  </Text>
                </Text>
              </Flex>
              {v.description && (
                <Text as="pre" whiteSpace="pre-wrap">
                  {v.description}
                </Text>
              )}
            </Box>
          </Link>
        );
      })}
    </Box>
  );
};
