# Architecture

## Language

Go. fx-repro is a local CLI for server admins on Windows and Linux. The work is directory walks, small text manifests, a dependency graph, and a resumable history file. The standard library covers that, and `CGO_ENABLED=0` produces one static binary. Cross-compiling is `GOOS` and `GOARCH` with no C toolchain, so a release is a single file an admin downloads and runs.

Rust would also ship a static binary. It is more toolchain than this tool needs.

## What v1 is

No FiveM resource bisector exists. Owners still cut a resources folder down by hand. Most reports are client crashes, which this process cannot observe: the server stays up while the player's game dies. v1 is a guided bisect. The tool proposes the next dependency-closed set, the human runs that trial, and the human records `crash`, `works`, or `unknown`. Automatic crash detection is a later, optional step.

The search shape — bisection that keeps required dependencies enabled, driven by a human verdict — is the same idea as [Qendolin/mod-bisect-tool](https://github.com/Qendolin/mod-bisect-tool) (Minecraft, MPL-2.0). That project is design inspiration only. This repo copies none of its code.

## Modules

These packages are the planned layout. Only `cmd/fx-repro` and `internal/cli` exist today.

| Package | Job |
| --- | --- |
| `cmd/fx-repro` | CLI entry: `scan`, `plan`, `record`, `status`, `export` |
| `internal/manifest` | Read `fxmanifest.lua` and `__resource.lua` as text. Never execute Lua. |
| `internal/graph` | Resources and their `dependency` / `dependencies` edges. |
| `internal/planner` | Delta-debugging / bisection. A trial set stays closed under required dependencies: enabling a resource enables what it requires, and a required resource stays while a dependent is still on. |
| `internal/history` | Local record of each trial (`crash`, `works`, `unknown`) so a run can resume. |
| `internal/export` | Minimal shareable repro: resource names and dependency edges. No commercial or escrowed file bytes. No `server.cfg` secrets, license keys, or connection strings. |

## Trial loop

`scan` fills the graph from a resources directory. `plan` asks the planner for the next set. The human applies that set and tries the crash. `record` stores the verdict. `plan` repeats on what is still unknown. `export` writes the smallest set still marked crashing, plus its dependency closure.

That is the whole program. No plugin host, no embedded server, no crash watcher in v1.
