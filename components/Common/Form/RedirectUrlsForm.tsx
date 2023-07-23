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

export interface RedirectUrlsFormValue {
  redirectUrls: {value: string}[];
}

interface Props {
  maxCreatedCount: number;
}

export const RedirectUrlsForm: React.FC<Props> = props => {
  const {
    register,
    control,
    formState: {errors},
  } = useFormContext<RedirectUrlsFormValue>();

  const {fields, remove, append} = useFieldArray({
    control,
    name: 'redirectUrls',
  });

  return (
    <>
      {fields.map((field, index) => (
        <React.Fragment key={field.id}>
          <Flex key={field.id} alignItems="center" my=".5rem">
            <Input
              type="url"
              placeholder="https://"
              {...register(`redirectUrls.${index}.value`, {
                required: {
                  value: true,
                  message: 'リダイレクトURLを空に設定することはできません',
                },
              })}
            />
            {fields.length > 1 && (
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
            )}
          </Flex>
          <FormErrorMessage>
            {errors.redirectUrls &&
              errors.redirectUrls[index]?.value &&
              errors.redirectUrls[index]?.value?.message}
          </FormErrorMessage>
        </React.Fragment>
      ))}
      <ButtonGroup size="sm" isAttached>
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
          onClick={() => {
            if (fields.length <= 1) return;

            remove(fields.length - 1);
          }}
        />
      </ButtonGroup>
    </>
  );
};
