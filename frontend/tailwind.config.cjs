const defaultTheme = require('tailwindcss/defaultTheme');

/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	theme: {
		extend: {
			fontFamily: {
				sans: ['"Work Sans"', ...defaultTheme.fontFamily.sans]
			},
			screens: {
				'3xl': '1792px',
				'4xl': '2048px'
			}
		}
	},
	plugins: [require('@tailwindcss/line-clamp')]
};
