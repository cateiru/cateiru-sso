import {useToast} from '@chakra-ui/react';
import {api} from '../../utils/api';
import {
  ErrorSchema,
  ErrorType,
  ErrorUniqueMessage,
} from '../../utils/types/error';

interface Options {
  errorCallback?: (message: string | undefined) => void;
  customError?: (e: ErrorType) => void;
}

interface Returns {
  request: (data: RequestInit) => Promise<Response | undefined>;
}

export const useRequest = (path: string, options?: Options): Returns => {
  const toast = useToast();

  const request = async (data: RequestInit) => {
    try {
      const res = await fetch(api(path), data);

      if (!res.ok) {
        const error = ErrorSchema.parse(await res.json());

        // カスタムでエラーできる便利機能
        if (options?.customError) {
          options.customError(error);
          return;
        }

        const message = error.unique_code
          ? ErrorUniqueMessage[error.unique_code] ?? error.message
          : error.message;
        toast({
          title: message,
          status: 'error',
          duration: 5000,
          isClosable: true,
        });
        options?.errorCallback?.(message);

        return;
      }

      return res;
    } catch (e) {
      if (e instanceof Error) {
        toast({
          title: 'エラー',
          description: e.message,
          status: 'error',
          duration: 5000,
          isClosable: true,
        });
      }
      options?.errorCallback?.(undefined);
    }
    return;
  };

  return {
    request,
  };
};
