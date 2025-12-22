<p align="center">
<img src="https://github.com/andygeiss/cloud-native-app/blob/main/logo.png?raw=true" />
</p>

# Cloud Native App

[![License](https://img.shields.io/github/license/andygeiss/cloud-native-app)](https://github.com/andygeiss/cloud-native-app/blob/master/LICENSE)
[![Releases](https://img.shields.io/github/v/release/andygeiss/cloud-native-app)](https://github.com/andygeiss/cloud-native-app/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/andygeiss/cloud-native-app)](https://goreportcard.com/report/github.com/andygeiss/cloud-native-app)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/30de4509fb6d4b388e2dc52157afb0fd)](https://app.codacy.com/gh/andygeiss/cloud-native-app/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)
[![Codacy Badge](https://app.codacy.com/project/badge/Coverage/30de4509fb6d4b388e2dc52157afb0fd)](https://app.codacy.com/gh/andygeiss/cloud-native-app/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_coverage)

**Cloud Native App** is a small Go CLI that scaffolds a new Go module (repository) from embedded templates.
It creates a ready-to-run HTTP service skeleton (UI pages, static assets, hexagonal-ish layout), initializes `go.mod`, and bootstraps a fresh Git repo.

## Features

- **Project scaffolding** — creates a new folder named after the module basename and populates it from templates
- **HTTP service skeleton** — `cmd/service` with embedded assets and `/ui` routes
- **OIDC-friendly UI** — templates link to `/auth/login` and `/auth/logout` (provided by `cloud-native-utils/security`)
- **Opt-in profiling** — `/debug/pprof/` route when `ENABLE_PPROF=true`
- **Dev automation** — generated `.justfile` includes build/run/test tasks (service build uses PGO)
- **CI** — GitHub Actions workflow runs `just test` and uploads coverage to Codacy

## Technologies Used

- [Go](https://golang.org/) — language/runtime
- [cloud-native-utils](https://github.com/andygeiss/cloud-native-utils) — templating, security/mux helpers, logging, messaging abstractions
- [Just](https://github.com/casey/just) — task runner (optional)
- (Generated module) Keycloak — for OIDC when you wire it up via env
- (Generated module) Kafka — only referenced via `.env` placeholders (dispatcher is created; app code doesn’t publish/consume by default)

## What This Repo Contains

This repository contains:

- A CLI entrypoint at `cmd/app/main.go` that embeds templates from `cmd/app/templates`
- A template rendering adapter at `internal/app/adapters/outbound/template.go`
- A module generator service at `internal/app/core/services/module.go`
- Unit tests that validate template rendering and module generation

The generator writes a new module to a directory named `filepath.Base(<module-arg>)` and then runs:

- `go mod init <module-arg>`
- `go mod tidy`
- `git init && git add . && git commit -m "feat: init repo"`

If Git isn’t configured (user name/email), the final commit step can fail.

## Generated Module Layout (High-Level)

The scaffolded module includes (among others):

```
.
├── .env
├── .gitignore
├── .justfile
├── README.md
├── cmd/
│   └── service/
│       ├── main.go
│       └── assets/
│           ├── static/
│           │   ├── base.css
│           │   ├── theme.css
│           │   └── styles.css
│           └── templates/
│               ├── index.tmpl
│               └── login.tmpl
└── internal/
		└── app/
				├── adapters/
				│   └── ingres/
				│       ├── router.go
				│       ├── middleware.go
				│       └── ui/
				│           ├── index.go
				│           ├── login.go
				│           └── view.go
				├── config/
				│   └── config.go
				└── errors.go
```

Note: the generated HTML templates reference `/static/htmx.min.js`; the generator does not currently write that file into `assets/static`. Add it yourself (or remove the `<script>` tag) if you want HTMX.

## Build / Install This CLI

### Prerequisites

- Go (this repo targets Go 1.25.4 via `go.mod`)
- `git` on your PATH
- `just` (optional, but recommended)
- Ensure `$HOME/bin` is on your `PATH` if you want `just install` to work as intended

### Install

Build and install `cloud-native-app` into `$HOME/bin`:

```bash
just install
```

Or build locally:

```bash
just build
./bin/cloud-native-app <module>
```

## Usage

Run the generator from the directory where you want the new module folder to be created:

```bash
cloud-native-app github.com/your-org/demo
```

This creates a new folder `demo/` (basename of the module) and writes the scaffold into it.

### Run The Generated Service

```bash
cd demo
just run
```

The UI routes are mounted under `/ui/`.

## Local Keycloak / Kafka (Optional)

The generated `.env` contains OIDC and Kafka placeholders. If you want to exercise the login flow, you’ll need an OIDC provider (e.g., Keycloak).

Example with Podman:

```bash
podman run --restart=always --name keycloak -d -p 127.0.0.1:8180:8080 \
	-e KC_BOOTSTRAP_ADMIN_USERNAME=admin \
	-e KC_BOOTSTRAP_ADMIN_PASSWORD=admin \
	quay.io/keycloak/keycloak:26.3.4 start-dev

podman run --restart=always --name kafka -d -p 127.0.0.1:9092:9092 apache/kafka:latest
```

## Tests

Run this repo’s unit tests:

```bash
just test
```

## Repository Structure

```
cmd/app/                      # CLI entrypoint
cmd/app/templates/            # embedded templates for generated module
internal/app/adapters/outbound/  # template rendering adapter
internal/app/core/services/   # module generator service + tests
bin/                          # compiled binaries (local)
```

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.
