import {useToast} from '@chakra-ui/react';
import {useAtom} from 'jotai';
import {UserState} from '../../utils/state/atom';
import {UserSchema} from '../../utils/types/user';
import {type NameFormData} from '../Common/Form/NameForm';
import {type UserNameFormData} from '../Common/Form/UserNameForm';
import {useRequest} from '../Common/useRequest';

export interface ProfileFormData extends UserNameFormData, NameFormData {
  gender?: string;
  birthdate?: string;
}

interface Returns {
  updateProfile: (data: ProfileFormData) => Promise<void>;
}

export const useUpdateProfile = (): Returns => {
  const [user, setUser] = useAtom(UserState);
  const toast = useToast();
  const {request} = useRequest('/v2/user/');

  const updateProfile = async (data: ProfileFormData) => {
    const form = new FormData();

    // ユーザー名は更新した & 今のユーザーと違う場合のみ更新する
    if (data.user_name && data.user_name !== user?.user.user_name) {
      form.append('user_name', data.user_name);
    }

    // 名前
    if (data.family_name) form.append('family_name', data.family_name);
    if (data.middle_name) form.append('middle_name', data.middle_name);
    if (data.given_name) form.append('given_name', data.given_name);

    form.append('gender', data.gender ? data.gender : '0');

    if (data.birthdate) form.append('birth_date', data.birthdate);

    const res = await request({
      method: 'PUT',
      body: form,
      mode: 'cors',
      credentials: 'include',
    });

    if (res) {
      const data = UserSchema.safeParse(await res.json());
      if (data.success) {
        if (user) {
          setUser({
            ...user,
            user: data.data,
          });
        }

        toast({
          title: 'プロフィールを更新しました',
          status: 'success',
        });
      } else {
        console.error(data.error);
      }
    }
  };

  return {updateProfile};
};
