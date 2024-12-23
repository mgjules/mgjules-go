// vite.config.ts
import { defineConfig, presetUno, presetWebFonts } from "unocss";
import { theme } from "@unocss/preset-mini";

export default defineConfig({
  cli: {
    entry: [
      {
        patterns: [
          // "./templates/**/*.dhtml",
          // "./templates/**/*.scss",
          "./internal/templates/**/*.templ",
        ],
        outFile: "./internal/static/css/style.css",
      },
    ],
  },
  shortcuts: {
    pill: "w-fit border border-brand-secondary/50 py-1 px-2 text-xs text-brand-primary/90 shadow-md select-none print:border-brand-secondary print:text-brand-secondary print:shadow-none",
    "pill-link":
      "decoration-none hover:text-brand-accent hover:shadow-sm hover:border-brand-accent",
    btn: "inline-block px-2 py-1 transition-all decoration-none border border-brand-accent text-brand-primary shadow-md hover:text-brand-accent hover:shadow-sm print:shadow-none print:border-none print:p-0",
    "wrapped-dash": "py-2 border-dashed border-y border-brand-tertiary",
    link: "text-brand-accent hover:text-brand-primary no-underline",
  },
  theme: {
    colors: {
      brand: {
        foreground: "#252D38",
        background: "#191F28",
        primary: theme.colors?.gray[300],
        secondary: theme.colors?.gray[400],
        tertiary: theme.colors?.gray[500],
        accent: theme.colors?.cyan[500],
      },
    },
  },
  presets: [
    presetUno(),
    presetWebFonts({
      fonts: {
        sans: "Roboto",
        mono: ["Fira Code", "Fira Mono:400,700"],
      },
    }),
  ],
});
