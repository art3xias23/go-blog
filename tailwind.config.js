/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ["./components/*.html", "./components/*.templ", "./components/*.go",],
	theme: {
		screens: {
			'sm': '640px',

			'md': '768px',

			'lg': '1024px',

			'xl': '1280px',

			'2xl': '1536px',

			'3xl': '1920px',

			'4xl': '2560px'
		},
		extend: {
			gridTemplateRows: {
				// Define custom row sizes
				layout: '0.1fr 0.05fr 0.1fr 1fr',
			}
		},
	},
	safelist: [
		'bg-gray-800',
		'bg-opacity-70',
	],
	plugins: [],
}

