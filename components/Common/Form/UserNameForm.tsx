import {
  FormControl,
  FormErrorMessage,
  FormLabel,
  Input,
  InputGroup,
  InputLeftAddon,
  InputRightElement,
  Spinner,
  Tooltip,
  useColorModeValue,
} from '@chakra-ui/react';
import React from 'react';
import {useFormContext} from 'react-hook-form';
import {TbCheck, TbX} from 'react-icons/tb';
import {UserUserName} from '../../../utils/types/user';
import {useRequest} from '../useRequest';

export interface UserNameFormData {
  user_name: string;
}

interface Props {
  userName: string;
}

export const UserNameForm: React.FC<Props> = ({userName}) => {
  const checkMarkColor = useColorModeValue('#68D391', '#38A169');
  const {
    register,
    clearErrors,
    setError,
    formState: {errors},
  } = useFormContext<UserNameFormData>();

  const [ok, setOk] = React.useState<boolean | null>(null);
  const [loading, setLoading] = React.useState(false);
  const [name, setName] = React.useState('');

  const {request} = useRequest('/v2/user/user_name', {
    errorCallback: () => {
      setOk(false);
      setLoading(false);
    },
  });

  const onSubmit = async () => {};

  const onBlur = async () => {
    if (name === '' || name === userName) {
      setOk(null);
      return;
    }

    setLoading(true);

    const form = new FormData();
    form.append('user_name', name);

    const res = await request({
      method: 'POST',
      mode: 'cors',
      credentials: 'include',
      body: form,
    });

    if (res) {
      const data = UserUserName.safeParse(await res.json());
      if (data.success) {
        setOk(data.data.ok);
        if (data.data.ok) {
          clearErrors('user_name');
        } else {
          setError('user_name', {
            type: 'manual',
            message: 'このユーザー名は既に使用されています',
          });
        }
      } else {
        console.error(data.error);
      }
    }

    setLoading(false);
  };

  const Status = React.useCallback(() => {
    if (loading) {
      return (
        <InputRightElement>
          <Spinner />
        </InputRightElement>
      );
    }

    if (ok === null) return null;

    if (ok) {
      return (
        <Tooltip
          label="このユーザー名は使用可能です"
          hasArrow
          borderRadius="7px"
        >
          <InputRightElement>
            <TbCheck size="30px" strokeWidth="3px" color={checkMarkColor} />
          </InputRightElement>
        </Tooltip>
      );
    }
    return (
      <Tooltip
        label="このユーザー名はすでに使用されています"
        hasArrow
        borderRadius="7px"
      >
        <InputRightElement>
          <TbX size="30px" strokeWidth="3px" color="#E53E3E" />{' '}
        </InputRightElement>
      </Tooltip>
    );
  }, [ok, loading]);

  return (
    <FormControl isInvalid={!!errors.user_name} mt="1rem">
      <FormLabel htmlFor="user_name">ユーザー名</FormLabel>
      <InputGroup size="md">
        <InputLeftAddon>@</InputLeftAddon>
        <Input
          id="user_name"
          type="text"
          autoComplete="username"
          {...register('user_name', {
            required: 'ユーザー名は必須です',
            onChange: e => {
              setName(e.target.value);
              clearErrors('user_name');
            },
            onBlur: onBlur,
          })}
        />
        <Status />
      </InputGroup>
      <FormErrorMessage>
        {errors.user_name && errors.user_name.message}
      </FormErrorMessage>
    </FormControl>
  );
};
