import {API} from './api';

/**
 * Pingを出す
 * APIはCloud Runで動いており、アクセスごとにインスタンスが立つため初期リクエスト時にレスポンス時間が多くなる
 * それを解決するため、cookieを変えるリクエストの前にpingを出してリクエスト中にブラウザバックしてもcookieが消滅する
 * ことがなくなる。はず。
 */
export default async function ping(): Promise<void> {
  const api = new API();
  api.get();

  await api.connect('/');
}
