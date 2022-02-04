const defaultTheme = require('tailwindcss/defaultTheme')

function color(variableName) {
  return () => `var(${variableName})`
}

module.exports = {
  mode: 'jit',
  content: [
    "./components/**/*.{js,vue,ts}",
    "./layouts/**/*.vue",
    "./pages/**/*.vue",
    "./plugins/**/*.{js,ts}",
    // "./nuxt.config.{js,ts}",
  ],
  darkMode: 'class', // or 'media' or 'class'
  theme: {
    extend: {
      fontFamily: {
        sans: ['Poppins',...defaultTheme.fontFamily.sans],
      },
    },
  },
  variants: {
    extend: {
      fill: ['hover', 'focus'],
    },
  },
  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/aspect-ratio'),
    require('@tailwindcss/typography'),
    require('@tailwindcss/line-clamp'),
  ],
}
