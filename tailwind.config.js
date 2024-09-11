/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["internal/templates/**/*.templ"],
    theme: {
        extend: {
            fontFamily: {
                acme: ["acme", "sans-serif"],
            },
            colors: {
                primary: "hsl(var(--color-primary))", // Light beige
                accent: "hsl(var(--color-accent))", // Light brown
                cream: "hsl(var(--color-cream))", // Cream like
            },
        },
    },
    plugins: [
        require("@tailwindcss/forms"),
        require("@tailwindcss/typography"),
    ],
};
