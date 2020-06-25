'use strict';

class Chat {
    constructor(url, log) {
        this.conn = new WebSocket(url);
        this.log = log;
    }

    onclose(f) {
        this.conn.onclose = f
    }

    onopen(f) {
        this.conn.onopen = f
    }

    onmessage(f) {
        this.conn.onmessage = f
    }

    appendLog(item) {
        let doScroll = this.log.scrollTop > this.log.scrollHeight - this.log.clientHeight - 1;
        this.log.appendChild(item);
        if (doScroll) {
            this.log.scrollTop = this.log.scrollHeight - this.log.clientHeight;
        }
    }
}
