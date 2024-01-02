/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./components/**/*.templ", "./pages/**/*.templ"],
  theme: {
    fontFamily: {
      sans: ['Cormorant Garamond', 'sans-serif'],
    },
    extend: {
      colors: {
        "gold": {
          "50": "#fffef7",
          "100": "#fffdf0",
          "200": "#fffbe0",
          "300": "#fff9cf",
          "400": "#fff5ad",
          "500": "#fff189",
          "600": "#ffe666",
          "700": "#ffd03f",
          "800": "#ffbb19",
          "900": "#e6a600"
        }
      },
      backgroundImage: {
        "radial-gradient-center": "radial-gradient(circle at 50% 50%, var(--tw-gradient-stops))",
      },
    },
  },
  plugins: [],
}

