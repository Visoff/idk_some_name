import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';
import mksert from 'vite-plugin-mkcert'

export default defineConfig({
	plugins: [sveltekit(), mksert()],
	server:{ https:true }
});
