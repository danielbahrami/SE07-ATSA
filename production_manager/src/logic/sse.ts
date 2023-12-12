/* "use client"
const nofification_system = process.env.NOTIFICATION_SYSTEM;
const sse = new EventSource(`http://${nofification_system}/sse`);

export type OnMessage = (data: string) => void;

const callbacks: OnMessage[] = []

const addOnMessage = (callback: OnMessage) => {
    callbacks.push(callback);
}

sse.onopen = () => {
    console.log("Listening for events ...");
}

sse.onmessage = e => {
    console.log(e.data)
    callbacks.forEach(c => c(e.data))
}

export {sse, addOnMessage} */