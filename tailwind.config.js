/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["internal/templates/**/*.templ"],
    theme: {
        extend: {
            fontFamily: {
                acme: ["acme", "sans-serif"],
            },
            colors: {
                primary: "var(--primary)", // Light beige
                accent: "var(--accent)", // Light brown
                cream: "var(--cream)", // Cream like
            },
        },
    },
    plugins: [
        require("@tailwindcss/forms"),
        require("@tailwindcss/typography"),
    ],
};
