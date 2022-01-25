import {Box, Heading} from '@chakra-ui/react';
import {useSteps} from 'chakra-ui-steps';
import {useRouter} from 'next/router';
import React from 'react';
import {FieldValues} from 'react-hook-form';
import {useSetRecoilState, useRecoilValue} from 'recoil';
import {useCreateTemp} from '../../hooks/useCreate';
import {CTState, CreateNextState} from '../../utils/state/atom';
import CreateInfo from './CreateInfo';
import Flow from './Flow';
import UserPassword from './UserPassword';
import ValidateMail from './ValidateMai';
import WaitMail from './WaitMailValidate';

export enum CreateType {
  Initialize,
  SendMail,
  ValidateMail,
  Info,
  Error,
}

const SelectCreate: React.FC = () => {
  const {nextStep, setStep, activeStep} = useSteps({
    initialStep: 0,
  });
  const [selectType, setSelectType] = React.useState(CreateType.Initialize);
  const [mailToken, setMailToken] = React.useState('');
  const [mail, setMail] = React.useState('＼(^o^)／');
  const [create, clientToken, err] = useCreateTemp();
  const setCT = useSetRecoilState(CTState);
  const next = useRecoilValue(CreateNextState);

  // クエリパラメータから取得する（あれば）
  const router = useRouter();
  React.useEffect(() => {
    if (!router.isReady) return;
    const query = router.query;

    if (typeof query['m'] === 'string') {
      setMailToken(query['m']);
      setSelectType(CreateType.ValidateMail);
      setStep(1);
    }
  }, [router.isReady, router.query]);

  // clientTokenを移動する
  React.useEffect(() => {
    if (clientToken.length !== 0) {
      setCT(clientToken);
    }
  }, [clientToken]);

  // APIでエラーが発生した場合
  React.useEffect(() => {
    if (err) {
      setSelectType(CreateType.Error);
    }
  }, [err]);

  // メールアドレス確認の次へ行く
  React.useEffect(() => {
    if (next) {
      nextStep();
      setSelectType(CreateType.Info);
    }
  }, [next]);

  const submit = (values: FieldValues, recaptcha: string) => {
    setMail(values.email);
    nextStep();
    setSelectType(CreateType.SendMail);

    // API叩く
    create(values.email, values.password, recaptcha);
  };

  const Validate = () =>
    React.useMemo(() => {
      return <ValidateMail token={mailToken} />;
    }, [mailToken]);

  const Select = () =>
    React.useMemo(() => {
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
              <Validate />
            </>
          );
        case CreateType.Info:
          return (
            <>
              <CreateInfo />
            </>
          );
        case CreateType.Error:
          return <Heading>Oops!</Heading>;
        default:
          return <></>;
      }
    }, [selectType]);

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
