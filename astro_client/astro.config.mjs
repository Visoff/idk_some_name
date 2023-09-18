import { defineConfig } from 'astro/config';

import svelte from "@astrojs/svelte";
import tailwind from "@astrojs/tailwind";
import node from "@astrojs/node"

// https://astro.build/config
export default defineConfig({
  output:"hybrid",
  adapter:node({
    mode:"standalone"
  }),
  server:{
    host:"0.0.0.0",
    port:3000
  },
  integrations: [svelte(), tailwind()]
});