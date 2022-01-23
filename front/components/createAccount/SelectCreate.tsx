import {Box} from '@chakra-ui/react';
import {useSteps} from 'chakra-ui-steps';
import React from 'react';
import {useGoogleReCaptcha} from 'react-google-recaptcha-v3';
import {FieldValues} from 'react-hook-form';
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
  const [recaptcha, setRecaptcha] = React.useState('');

  React.useEffect(() => {
    if (selectType === CreateType.ValidateMail) {
      setStep(1);
    }
  }, [selectType]);

  const submit = (values: FieldValues) => {
    setMail(values.email);
    nextStep();
    setSelectType(CreateType.SendMail);
  };

  const {executeRecaptcha} = useGoogleReCaptcha();
  const handleReCaptchaVerify = React.useCallback(async () => {
    if (!executeRecaptcha) {
      console.log('Execute recaptcha not yet available');
      return;
    }

    const token = await executeRecaptcha();
    setRecaptcha(token);
  }, [executeRecaptcha, setRecaptcha]);

  // reCAPTCHAのトークンを取得する
  React.useEffect(() => {
    if (selectType === CreateType.Initialize) {
      handleReCaptchaVerify();
    }
  }, [executeRecaptcha, selectType]);

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
