/**
 *
 * @param {"streamchange"} event
 * @param {(event: any) => void} func
 */
export function addEventListener(event: "streamchange", func: (event: any) => void): void;
/**
 *
 * @param {{url:string, room_id:string, stream:MediaStream}}
 */
export function connect({ url, room_id, stream }: {
    url: string;
    room_id: string;
    stream: MediaStream;
}): void;
