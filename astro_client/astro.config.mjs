import { defineConfig } from 'astro/config';
import tailwind from "@astrojs/tailwind";
import svelte from "@astrojs/svelte";
import node from "@astrojs/node"

// https://astro.build/config
export default defineConfig({
  integrations: [tailwind(), svelte()],
  experimental:{
    viewTransitions:true
  },
  adapter:node({
    mode:"standalone"
  }),
  output:"hybrid",
  server:{
    host:"0.0.0.0"
  }
});