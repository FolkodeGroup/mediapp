/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx,html}",
  ],
  darkMode: 'class', // Habilita modo oscuro basado en clase
  theme: {
    extend: {
      colors: {
        // Colores personalizados para el modo oscuro
        dark: {
          bg: '#1a1a1a',
          surface: '#2a2a2a',
          text: 'rgba(255, 255, 255, 0.87)',
          'text-secondary': 'rgba(255, 255, 255, 0.6)',
        }
      }
    },
  },
  plugins: [],
}

