// "use client"
// const sse = new EventSource("http://127.0.0.1:5000/sse");
//
// export type OnMessage = (data: string) => void;
//
// const callbacks: OnMessage[] = []
//
// const addOnMessage = (callback: OnMessage) => {
//     callbacks.push(callback);
// }
//
// sse.onopen = () => {
//     console.log("Listening for events ...");
// }
//
// sse.onmessage = e => {
//     console.log(e.data)
//     callbacks.forEach(c => c(e.data))
// }
//
// export {sse, addOnMessage}