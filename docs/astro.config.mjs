// @ts-check

import { defineConfig } from "astro/config";
import starlight from "@astrojs/starlight";
import starlightThemeSix from "@six-tech/starlight-theme-six";

// https://astro.build/config
export default defineConfig({
  site: "https://jspackr.dev",

  integrations: [
    starlight({
      title: "jspackr",
      description:
        "Documentation for jspackr, a fast and modern JavaScript and TypeScript bundler.",

      logo: {
        src: "./src/assets/logo.svg",
        alt: "jspackr",
      },

      favicon: "/favicon.svg",

      customCss: ["./src/styles/global.css"],

      plugins: [
        starlightThemeSix({
          footerText: "Created by Kaloka Radia Nanda.",
        }),
      ],

      lastUpdated: true,

      social: [
        {
          icon: "github",
          label: "GitHub",
          href: "https://github.com/kalokaradia/jspackr",
        },
      ],

      editLink: {
        baseUrl: "https://github.com/kalokaradia/jspackr/edit/main/docs/",
      },

      sidebar: [
        {
          label: "Get started",
          items: [
            {
              label: "Overview",
              slug: "index",
            },
            {
              label: "Installation",
              slug: "getting-started/installation",
            },
            {
              label: "Quick start",
              slug: "getting-started/quick-start",
            },
          ],
        },

        {
          label: "Guides",
          items: [
            {
              label: "Configuration",
              slug: "guides/configuration",
            },
            {
              label: "Bundling",
              slug: "guides/bundling",
            },
            {
              label: "Minification",
              slug: "guides/minification",
            },
            {
              label: "Source maps",
              slug: "guides/source-maps",
            },
            {
              label: "Watch mode",
              slug: "guides/watch-mode",
            },
            {
              label: "TypeScript and JSX",
              slug: "guides/typescript-jsx",
            },
          ],
        },

        {
          label: "CLI",
          items: [
            {
              label: "CLI reference",
              slug: "cli/reference",
            },
          ],
        },

        {
          label: "Reference",
          items: [
            {
              label: "Configuration reference",
              slug: "reference/configuration",
            },
            {
              label: "Supported features",
              slug: "reference/supported-features",
            },
            {
              label: "Troubleshooting",
              slug: "reference/troubleshooting",
            },
          ],
        },
      ],
    }),
  ],
});
