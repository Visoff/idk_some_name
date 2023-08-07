import Clerk from "@clerk/clerk-js";
import { writable, type Writable } from "svelte/store";

export let CurrentChat:Writable<string/*uuid*/|false> = writable(false)
export let CurrentApp:Writable<"Chat"|"Video"> = writable("Chat")

export let clerk:Writable<Clerk> = writable(new Clerk("pk_test_Y2hhbXBpb24ta2luZ2Zpc2gtMTguY2xlcmsuYWNjb3VudHMuZGV2JA"))
export let clerk_ready:Writable<boolean> = writable(false)