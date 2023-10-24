import {useSearchParams} from 'next/navigation';

export const useGetOauthLoginSession = () => {
  const searchparams = useSearchParams();

  const getOauthLoginSession = (): HeadersInit => {
    const oauthLoginSession = searchparams.get('oauth_login_session');

    if (oauthLoginSession === null) {
      return {};
    }

    return {
      'X-Oauth-Login-Session': oauthLoginSession,
    };
  };

  return getOauthLoginSession;
};
