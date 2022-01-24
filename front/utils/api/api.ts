import decodeErrorCode from './errorCode';

export class API {
  private apiUrl = process.env.NEXT_PUBLIC_API_URL;
  private config: RequestInit = {};

  /**
   * Connect to API
   *
   * @param {string} apiPath -api path.
   * @returns {Promise<Response>} - response data.
   */
  async connect(apiPath: string): Promise<Response> {
    const response = await fetch(`${this.apiUrl}${apiPath}`, this.config);

    if (!response.ok) {
      const resp = await response.json();
      if (resp['code']) {
        throw new Error(decodeErrorCode(resp['code']));
      }
      throw new Error(`${resp['status_code']}::ID:${resp['error_id']}`);
    }

    return response;
  }

  post(data: string) {
    this.config = {
      method: 'POST',
      credentials: 'include',
      mode: 'cors',
      headers: {
        'Content-Type': 'application/json',
      },
      body: data,
    };
  }

  get() {
    this.config = {
      method: 'GET',
      credentials: 'include',
      mode: 'cors',
    };
  }

  delete() {
    this.config = {
      method: 'DELETE',
      credentials: 'include',
      mode: 'cors',
    };
  }
}
