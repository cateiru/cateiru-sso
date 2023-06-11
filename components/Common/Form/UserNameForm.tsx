import {
  FormControl,
  FormErrorMessage,
  FormLabel,
  Input,
  InputGroup,
  InputLeftAddon,
  InputRightElement,
  useColorModeValue,
} from '@chakra-ui/react';
import React from 'react';
import {useFormContext} from 'react-hook-form';
import {TbCheck, TbX} from 'react-icons/tb';
import {UserUserName} from '../../../utils/types/user';
import {Tooltip} from '../Chakra/Tooltip';
import {Spinner} from '../Icons/Spinner';
import {useRequest} from '../useRequest';

export interface UserNameFormData {
  user_name: string;
}

interface Props {
  userName: string;
}

export const UserNameForm: React.FC<Props> = ({userName}) => {
  const checkMarkSuccessColor = useColorModeValue('#68D391', '#38A169');
  const checkMarkNoSuccessColor = useColorModeValue('#F56565', '#C53030');

  const {
    register,
    clearErrors,
    setError,
    formState: {errors},
  } = useFormContext<UserNameFormData>();

  const [ok, setOk] = React.useState<boolean | null>(null);
  const [message, setMessage] = React.useState('');
  const [loading, setLoading] = React.useState(false);
  const [name, setName] = React.useState('');

  const {request} = useRequest('/v2/user/user_name', {
    errorCallback: () => {
      setOk(false);
      setLoading(false);
    },
  });

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
        setMessage(data.data.message);

        if (data.data.ok) {
          clearErrors('user_name');
        } else {
          setError('user_name', {
            type: 'manual',
            message: data.data.message,
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
        <Tooltip label={message}>
          <InputRightElement>
            <TbCheck
              size="30px"
              strokeWidth="3px"
              color={checkMarkSuccessColor}
            />
          </InputRightElement>
        </Tooltip>
      );
    }
    return (
      <Tooltip label={message}>
        <InputRightElement>
          <TbX size="30px" strokeWidth="3px" color={checkMarkNoSuccessColor} />
        </InputRightElement>
      </Tooltip>
    );
  }, [ok, loading]);

  return (
    <FormControl isInvalid={!!errors.user_name} mt="1rem">
      <FormLabel htmlFor="user_name">ユーザー名</FormLabel>
      <InputGroup size="md">
        <InputLeftAddon fontWeight="bold">@</InputLeftAddon>
        <Input
          id="user_name"
          type="text"
          autoComplete="username"
          {...register('user_name', {
            required: 'ユーザー名は必須です',
            pattern: {
              value: /^[a-zA-Z0-9_]+$/,
              message: "ユーザー名は半角英数字と'_'のみ使用できます",
            },
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
