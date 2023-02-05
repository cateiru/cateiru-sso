/**********************************************************
 * connect for websocket
 *
 * @author Yuto Watanabe <yuto.w51942@gmail.com>
 * @version 1.0.0
 *
 * Copyright (C) 2021 hello-slide
 **********************************************************/

export default class Socket {
  protected url: string;

  protected socket: WebSocket;

  constructor(path: string) {
    const domain = process.env.NEXT_PUBLIC_WS_API_URL;

    this.url = `${domain}${path}`;
    this.socket = new WebSocket(this.url);
  }

  public initSend() {
    this.socket.onopen = () => {};
  }

  /**
   * Get the data.
   *
   * @param {(data: string) => void} fn - response func.
   */
  public get(fn: (data: string) => void) {
    this.socket.onmessage = (event: MessageEvent<string>) => {
      fn(event.data);
    };
  }

  /**
   *  error handling
   *
   * @param {(ev: Event) => void} fn - error call func
   */
  public error(fn: (ev: Event) => void) {
    this.socket.onerror = fn;
  }

  /**
   * close event.
   *
   * @param {() => void} fn close handler.
   */
  public end(fn: () => void) {
    this.socket.onclose = fn;
  }

  /**
   * Close socket.
   */
  public close() {
    this.socket.close();
  }
}
