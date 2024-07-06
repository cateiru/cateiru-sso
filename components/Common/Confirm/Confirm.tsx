import {
  Button,
  ButtonGroup,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalHeader,
  ModalOverlay,
} from '@chakra-ui/react';
import React, {useState} from 'react';

export interface ConfirmText {
  confirmHeader: string;
  confirmOkText: string;
  confirmOkTextColor?: string;
  confirmCancelText?: string;
}

interface Props {
  children?: React.ReactNode;
  onSubmit: () => Promise<void> | void;
  text: ConfirmText;
  isOpen: boolean;
  onClose: () => void;
}

export const Confirm: React.FC<Props> = props => {
  const [loading, setLoading] = useState(false);

  // ロードが終了したら閉じる
  React.useEffect(() => {
    if (!loading) {
      onCloseHandler();
    }
  }, [loading]);

  const onCloseHandler = () => {
    if (loading) return;
    props.onClose();
  };

  const onSubmitHandler = () => {
    const f = async () => {
      setLoading(true);
      await props.onSubmit();
      setLoading(false);
    };
    f();
  };

  return (
    <>
      <Modal isOpen={props.isOpen} onClose={onCloseHandler} isCentered>
        <ModalOverlay />
        <ModalContent>
          <ModalHeader maxW="90%">{props.text.confirmHeader}</ModalHeader>
          <ModalCloseButton size="lg" />
          <ModalBody>
            {props.children}
            <ButtonGroup w="100%" mt="2rem" mb=".5rem">
              <Button
                w="100%"
                colorScheme={props.text.confirmOkTextColor ?? 'cateiru'}
                onClick={onSubmitHandler}
                isLoading={loading}
              >
                {props.text.confirmOkText}
              </Button>
              {props.text.confirmCancelText && (
                <Button w="100%" onClick={onCloseHandler} isDisabled={loading}>
                  {props.text.confirmCancelText}
                </Button>
              )}
            </ButtonGroup>
          </ModalBody>
        </ModalContent>
      </Modal>
    </>
  );
};
