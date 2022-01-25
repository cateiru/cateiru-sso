import {Step, Steps} from 'chakra-ui-steps';

const steps = [
  {label: 'アカウント作成'},
  {label: 'メール確認'},
  {label: 'ユーザ情報登録'},
];

const Flow: React.FC<{step: number}> = ({step}) => {
  return (
    <Steps activeStep={step} colorScheme="blue">
      {steps.map(({label}) => (
        <Step label={label} key={label} />
      ))}
    </Steps>
  );
};

export default Flow;
