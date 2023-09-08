/**
 *
 * @param {string} event
 * @param {(event: any) => void} func
 */
export function addEventListener(event: string, func: (event: any) => void): void;
/**
 *
 * @param {MediaStream} stream
 * @param {string} room_id
 */
export function connect(stream: MediaStream, room_id: string): Promise<void>;
