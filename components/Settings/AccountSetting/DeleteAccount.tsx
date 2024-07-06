import {Button} from '@chakra-ui/react';
import {SettingCard} from '../SettingCard';

export const DeleteAccount = () => {
  return (
    <SettingCard title="アカウント削除">
      <Button w="100%">アカウント削除</Button>
    </SettingCard>
  );
};
