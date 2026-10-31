# Contributing

fx-repro is an early scaffold. Behavior counts once a human has confirmed it.

## Pull requests

- Keep a change small, on the layout in `docs/ARCHITECTURE.md`.
- Run `go test ./...` before opening a PR.
- Disclose AI assistance. Much of this code is written with AI coding agents; a human maintainer reviews and confirms behavior. Say what an agent wrote and what you ran or checked yourself.
- Fixtures in `tests/fixtures/` stay synthetic. Do not add real resources, escrowed assets, or files from a live server.
- Do not commit license keys, `sv_licenseKey`, database URLs, or other secrets. Exports must not grow a path that copies them out.

## Checks

CI builds with `CGO_ENABLED=0` and runs `go test ./...` on Ubuntu and Windows.
