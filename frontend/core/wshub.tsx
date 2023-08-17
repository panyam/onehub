
const normalizeUrl = (url: string) : string => {
  return url.replace("https://", "wss://").replace("http://", "ws://")
}

class WSConn {
  private ws: null | WebSocket = null
  private intervalId: any = null
  private handlers = [] as any[]
  constructor(public readonly url: string) {
    url = normalizeUrl(url)
  }

  start() {
    if (this.ws != null) {
      console.log("Websocket connection already established to url: ", this.url)
      return
    }
    this.ws = new WebSocket(this.url);
    const ws = this.ws
    ws.onerror = (event) => {
      console.log("WS Error: ", event);
    };
    ws.onclose = (event) => {
      console.log("WS Closed: ", event);
      if (this.intervalId != null) {
        clearInterval(this.intervalId);
      }
    };
    ws.onmessage = (event) => {
      const msg = JSON.parse(event.data);
      for (const handler of this.handlers) {
        handler(msg);
      }
    };

    this.intervalId = setInterval(() => {
      console.log("Sending ping", new Date());
      ws.send(JSON.stringify({ type: "ping" }));
    }, 5000);
    return ws;
  }

  addHandler(handler: (msg: any) => void) {
    if (this.handlers.indexOf(handler) < 0) {
      this.handlers.push(handler)
    }
  }

  removeHandler(handler: (msg: any) => void) {
    const index = this.handlers.indexOf(handler)
    if (index >= 0) {
      this.handlers.splice(index, 1)
    }
  }
}

/**
 * A way to have a message hub architecture for many consumers
 * needing events to be sent to them from different WS hosts.
 * This way clients do not need to worry about retries, or managing
 * websocket connections etc.
 */
class WebSocketHub {
  wsConns = new Map<string, WSConn>()
  getConn(url: string, ensure = false): WSConn | null {
    url = normalizeUrl(url)
    if (this.wsConns.has(url)) {
      // already exists - so let it run
      return this.wsConns.get(url) || null
    } else {
      if (ensure) {
        const wsconn = new WSConn(url)
        this.wsConns.set(url, wsconn)
        wsconn.start()
        return wsconn
      } else {
        return null 
      }
    }
  }

  connect(url: string) {
    return this.getConn(url, true)
  }

  addHandler(url: string, handler: (msg: any) => void) {
    const wsconn = this.getConn(url, true)
    wsconn!.addHandler(handler)
  }
}
