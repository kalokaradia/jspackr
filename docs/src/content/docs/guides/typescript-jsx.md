---
title: TypeScript and JSX
description: Use esbuild's TypeScript and JSX transforms through jspackr.
---

esbuild selects a loader from the entry and imported file extensions. jspackr passes files to esbuild as part of the bundle, so `.ts`, `.tsx`, `.jsx`, and JavaScript syntax can be transformed.

```ts title="src/index.ts"
type User = { name: string };

const user: User = { name: 'Ada' };
console.log(`Hello, ${user.name}`);
```

```sh
jspackr -i src/index.ts -o dist/app.js
```

TypeScript annotations and types are erased during esbuild's transformation; transpilation is not type checking. Type checking is optional and can be enabled separately:

```sh
npm install --save-dev typescript
jspackr -i src/index.ts -o dist/app.js --type-check
```

You can also enable checking in `jspackr.config.json`:

```json title="jspackr.config.json"
{
  "input": "src/index.ts",
  "output": "dist/app.js",
  "typeCheck": true
}
```

With `--type-check` or `"typeCheck": true`, jspackr runs the local `node_modules/.bin/tsc` (on Windows, `tsc.cmd`) as `tsc --noEmit` before esbuild. TypeScript must be installed in the project. TypeScript diagnostics are shown as emitted by `tsc`; a type-check failure gives jspackr a failing exit status and prevents bundling. The CLI flag can turn checking on when the config value is false or omitted; as with other boolean CLI options, omitting the flag does not turn off a config value set to true.

When present, `tsconfig.json` is read by TypeScript itself. Its file selection and compiler settings—including `strict`, module settings, `paths`, `baseUrl`, `jsx`, and `include`/`exclude`—determine the check. jspackr does not translate or override those settings. TypeScript and TSX checking, including JSX semantics, follows the project's TypeScript configuration; esbuild still performs the transformation and bundling.

JavaScript files are only checked when the TypeScript project enables JavaScript checking, usually with both `allowJs` and `checkJs`. jspackr does not turn on either option.

Without `tsconfig.json`, jspackr still invokes `tsc --noEmit` without adding file arguments or generating configuration. TypeScript's CLI then displays its help and exits unsuccessfully because no project or input files were selected. Add a `tsconfig.json` to define what should be checked.

For a selected TypeScript file, this is valid:

```ts
const name: string = "Kaloka";
const age: number = 15;
export {};
```

This assignment produces a TypeScript error and prevents the bundle from being built:

```ts
const age: number = "hello";
export {};
```

JSX/TSX transformation follows esbuild's JSX behavior. jspackr exposes no JSX runtime, factory, or import-source flags; configure your source and runtime accordingly. TypeScript and JSX syntax support does not imply framework-specific compilation.
