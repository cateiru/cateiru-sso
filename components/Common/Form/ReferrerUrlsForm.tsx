import {
  ButtonGroup,
  Flex,
  FormErrorMessage,
  IconButton,
  Input,
} from '@chakra-ui/react';
import React from 'react';
import {useFieldArray, useFormContext} from 'react-hook-form';
import {TbMinus, TbPlus, TbX} from 'react-icons/tb';

export interface ReferrerUrlsFormValue {
  referrerUrls: {value: string}[];
}

interface Props {
  maxCreatedCount: number;
}

export const ReferrerUrlsForm: React.FC<Props> = props => {
  const {
    register,
    control,
    formState: {errors},
  } = useFormContext<ReferrerUrlsFormValue>();

  const {fields, remove, append} = useFieldArray({
    control,
    name: 'referrerUrls',
    rules: {
      minLength: {
        value: 1,
        message: 'リダイレクトURLは最低でも1つは設定する必要があります',
      },
      maxLength: {
        value: props.maxCreatedCount,
        message: `リダイレクトURLは最大${props.maxCreatedCount}個まで設定できます`,
      },
    },
  });

  return (
    <>
      {fields.map((field, index) => (
        <React.Fragment key={field.id}>
          <Flex key={field.id} alignItems="center" my=".5rem">
            <Input
              type="url"
              placeholder="https://"
              {...register(`referrerUrls.${index}.value`, {
                required: {
                  value: true,
                  message: 'リダイレクトURLを空に設定することはできません',
                },
              })}
            />
            <IconButton
              ml=".5rem"
              aria-label="add"
              icon={<TbX size="20px" />}
              borderRadius="50%"
              variant="ghost"
              onClick={() => {
                remove(index);
              }}
            />
          </Flex>
          <FormErrorMessage>
            {errors.referrerUrls &&
              errors.referrerUrls[index]?.value &&
              errors.referrerUrls[index]?.value?.message}
          </FormErrorMessage>
        </React.Fragment>
      ))}
      <ButtonGroup size="sm" isAttached mt=".5rem">
        <IconButton
          aria-label="add"
          icon={<TbPlus size="20px" />}
          onClick={() => {
            if (fields.length >= props.maxCreatedCount) return;

            append({value: ''});
          }}
        />
        <IconButton
          aria-label="remove"
          icon={<TbMinus size="20px" />}
          onClick={() => remove(fields.length - 1)}
        />
      </ButtonGroup>
    </>
  );
};
