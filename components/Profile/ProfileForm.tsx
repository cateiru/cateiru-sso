import {
  Button,
  Flex,
  FormControl,
  FormLabel,
  Input,
  Select,
} from '@chakra-ui/react';
import {format} from 'date-fns';
import React from 'react';
import {FormProvider, useForm} from 'react-hook-form';
import {useRecoilValue} from 'recoil';
import {UserState} from '../../utils/state/atom';
import {NameForm} from '../Common/Form/NameForm';
import {UserNameForm} from '../Common/Form/UserNameForm';
import {ProfileFormData, useUpdateProfile} from './useUpdateProfile';

export const ProfileForm = () => {
  const user = useRecoilValue(UserState);
  const {updateProfile} = useUpdateProfile();
  const methods = useForm<ProfileFormData>({
    defaultValues: {
      user_name: user?.user.user_name ?? '',
      family_name: user?.user.family_name ?? undefined,
      middle_name: user?.user.middle_name ?? undefined,
      given_name: user?.user.given_name ?? undefined,
      gender: user?.user.gender,
      birthdate: user?.user.birthdate
        ? format(new Date(user.user.birthdate), 'yyyy-MM-dd')
        : undefined,
    },
  });
  const {
    handleSubmit,
    register,
    setValue,
    formState: {errors, isSubmitting},
  } = methods;

  React.useEffect(() => {
    if (user) {
      setValue('user_name', user.user.user_name);
      setValue('family_name', user.user.family_name ?? undefined);
      setValue('middle_name', user.user.middle_name ?? undefined);
      setValue('given_name', user.user.given_name ?? undefined);
      setValue('gender', user.user.gender);
      setValue(
        'birthdate',
        user?.user.birthdate
          ? format(new Date(user.user.birthdate), 'yyyy-MM-dd')
          : undefined
      );
    }
  }, [user]);

  return (
    <FormProvider {...methods}>
      <form onSubmit={handleSubmit(updateProfile)}>
        <UserNameForm userName={user?.user.user_name ?? ''} />
        <NameForm isMiddleName={!!user?.user.middle_name} />
        <Flex mt="1rem">
          <FormControl isInvalid={!!errors.gender} mr=".5rem">
            <FormLabel htmlFor="gender">性別</FormLabel>
            <Select autoComplete="sex" id="gender" {...register('gender')}>
              <option value="0">未設定</option>
              <option value="1">男性</option>
              <option value="2">女性</option>
              <option value="9">その他</option>
            </Select>
          </FormControl>
          <FormControl isInvalid={!!errors.birthdate}>
            <FormLabel htmlFor="birthdate">誕生日</FormLabel>
            <Input
              id="birthdate"
              type="date"
              autoComplete="bday"
              {...register('birthdate')}
            />
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
