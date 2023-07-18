/** @type {import('tailwindcss').Config} */
export default {
    content: ['./src/**/*.{html,js,svelte,ts}'],
    theme: {
      extend: {
        colors: {
          'ease-grey': '#F6F6F8'
        },
        fontFamily: {
          s: "sf-regular",
          m: "sf-medium",
          l: "sf-semibold",
          xl: "sf-bold",
        },
    },
    plugins: [],
  }
}

