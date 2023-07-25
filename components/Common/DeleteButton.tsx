import {IconButton, useColorModeValue, useDisclosure} from '@chakra-ui/react';
import React from 'react';
import {Tooltip} from './Chakra/Tooltip';
import {Confirm, ConfirmText} from './Confirm/Confirm';

interface Props {
  tooltipLabel: string;
  onSubmit: () => Promise<void> | void;
  text: ConfirmText;
  children?: React.ReactNode;
  icon: React.ReactElement;
}

export const DeleteButton: React.FC<Props> = props => {
  const defaultTrashColor = useColorModeValue('#CBD5E0', '#4A5568');
  const hoverTrashColor = useColorModeValue('#F56565', '#C53030');

  const confirmModal = useDisclosure();

  return (
    <>
      <Tooltip label={props.tooltipLabel} placement="top">
        <IconButton
          icon={props.icon}
          aria-label="delte"
          variant="unstyled"
          color={defaultTrashColor}
          _hover={{color: hoverTrashColor}}
          onClick={confirmModal.onOpen}
          minW="0"
          w="auto"
          h="auto"
          m="0"
        />
      </Tooltip>
      <Confirm
        text={props.text}
        onSubmit={props.onSubmit}
        isOpen={confirmModal.isOpen}
        onClose={confirmModal.onClose}
      >
        {props.children}
      </Confirm>
    </>
  );
};
