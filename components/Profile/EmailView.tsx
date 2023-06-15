import {
  Flex,
  IconButton,
  Spacer,
  Text,
  useColorModeValue,
} from '@chakra-ui/react';
import React from 'react';
import {TbEye, TbEyeOff} from 'react-icons/tb';
import {hideEmail} from '../../utils/hide';

export const EmailView: React.FC<{email: string}> = ({email}) => {
  const color = useColorModeValue('#718096', '#A0AEC0');
  const [show, setShow] = React.useState(false);

  const hide = React.useCallback(() => {
    return hideEmail(email);
  }, [email]);

  return (
    <Flex w="150px" alignItems="center">
      <IconButton
        mr=".1rem"
        icon={
          show ? (
            <TbEye size="20px" color={color} />
          ) : (
            <TbEyeOff size="20px" color={color} />
          )
        }
        aria-label={show ? 'hide' : 'show'}
        onClick={() => setShow(v => !v)}
        size="sm"
        variant="ghost"
      />
      <Spacer />
      <Text color={color}>{show ? email : hide()}</Text>
    </Flex>
  );
};
