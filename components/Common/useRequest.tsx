import {useToast} from '@chakra-ui/react';
import {api} from '../../utils/api';
import {ErrorSchema, ErrorUniqueMessage} from '../../utils/types/error';

interface Options {
  errorCallback?: () => void;
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
        toast({
          title: ErrorUniqueMessage[error.unique_code] ?? error.message,
          status: 'error',
          duration: 5000,
          isClosable: true,
        });
        options?.errorCallback?.();

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
      options?.errorCallback?.();
    }
    return;
  };

  return {
    request,
  };
};
