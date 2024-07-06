import {
  Box,
  Button,
  SimpleGrid,
  useClipboard,
  useColorModeValue,
} from '@chakra-ui/react';
import React from 'react';
import {TbCheck} from 'react-icons/tb';
import {useSecondaryColor} from '../../../Common/useColor';

interface Props {
  backups: string[];
}

export const OtpBackupTable: React.FC<Props> = props => {
  const backupColor = useSecondaryColor();
  const checkMarkColor = useColorModeValue('#68D391', '#38A169');

  const {onCopy, hasCopied} = useClipboard(props.backups.join(', '));

  return (
    <>
      <SimpleGrid columns={2} spacing="1rem" mb="1rem">
        {props.backups.map((v, i) => {
          return (
            <Box
              key={`backup-${i}-${v}`}
              textAlign="center"
              color={backupColor}
            >
              {v}
            </Box>
          );
        })}
      </SimpleGrid>
      <Button variant="ghost" w="100%" onClick={onCopy}>
        {hasCopied ? (
          <TbCheck size="30px" strokeWidth="3px" color={checkMarkColor} />
        ) : (
          'コピーする'
        )}
      </Button>
    </>
  );
};
