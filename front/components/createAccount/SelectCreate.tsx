import {Box} from '@chakra-ui/react';
import {useSteps} from 'chakra-ui-steps';
import React from 'react';
import {GoogleReCaptchaProvider} from 'react-google-recaptcha-v3';
import {FieldValues} from 'react-hook-form';
import Flow from './Flow';
import UserPassword from './UserPassword';
import ValidateMail from './ValidateMai';
import WaitMail from './WaitMailValidate';

const reCAPTCHA = process.env.NEXT_PUBLIC_RE_CAPTCHA;

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

  React.useEffect(() => {
    if (selectType === CreateType.ValidateMail) {
      setStep(1);
    }
  }, [selectType]);

  const submit = (values: FieldValues, recaptcha: string) => {
    setMail(values.email);
    nextStep();
    setSelectType(CreateType.SendMail);

    console.log(recaptcha);
  };

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
