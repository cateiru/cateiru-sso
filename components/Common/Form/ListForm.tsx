import {ButtonGroup, Flex, IconButton, Input} from '@chakra-ui/react';
import React from 'react';
import {useFieldArray, useFormContext} from 'react-hook-form';
import {TbMinus, TbPlus, TbX} from 'react-icons/tb';

interface Props {
  name: string;
  placeholder?: string;
}

export const ListForm: React.FC<Props> = ({name, placeholder}) => {
  const {register, control} = useFormContext();

  const {fields, append, remove} = useFieldArray({
    control,
    name: name,
  });

  return (
    <>
      {fields.map((field, index) => (
        <Flex key={field.id} alignItems="center" my=".5rem">
          <Input placeholder={placeholder} {...register(`${name}.${index}`)} />
          <IconButton
            ml=".5rem"
            aria-label="add"
            icon={<TbX size="20px" />}
            borderRadius="50%"
            variant="ghost"
            onClick={() => remove(index)}
          />
        </Flex>
      ))}
      <ButtonGroup size="sm" isAttached>
        <IconButton
          aria-label="add"
          icon={<TbPlus size="20px" />}
          onClick={() => append('')}
        />
        <IconButton
          aria-label="remove"
          icon={<TbMinus size="20px" />}
          onClick={() => {
            remove(fields.length - 1);
          }}
        />
      </ButtonGroup>
    </>
  );
};
