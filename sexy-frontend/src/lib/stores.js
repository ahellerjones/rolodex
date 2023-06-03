import { writable } from "svelte/store";

export const auth = writable(true);
export const user = writable({
  username: "temp",
});
