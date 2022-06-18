import {API} from './api';

export interface OTPGetResponse {
  id: string;
  otp_token: string;
}

interface OTPSetResponse {
  backups: string[];
}

interface OTPBackupsResponse {
  codes: string[];
}

export const getToken = async (): Promise<OTPGetResponse> => {
  const api = new API();

  api.get();

  return (await (await api.connect('/user/otp')).json()) as OTPGetResponse;
};

export const setToken = async (
  id: string,
  passcode: string
): Promise<string[]> => {
  const api = new API();

  api.post(JSON.stringify({type: 'enable', passcode: passcode, id: id}));

  const response = await api.connect('/user/otp');

  return ((await response.json()) as OTPSetResponse).backups;
};

export const getBackups = async (password: string): Promise<string[]> => {
  const api = new API();

  console.log(password);

  api.postFormURL(`password=${encodeURIComponent(password)}`);

  const response = await api.connect('/user/otp/backup');

  return ((await response.json()) as OTPBackupsResponse).codes;
};

export const deleteotp = async (passcode: string) => {
  const api = new API();

  api.post(JSON.stringify({type: 'disable', passcode: passcode}));

  await api.connect('/user/otp');
};
