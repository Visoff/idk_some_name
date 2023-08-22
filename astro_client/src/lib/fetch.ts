export default function (path:string, other:RequestInit | undefined) {
    const url = new URL(path, import.meta.env.PUBLIC_SERVER_HOST)
    return fetch(url, other)
}