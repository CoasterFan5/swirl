# Swirl playground

This is a small consumer project kept in the repository to exercise Swirl's build and editor/LSP support.

- `src/` contains app routes and project code.
- `.swirl/` contains generated project-local declarations and TypeScript config.
- `swirl-config.json` and `tsconfig.json` are project-level configuration.
- `go.mod` makes the app's Go code its own module; the repository-level `go.work` includes it during framework development.

From the repository root, run the development server with:

```sh
go run ./cmd/swirl -project ./examples/playground
```
