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

export interface ToURLForm {
  toUrls: {
    url: string;
  }[];
}

const ToURLs = () => {
  const {
    control,
    register,
    formState: {errors},
  } = useFormContext<ToURLForm>();

  const {fields, append, remove} = useFieldArray({
    control,
    name: 'toUrls',
  });

  return (
    <FormControl isInvalid={Boolean(errors.toUrls)} my=".5rem">
      <FormLabel mt="1rem">リダイレクトURL（しない場合はdirect）</FormLabel>
      {fields.map((_, index) => {
        return (
          <Input
            key={`answers_key_${index}`}
            id="fromurls"
            type="text"
            mb=".5rem"
            placeholder={`リダイレクトURL ${index + 1}`}
            {...register(`toUrls.${index}.url` as const, {
              required: 'リダイレクトURL の入力が必要です',
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
        {errors.toUrls && errors.toUrls.message}
      </FormErrorMessage>
    </FormControl>
  );
};

export default ToURLs;
