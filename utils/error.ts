import {ErrorType} from './types/error';

export class HTTPError extends Error implements ErrorType {
  readonly unique_code: number | undefined;

  constructor(message: string, unique?: number) {
    super(message);
    this.message = message;
    this.unique_code = unique;
  }
}
