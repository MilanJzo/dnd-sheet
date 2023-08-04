// place files you want to import through the `$lib` alias in this folder.

import { writable } from "svelte/store";

export const theme = writable("LIGHT");
export const secondary = writable("#6c6c6c");
export const primary = writable("#fff3cf");
export const text = writable("#000000");