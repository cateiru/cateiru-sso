import {Button} from '@chakra-ui/react';
import React from 'react';
import {hideEmail} from '../../utils/hide';
import {useSecondaryColor} from '../Common/useColor';

export const EmailView: React.FC<{email: string}> = ({email}) => {
  const color = useSecondaryColor();
  const [show, setShow] = React.useState(false);

  const hide = React.useCallback(() => {
    return hideEmail(email);
  }, [email]);

  return (
    <Button
      variant="unstyled"
      fontWeight="400"
      color={color}
      onClick={() => setShow(v => !v)}
      cursor="pointer"
    >
      {show ? email : hide()}
    </Button>
  );
};
