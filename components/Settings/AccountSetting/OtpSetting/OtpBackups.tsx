import {Box, Divider, Text} from '@chakra-ui/react';
import React from 'react';
import {useSecondaryColor} from '../../../Common/useColor';
import {OtpBackupTable} from './OtpBackupTable';

interface Props {
  backups: string[];
  title?: string;
}

export const OtpBackups: React.FC<Props> = props => {
  const descriptionColor = useSecondaryColor();

  return (
    <Box>
      {props.title && (
        <Text fontSize="1.5rem" fontWeight="bold" textAlign="center" mb=".5rem">
          {props.title}
        </Text>
      )}
      <Text color={descriptionColor} textAlign="center" mb=".5rem">
        バックアップコードは、印刷もしくはメモなどに残して安全な場所に保管してください。
      </Text>
      <Divider my="1rem" />
      <OtpBackupTable backups={props.backups} />
    </Box>
  );
};
