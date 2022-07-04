import {
  FormControl,
  Input,
  FormErrorMessage,
  FormLabel,
  ButtonGroup,
  IconButton,
} from '@chakra-ui/react';
import {useFormContext, useFieldArray} from 'react-hook-form';
import {IoAddOutline, IoRemoveOutline} from 'react-icons/io5';

export interface FromURLForm {
  fromUrls: {
    url: string;
  }[];
}

const FromURLs = () => {
  const {
    control,
    register,
    formState: {errors},
  } = useFormContext<FromURLForm>();

  const {fields, append, remove} = useFieldArray({
    control,
    name: 'fromUrls',
  });

  return (
    <FormControl isInvalid={Boolean(errors.fromUrls)} my=".5rem">
      <FormLabel mt="1rem">送信元URL</FormLabel>
      {fields.map((_, index) => {
        return (
          <Input
            key={`answers_key_${index}`}
            id="fromurls"
            type="text"
            mb=".5rem"
            placeholder={`送信元URL ${index + 1}`}
            {...register(`fromUrls.${index}.url` as const, {
              required: '送信元URL の入力が必要です',
              pattern: {
                value: /(https:\/\/[\w/:%#$&?()~.=+-]+|http:\/\/.+|direct)/,
                message: 'URLの形式が違うようです',
              },
            })}
          />
        );
      })}
      <ButtonGroup size="sm" isAttached mt=".5rem">
        <IconButton
          aria-label="Add"
          icon={<IoAddOutline />}
          onClick={() => {
            if (fields.length < 5) {
              append({url: ''});
            }
          }}
        />
        <IconButton
          aria-label="Remove"
          icon={<IoRemoveOutline />}
          onClick={() => {
            if (fields.length > 1) {
              remove(-1);
            }
          }}
        />
      </ButtonGroup>
      <FormErrorMessage>
        {errors.fromUrls && errors.fromUrls.message}
      </FormErrorMessage>
    </FormControl>
  );
};

export default FromURLs;
