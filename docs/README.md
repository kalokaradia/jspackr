# jspackr documentation site

The official jspackr documentation is built with Astro and Starlight. Content lives in `src/content/docs`; navigation, repository links, and Starlight integration settings live in `astro.config.mjs`. Theme tokens and component styling are in `src/styles/global.css`.

## Development

Install dependencies from this directory, then run:

```sh
npm run dev
npm run build
npm run preview
```

The production build is written to `dist/`. Starlight provides the responsive documentation layout, sidebar, table of contents, search index, code highlighting, and color-mode control.

## Content structure

- `getting-started/`: install and make a first bundle
- `guides/`: explain build workflows and individual build options
- `cli/`: command-line reference
- `reference/`: configuration, current feature scope, and troubleshooting

Document behavior from the Go implementation and tests in the repository root. The Go CLI and configuration code are authoritative when older README text differs.
