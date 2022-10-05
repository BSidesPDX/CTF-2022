/*!
* Start Bootstrap - Coming Soon v6.0.6 (https://startbootstrap.com/theme/coming-soon)
* Copyright 2013-2022 Start Bootstrap
* Licensed under MIT (https://github.com/StartBootstrap/startbootstrap-coming-soon/blob/master/LICENSE)
*/

class Message {
    constructor(opcode, argument) {
        this.opcode = opcode;
        this.argument = argument;
    }
}

function twit() {
    let twitMsg = new Message("twit", null);
    let data = JSON.stringify(twitMsg);
    socket.send(data);
}

function init() {
    url = new URL(window.location.href);
    url.protocol = "ws";
    url.pathname = "/ws";

    var socket = new WebSocket(url.href);
    socket.onopen = () => {
        console.log("Successfully Connected");
        setInterval(function() {
            let heartbeat = new Message("heartbeat", null);
            let data = JSON.stringify(heartbeat);
            socket.send(data);
        }, 2000);
    };

    socket.onclose = event => {
        console.log("Socket Closed Connection: ", event);
    };

    socket.onerror = error => {
        console.log("Socket Error: ", error);
    };

    socket.onmessage = (event) => {
        // TODO: build out the dynamic content updater
        console.log(event.data);
    }
}
