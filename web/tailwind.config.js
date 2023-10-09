/** @type {import('tailwindcss').Config} */
export default {
    content: ['./index.html', './src/**/*.{js,ts,jsx,tsx}'],
    theme: {
        extend: {
            colors: {
                blue: {
                    700: '#122854',
                    800: '#09142a',
                },
                'golden-rod': '#efc96e',
            },
            fontFamily: {
                lora: ['Lora', 'serif'],
            },
        },
    },
    plugins: [],
}
