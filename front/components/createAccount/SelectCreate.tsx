import {Center, Heading} from '@chakra-ui/react';
import {useSteps} from 'chakra-ui-steps';
import {useRouter} from 'next/router';
import React from 'react';
import {FieldValues} from 'react-hook-form';
import {useRecoilValue, useRecoilState} from 'recoil';
import {useCreateTemp} from '../../hooks/useCreate';
import useVerify from '../../hooks/useVerify';
import useVerifySurveillance from '../../hooks/useVerifySurveillance';
import {CTState, CreateNextState} from '../../utils/state/atom';
import Spinner from '../common/Spinner';
import CreateInfo from './CreateInfo';
import Flow from './Flow';
import UserPassword from './UserMail';
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
  const [create, err] = useCreateTemp();
  const [surveillance, receive, close] = useVerifySurveillance();
  const [verify, isKeep, loadVerify, verifyError] = useVerify();
  const [next, setNext] = useRecoilState(CreateNextState);
  const ct = useRecoilValue(CTState);

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
      setNext(false);
    }
  }, [next]);

  const submit = (values: FieldValues, recaptcha: string) => {
    setMail(values.email);
    nextStep();
    setSelectType(CreateType.SendMail);

    // API叩く
    create(values.email, recaptcha);
  };

  // 認証確認Websocket
  React.useEffect(() => {
    let unmounted = false;
    if (ct.length !== 0 && !unmounted && selectType === CreateType.SendMail) {
      surveillance(ct);
    }

    return () => {
      unmounted = true;
    };
  }, [ct, selectType]);

  // 確認API叩く
  React.useEffect(() => {
    if (mailToken.length !== 0 && selectType === CreateType.ValidateMail) {
      verify(mailToken);
    }
  }, [mailToken, selectType]);

  React.useEffect(() => {
    // isKeepがtrueの場合は強制的に次へ進む
    if (isKeep) {
      setNext(true);
    }
  }, [isKeep]);

  const Select = () =>
    React.useMemo(() => {
      switch (selectType) {
        case CreateType.Initialize:
          return (
            <>
              <Heading fontSize="1.6rem" marginBottom="1.5rem">
                メールアドレスを認証してアカウントを作成します
              </Heading>
              <UserPassword submit={submit} />
            </>
          );
        case CreateType.SendMail:
          return (
            <>
              <WaitMail mail={mail} receive={receive} close={close} />
            </>
          );
        case CreateType.ValidateMail:
          return (
            <>
              <ValidateMail
                isKeep={isKeep}
                loadVerify={loadVerify}
                verifyError={verifyError}
              />
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
          return (
            <>
              <Spinner />
            </>
          );
      }
    }, [selectType]);

  return (
    <>
      <Center
        width={{base: '90%', md: '50rem'}}
        marginBottom={{base: '1rem', sm: '3rem', md: '4rem'}}
      >
        <Flow step={activeStep} />
      </Center>
      <Select />
    </>
  );
};

export default SelectCreate;
