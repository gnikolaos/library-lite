/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["internal/templates/**/*.templ"],
    theme: {
        extend: {
            fontFamily: {
                acme: ["acme", "sans-serif"],
            },
            colors: {
                primary: "#d3cabb", // Light beige
                accent: "#a39171", // Light brown
            },
        },
    },
    plugins: [
        require("@tailwindcss/forms"),
        require("@tailwindcss/typography"),
    ],
};
