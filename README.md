# fx-repro

Reduce a crashing FiveM (Cfx.re FXServer) server to a minimal, shareable reproduction, resource by resource.

**Status: early scaffold, not usable yet.** The CLI builds and prints help. `scan`, `plan`, `record`, `status`, and `export` print `not implemented yet`.

## The problem

A server crashes, or clients crash on connect, and the owner does not know which resource did it. The usual fix is manual: disable resources one by one, or disable half, restart, and repeat. That is the advice on the Cfx.re forum in [An MLO is causing a crash and idk which one](https://forum.cfx.re/t/an-mlo-is-causing-a-crash-and-idk-which-one-how-can-i-figure-this-out/5157230).

No FiveM resource bisector exists. The same hand narrowing shows up when people isolate client crashes in citizenfx/rfc discussions [#380](https://github.com/citizenfx/rfc/discussions/380) and [#162](https://github.com/citizenfx/rfc/discussions/162).

Most of the real demand is **client** crashes. The server process often stays up; the player's game dies, and only the human can say whether that trial crashed. v1 is an interactive, dependency-aware guided bisect: the tool proposes a reduced set that still includes required dependencies, and the human reports each trial as crash, works, or unknown. Automatic crash detection is a later, optional step. The tool never claims to detect client crashes itself.

## Goal

Read the server's resources and their manifests (`fxmanifest.lua` / `__resource.lua`), including dependencies. Propose reduced configurations that keep required dependencies. Let the operator record each trial, keep a history that can resume, and converge on a minimal repro they can share.

## Workflow

1. **Scan.** Read resource folders and manifests into a dependency graph.
2. **Plan.** Propose the next reduced set. The set stays closed under required dependencies.
3. **Trial.** The human applies that set and tries to reproduce the crash.
4. **Record.** The human marks the trial crash, works, or unknown.
5. **Reduce.** Repeat plan and record until the crashing set is minimal.
6. **Export.** Write that set as a shareable repro.

## Non-goals

- No automatic client-crash detection. The human confirms whether a crash is real.
- No redistribution of commercial or escrowed assets. An export names resources and dependency edges; it does not pack those files.
- No export of secrets: `server.cfg` keys, license keys, database URLs, or connection strings.

## Neighbors

These solve different jobs. They are complementary, not competitors.

- [fivem-server-doctor-mcp](https://github.com/mkcrw/fivem-server-doctor-mcp) — static checks of config, manifests, and logs.
- [fivem-mcp](https://github.com/eeharumt/fivem-mcp) — live RCON control of a running server.

## Design inspiration

[Qendolin/mod-bisect-tool](https://github.com/Qendolin/mod-bisect-tool) (Minecraft, MPL-2.0) is the prior art for a guided, dependency-aware mod bisect. It is design inspiration only. This repo copies none of its code.

## Build

Requires Go 1.22.

```sh
go build -o fx-repro ./cmd/fx-repro
./fx-repro
go test ./...
```

`CGO_ENABLED=0` is the intended release build: one static binary for Windows or Linux.

## AI assistance

Much of this code is written with AI coding agents; a human maintainer reviews and confirms behavior.
