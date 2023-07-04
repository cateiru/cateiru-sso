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
  request: (
    data: RequestInit,
    urlParams?: URLSearchParams
  ) => Promise<Response | undefined>;
}

export const useRequest = (path: string, options?: Options): Returns => {
  const toast = useToast();

  const request = async (data: RequestInit, urlParams?: URLSearchParams) => {
    try {
      const res = await fetch(api(path, urlParams), data);

      if (!res.ok) {
        let message: string;
        if (res.status >= 500 && res.status < 600) {
          // サーバーエラーの場合はそのまま表示する
          message = await res.text();
        } else {
          // クライアントエラーの場合はレスポンスからエラーを読む
          const error = ErrorSchema.parse(await res.json());

          // カスタムでエラーできる便利機能
          if (options?.customError) {
            options.customError(error);
            return;
          }

          message = error.unique_code
            ? ErrorUniqueMessage[error.unique_code] ?? error.message
            : error.message;
        }

        toast({
          title: message,
          status: 'error',
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
