import {
  Flex,
  FormControl,
  FormErrorMessage,
  FormLabel,
  Input,
  Switch,
} from '@chakra-ui/react';
import React from 'react';
import {useFormContext} from 'react-hook-form';

export interface NameFormData {
  family_name?: string;
  middle_name?: string;
  given_name?: string;
}

interface Props {
  isMiddleName: boolean;
}

export const NameForm: React.FC<Props> = props => {
  const {
    register,
    formState: {errors},
    setValue,
  } = useFormContext<NameFormData>();

  const [middleName, setMiddleName] = React.useState('');
  const [isMiddleName, setIsMiddleName] = React.useState(props.isMiddleName);

  const onChange = (v: boolean) => {
    if (v) {
      if (middleName !== '') {
        setValue('middle_name', middleName);
      }
    } else {
      setValue('middle_name', undefined);
    }

    setIsMiddleName(v);
  };

  return (
    <>
      <Flex mt="1rem">
        <FormControl isInvalid={!!errors.family_name} mr=".5rem">
          <FormLabel htmlFor="family_name">姓</FormLabel>
          <Input
            id="family_name"
            type="text"
            autoComplete="family-name"
            {...register('family_name')}
          />
        </FormControl>
        {isMiddleName && (
          <FormControl isInvalid={!!errors.middle_name} mr=".5rem">
            <FormLabel htmlFor="middle_name">ミドルネーム</FormLabel>
            <Input
              id="middle_name"
              autoComplete="additional-name"
              {...register('middle_name', {
                onChange: v => setMiddleName(v.target.value),
              })}
            />
          </FormControl>
        )}
        <FormControl isInvalid={!!errors.given_name}>
          <FormLabel htmlFor="given_name">名</FormLabel>
          <Input
            id="given_name"
            autoComplete="given-name"
            {...register('given_name')}
          />
        </FormControl>
      </Flex>
      <FormErrorMessage>
        {errors.family_name && errors.family_name.message}
        {errors.middle_name && errors.middle_name.message}
        {errors.given_name && errors.given_name.message}
      </FormErrorMessage>
      <FormControl display="flex" alignItems="center" mt=".5rem">
        <FormLabel htmlFor="is_middle_name" mb="0">
          ミドルネーム
        </FormLabel>
        <Switch
          id="is_middle_name"
          colorScheme="cateiru"
          isChecked={isMiddleName}
          onChange={v => onChange(v.target.checked)}
        />
      </FormControl>
    </>
  );
};
