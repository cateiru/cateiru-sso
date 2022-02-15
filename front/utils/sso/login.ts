import {ParsedUrlQuery} from 'querystring';

export interface OIDCRequestQuery {
  scopes: string[];
  responseType: string[];
  clientID: string;
  redirectURL: string;
  state: string;
  prompt: string;
}

export class Login {
  private scopes: string[];
  private responseType: string[];
  private clientID: string;
  private redirectURL: string;
  private state: string;
  private prompt: string;

  constructor(query: ParsedUrlQuery) {
    const scope = query.scope;
    const responseType = query.response_type;
    const clientID = query.client_id;
    const redirectURL = query.redirect_uri;

    const state = query.state;
    const prompt = query.prompt;

    this.scopes = this.parseMulti(scope);
    this.responseType = this.parseMulti(responseType);
    this.clientID = this.parseQuery(clientID);
    this.redirectURL = this.parseQuery(redirectURL);
    this.state = this.parseQuery(state);
    this.prompt = this.parseQuery(prompt);
  }

  public require(): boolean {
    if (!this.scopes.includes('openid')) {
      return false;
    }

    if (
      !(this.responseType.length === 1 && this.responseType.includes('code'))
    ) {
      return false;
    }

    if (this.redirectURL === '') {
      return false;
    }

    return true;
  }

  public parse(): OIDCRequestQuery {
    return {
      scopes: this.scopes,
      responseType: this.responseType,
      clientID: this.clientID,
      redirectURL: this.redirectURL,
      state: this.state,
      prompt: this.prompt,
    };
  }

  private parseQuery(q: string | string[] | undefined): string {
    if (typeof q === 'undefined') {
      return '';
    } else if (typeof q === 'string') {
      return q;
    }
    return q.join('');
  }

  private parseMulti(scope: string | string[] | undefined): string[] {
    if (typeof scope === 'undefined') {
      return [];
    } else if (typeof scope === 'string') {
      return scope.split(' ');
    }
    return scope;
  }
}
