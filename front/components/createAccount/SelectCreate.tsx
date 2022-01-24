import {Box} from '@chakra-ui/react';
import {useSteps} from 'chakra-ui-steps';
import React from 'react';
import {FieldValues} from 'react-hook-form';
import {useCreateTemp} from '../../hooks/useCreate';
import Flow from './Flow';
import UserPassword from './UserPassword';
import ValidateMail from './ValidateMai';
import WaitMail from './WaitMailValidate';

export enum CreateType {
  Initialize,
  SendMail,
  ValidateMail,
}

export interface SelectProps {
  selectType: CreateType;
  setSelectType: React.Dispatch<React.SetStateAction<CreateType>>;
  mailToken: string;
}

const SelectCreate: React.FC<SelectProps> = ({
  selectType,
  setSelectType,
  mailToken,
}) => {
  const {nextStep, setStep, activeStep} = useSteps({
    initialStep: 0,
  });
  const [mail, setMail] = React.useState('＼(^o^)／');
  const [create, clientToken] = useCreateTemp();

  React.useEffect(() => {
    if (selectType === CreateType.ValidateMail) {
      setStep(1);
    }
  }, [selectType]);

  const submit = (values: FieldValues, recaptcha: string) => {
    setMail(values.email);
    nextStep();
    setSelectType(CreateType.SendMail);

    // API叩く
    create(values.email, values.password, recaptcha);

    console.log(recaptcha);
  };

  // TODO: clientTokenでws
  React.useEffect(() => {
    console.log(clientToken);
  }, [clientToken]);

  const Select = () => {
    switch (selectType) {
      case CreateType.Initialize:
        return (
          <>
            <UserPassword submit={submit} />
          </>
        );
      case CreateType.SendMail:
        return (
          <>
            <WaitMail mail={mail} />
          </>
        );
      case CreateType.ValidateMail:
        return (
          <>
            <ValidateMail token={mailToken} />
          </>
        );
      default:
        return <></>;
    }
  };

  return (
    <>
      <Box width={{base: '100%', md: '50rem'}} marginBottom="4rem">
        <Flow step={activeStep} />
      </Box>
      <Select />
    </>
  );
};

export default SelectCreate;
