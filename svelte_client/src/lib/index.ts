import { writable, type Writable } from "svelte/store";

export let CurrentChat:Writable<string/*uuid*/|false> = writable(false)
export let CurrentApp:Writable<"Chat"|"Video"> = writable("Chat")