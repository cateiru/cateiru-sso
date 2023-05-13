import {
  Button,
  Flex,
  FormControl,
  FormLabel,
  Input,
  Select,
} from '@chakra-ui/react';
import React from 'react';
import {FormProvider, useForm} from 'react-hook-form';
import {useRecoilValue} from 'recoil';
import {UserState} from '../../utils/state/atom';
import {NameForm, NameFormData} from '../Common/Form/NameForm';
import {UserNameForm, UserNameFormData} from '../Common/Form/UserNameForm';

interface ProfileFormData extends UserNameFormData, NameFormData {
  gender: string;
  birthdate?: Date;
}

export const ProfileForm = () => {
  const user = useRecoilValue(UserState);
  const methods = useForm<ProfileFormData>({
    defaultValues: {
      user_name: user?.user.user_name ?? '',
      family_name: user?.user.family_name ?? undefined,
      middle_name: user?.user.middle_name ?? undefined,
      given_name: user?.user.given_name ?? undefined,
      gender: user?.user.gender,
      birthdate: user?.user.birthdate ?? undefined,
    },
  });
  const {
    handleSubmit,
    register,
    formState: {errors, isSubmitting},
  } = methods;

  const onSubmit = async (data: ProfileFormData) => {
    console.log(JSON.stringify(data));
  };

  return (
    <FormProvider {...methods}>
      <form onSubmit={handleSubmit(onSubmit)}>
        <UserNameForm userName={user?.user.user_name ?? ''} />
        <NameForm isMiddleName={!!user?.user.middle_name} />
        <Flex mt="1rem">
          <FormControl isInvalid={!!errors.family_name} mr=".5rem">
            <FormLabel htmlFor="family_name">性別</FormLabel>
            <Select placeholder="性別を選択してください">
              <option value="1">男性</option>
              <option value="2">女性</option>
              <option value="9">その他</option>
            </Select>
          </FormControl>
          <FormControl isInvalid={!!errors.birthdate}>
            <FormLabel htmlFor="birthdate">誕生日</FormLabel>
            <Input id="birthdate" type="date" {...register('birthdate')} />
          </FormControl>
        </Flex>

        <Button
          mt="1rem"
          isLoading={isSubmitting}
          colorScheme="cateiru"
          type="submit"
          w="100%"
        >
          プロフィールを更新する
        </Button>
      </form>
    </FormProvider>
  );
};
