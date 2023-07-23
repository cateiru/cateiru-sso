import {
  ButtonGroup,
  Flex,
  FormErrorMessage,
  IconButton,
  Input,
  Select,
} from '@chakra-ui/react';
import React from 'react';
import {useFieldArray, useFormContext} from 'react-hook-form';
import {TbMinus, TbPlus, TbX} from 'react-icons/tb';

export interface ScopesFormValue {
  scopes: {
    value: string;
    isRequired: boolean;
  }[];
}

interface Props {
  scopes: string[];
}

export const ScopesForm: React.FC<Props> = props => {
  const {
    register,
    control,
    formState: {errors},
  } = useFormContext<ScopesFormValue>();

  const {fields, append, remove} = useFieldArray({
    control,
    name: 'scopes',
    rules: {
      maxLength: {
        value: props.scopes.length,
        message: `スコープは最大${props.scopes.length}個まで設定できます`,
      },
      minLength: 1, // openid は必須
    },
  });

  return (
    <>
      {fields.map((field, index) => (
        <React.Fragment key={field.id}>
          <Flex alignItems="center" my=".5rem">
            <Select
              placeholder="スコープを選択"
              {...register(`scopes.${index}.value`, {
                required: {
                  value: true,
                  message: 'スコープを空に設定することはできません',
                },
              })}
              isDisabled={field.isRequired}
            >
              {props.scopes.map(v => {
                return (
                  <option value={v} key={`scopes-${v}`}>
                    {v}
                  </option>
                );
              })}
            </Select>
            {!field.isRequired && (
              <IconButton
                ml=".5rem"
                aria-label="add"
                icon={<TbX size="20px" />}
                borderRadius="50%"
                variant="ghost"
                onClick={() => remove(index)}
              />
            )}
          </Flex>
          <FormErrorMessage>
            {errors.scopes &&
              errors.scopes[index]?.value &&
              errors.scopes[index]?.value?.message}
          </FormErrorMessage>
        </React.Fragment>
      ))}
      <ButtonGroup size="sm" isAttached mt=".5rem">
        <IconButton
          aria-label="add"
          icon={<TbPlus size="20px" />}
          onClick={() => append({value: '', isRequired: false})}
        />
        <IconButton
          aria-label="remove"
          icon={<TbMinus size="20px" />}
          onClick={() => {
            const prevIndex = fields.length - 1;
            const prevElement = fields[prevIndex];
            if (prevElement.isRequired) return;

            remove(prevIndex);
          }}
        />
      </ButtonGroup>
    </>
  );
};
