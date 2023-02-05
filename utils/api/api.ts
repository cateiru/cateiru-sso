import decodeErrorCode from './errorCode';

export class API {
  private apiUrl = process.env.NEXT_PUBLIC_API_URL;
  private config: RequestInit = {};

  private apiVersion = '/v1';

  /**
   * Connect to API
   *
   * @param {string} apiPath -api path.
   * @returns {Promise<Response>} - response data.
   */
  async connect(apiPath: string): Promise<Response> {
    const response = await this.connectNoErr(apiPath);

    if (!response.ok) {
      const resp = await response.json();
      if (resp['code'] !== 1) {
        throw new Error(decodeErrorCode(resp['code']));
      }
      throw new Error(`${resp['status_code']}::ID:${resp['error_id']}`);
    }

    return response;
  }

  async connectNoErr(apiPath: string): Promise<Response> {
    const response = await fetch(
      `${this.apiUrl}${this.apiVersion}${apiPath}`,
      this.config
    );
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

  postForm(data: FormData) {
    this.config = {
      method: 'POST',
      credentials: 'include',
      mode: 'cors',
      headers: {
        // ref. https://stackoverflow.com/questions/39280438/fetch-missing-boundary-in-multipart-form-data-post
        // 意図的にcontent-typeを削除するとboundaryをブラウザが自動で付与します
        // 'Content-Type': 'multipart/form-data',
      },
      body: data,
    };
  }

  postFormURL(data: string) {
    this.config = {
      method: 'POST',
      credentials: 'include',
      mode: 'cors',
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
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
