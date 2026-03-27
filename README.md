# Resume Website — A Full Tutorial

This project is a resume website built to teach two technologies at once:

- **Go** — a compiled, statically-typed language made by Google. We use it to build a REST API that serves resume data as JSON over HTTP.
- **React** — a JavaScript library for building user interfaces. We use it to fetch that JSON and render it as a webpage.

The two halves talk to each other over HTTP, just like any real-world web application. The Go server is the backend; the React app is the frontend.

If you are brand new to both, start at the top and read every section. Each one builds on the last.

---

## Table of Contents

1. [How the Web Works — The Basics](#how-the-web-works--the-basics)
2. [Project Structure](#project-structure)
3. [How to Run](#how-to-run)
4. [Devcontainer](#devcontainer)
5. [Backend — Go REST API](#backend--go-rest-api)
   - [What is a REST API?](#what-is-a-rest-api)
   - [go.mod — The Module File](#gomod--the-module-file)
   - [main.go — The Entry Point](#maingo--the-entry-point)
   - [models/resume.go — Defining Data Shapes](#modelsresumego--defining-data-shapes)
   - [data/resume.go — The Actual Content](#dataresumego--the-actual-content)
   - [handlers/resume.go — Responding to Requests](#handlersresumego--responding-to-requests)
   - [api.http — Testing the API](#apihttp--testing-the-api)
6. [Frontend — React](#frontend--react)
   - [What is React?](#what-is-react)
   - [package.json — The Project Manifest](#packagejson--the-project-manifest)
   - [vite.config.js — The Build Tool](#viteconfigjs--the-build-tool)
   - [index.html — The Shell](#indexhtml--the-shell)
   - [src/main.jsx — Booting React](#srcmainjsx--booting-react)
   - [src/App.jsx — The Root Component](#srcappjsx--the-root-component)
   - [src/hooks/useResume.js — Shared Fetch Logic](#srchooksuseresumeJS--shared-fetch-logic)
   - [src/components/Bio.jsx](#srccomponentsBiojsx)
   - [src/components/Experience.jsx](#srccomponentsExperiencejsx)
   - [src/components/Education.jsx](#srccomponentsEducationjsx)
   - [src/components/Skills.jsx](#srccomponentsSkillsjsx)
   - [CSS Modules — Scoped Styles](#css-modules--scoped-styles)
7. [How the Two Halves Connect](#how-the-two-halves-connect)

---

## How the Web Works — The Basics

Before touching any code, here is the mental model you need.

When you visit a website, your browser sends an **HTTP request** to a server. The server sends back an **HTTP response**. That's it. Every web application, no matter how complex, is built on top of this request/response cycle.

An HTTP request has:
- A **method** — the most common are `GET` (fetch something) and `POST` (send something)
- A **URL** — the address being requested, e.g. `/api/resume/bio`
- **Headers** — metadata about the request (who's asking, what format they accept, etc.)
- An optional **body** — data sent along with the request (common in POST)

An HTTP response has:
- A **status code** — a number that tells the client what happened. `200 OK` means success. `404 Not Found` means the URL doesn't exist. `500 Internal Server Error` means something broke on the server.
- **Headers** — metadata about the response (what format the data is in, caching rules, etc.)
- A **body** — the actual content being returned (HTML, JSON, an image, etc.)

**JSON** (JavaScript Object Notation) is a text format for sending structured data. It looks like this:

```json
{
  "name": "Your Name",
  "title": "Software Engineer",
  "skills": ["Go", "React", "SQL"]
}
```

JSON is language-agnostic — Go can produce it, JavaScript can consume it, and any other language can read it too. This is why it became the universal language of web APIs.

---

## Project Structure

```
Resume-Website/
├── .devcontainer/
│   ├── Dockerfile          # Blueprint for the development environment container
│   └── devcontainer.json   # VS Code instructions for using that container
├── backend/
│   ├── main.go             # Go program entry point — starts the HTTP server
│   ├── go.mod              # Declares this as a Go module (like package.json for Go)
│   ├── api.http            # Test file: fire HTTP requests from VS Code
│   ├── data/
│   │   └── resume.go       # Your actual resume content (edit this)
│   ├── handlers/
│   │   └── resume.go       # Functions that handle each API request
│   └── models/
│       └── resume.go       # Data type definitions (the shape of resume data)
└── frontend/
    ├── package.json        # Declares this as a Node.js project + lists dependencies
    ├── vite.config.js      # Configuration for the dev server and build tool
    ├── index.html          # The one HTML page React lives inside
    └── src/
        ├── main.jsx        # React boot file — connects React to index.html
        ├── App.jsx         # The root component — assembles the full page
        ├── index.css       # Global styles (resets, body font)
        ├── App.module.css  # Layout styles for the page wrapper
        ├── hooks/
        │   └── useResume.js        # Reusable data-fetching logic
        └── components/
            ├── Bio.jsx             # Name, title, contact info, links
            ├── Bio.module.css      # Styles for the Bio section
            ├── Experience.jsx      # Work history
            ├── Education.jsx       # Education history
            ├── Skills.jsx          # Skills grouped by category
            ├── Section.module.css  # Shared styles used by Experience/Education/Skills
            └── Skills.module.css   # Tag-pill styles specific to Skills
```

---

## How to Run

### Option A — Devcontainer (recommended for a consistent environment)

Open this project in VS Code. Press `Cmd+Shift+P` and run **Dev Containers: Rebuild and Reopen in Container**. Once VS Code reopens inside the container, open two terminals:

```bash
# Terminal 1 — Go API server (runs on port 8080)
cd /workspaces/Resume-Website/backend
go run .

# Terminal 2 — React dev server (runs on port 5173)
cd /workspaces/Resume-Website/frontend
npm run dev
```

Open `http://localhost:5173` in your browser.

### Option B — Run directly on your machine

Requires Go 1.21+ and Node.js 20+ installed.

```bash
# Terminal 1
cd backend
go run .

# Terminal 2
cd frontend
npm install   # only needed the first time
npm run dev
```

---

## Devcontainer

A **devcontainer** solves the "it works on my machine" problem. It is a Docker container that holds your entire development environment — the right version of Go, Node.js, and any other tools — packaged up so every developer gets an identical setup.

Docker works by building **images** from a set of instructions called a Dockerfile. An image is like a snapshot of a filesystem. When you start a container from that image, you get an isolated environment that runs the same way everywhere.

### `.devcontainer/Dockerfile`

A Dockerfile is read top to bottom. Each instruction (`FROM`, `RUN`, `ENV`, etc.) adds a layer to the image.

```dockerfile
FROM registry.access.redhat.com/ubi9/ubi:latest
```

`FROM` declares the **base image** — the starting point everything else builds on. `registry.access.redhat.com/ubi9/ubi` is **Red Hat Universal Base Image 9** (UBI9), a freely redistributable, minimal version of Red Hat Enterprise Linux 9. It uses `dnf` as its package manager, the same tool found on RHEL, CentOS, and Fedora.

```dockerfile
RUN dnf install -y --allowerasing \
    git curl wget tar gcc make \
    && dnf clean all
```

`RUN` executes a shell command during the image build. `dnf install -y` installs packages without interactive prompts. The `--allowerasing` flag is needed because UBI9 ships a stripped-down `curl-minimal` that conflicts with the full `curl` package — this flag tells dnf it's allowed to replace it. `dnf clean all` removes the package metadata cache to keep the image smaller.

```dockerfile
ARG GO_VERSION=1.23.4
RUN ARCH=$(uname -m | sed 's/x86_64/amd64/;s/aarch64/arm64/') \
    && curl -fsSL https://go.dev/dl/go${GO_VERSION}.linux-${ARCH}.tar.gz \
    | tar -C /usr/local -xz
```

`ARG` declares a **build-time variable** (can be overridden when building with `--build-arg`). `uname -m` prints the CPU architecture (`x86_64` on Intel/AMD chips, `aarch64` on Apple Silicon). The `sed` command substitutes those names into the ones Go uses (`amd64`, `arm64`). The tarball is piped directly into `tar` so no intermediate file is saved to disk.

```dockerfile
ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH="/home/developer/go"
ENV PATH="${GOPATH}/bin:${PATH}"
```

`ENV` sets **environment variables** that persist at runtime — when you open a terminal in the container, these variables are already set. `PATH` is the list of directories the shell searches when you type a command. Adding `/usr/local/go/bin` to `PATH` is what makes `go` available as a command. `GOPATH` is Go's workspace directory for downloaded modules and compiled tools.

```dockerfile
ARG NODE_VERSION=20.18.1
RUN ARCH=$(uname -m | sed 's/x86_64/x64/;s/aarch64/arm64/') \
    && curl -fsSL https://nodejs.org/dist/v${NODE_VERSION}/node-v${NODE_VERSION}-linux-${ARCH}.tar.gz \
    | tar -C /usr/local -xz --strip-components=1
```

Node.js is installed the same way as Go — download a prebuilt binary tarball and extract it. `--strip-components=1` removes the top-level folder inside the archive (e.g. `node-v20.18.1-linux-arm64/`) so the contents land directly in `/usr/local/bin`, `/usr/local/lib`, etc.

```dockerfile
RUN groupadd --gid 1000 developer \
    && useradd --uid 1000 --gid 1000 -m developer \
    && chown -R developer:developer /workspace 2>/dev/null || true

USER developer
```

By default, Docker runs everything as root (the all-powerful system user). Running development tools as root is risky and bad practice. These lines create an unprivileged `developer` user, give them ownership of the workspace, and switch to that user. Every command after `USER developer` — including your terminal sessions in VS Code — runs as this user.

### `.devcontainer/devcontainer.json`

This file tells VS Code how to use the container.

```json
"workspaceFolder": "/workspaces/Resume-Website"
```

The path inside the container where VS Code mounts your project files. Your local files are bind-mounted here — edits you make in VS Code on your Mac are immediately visible inside the container.

```json
"forwardPorts": [8080, 5173]
```

The container is an isolated network. Without port forwarding, `localhost:8080` on your Mac would not reach the Go server running inside the container. This setting punches holes in that isolation for exactly the ports we need.

```json
"extensions": ["golang.go", "dbaeumer.vscode-eslint", ...]
```

VS Code extensions installed automatically inside the container. These are separate from your local extensions — the container gets its own set.

```json
"postCreateCommand": "go -C /workspaces/Resume-Website/backend mod download && npm --prefix /workspaces/Resume-Website/frontend install"
```

A shell command that runs once after the container is first created. It downloads Go module dependencies and installs Node.js packages so you don't have to do it manually. `go -C <dir>` runs the Go command from the given directory; `npm --prefix <dir>` does the same for npm.

---

## Backend — Go REST API

### What is a REST API?

**REST** (Representational State Transfer) is a convention for designing web APIs. The core idea is simple: different URLs represent different **resources**, and you use HTTP methods to say what you want to do with them.

In this project, the resources are parts of a resume:

| HTTP Method | URL | What it does |
|---|---|---|
| `GET` | `/api/health` | Check if the server is running |
| `GET` | `/api/resume` | Get the entire resume |
| `GET` | `/api/resume/bio` | Get only the bio (name, title, etc.) |
| `GET` | `/api/resume/experience` | Get the list of jobs |
| `GET` | `/api/resume/education` | Get the list of schools |
| `GET` | `/api/resume/skills` | Get the list of skill groups |

Every response is JSON. The frontend fetches these URLs and renders the data.

### `go.mod` — The Module File

```
module github.com/dsluss/resume-website/backend

go 1.23
```

Every Go project begins with a `go.mod` file. It does two things:

1. **Declares the module path** — a globally unique name for your code. By convention it's a URL-like path (usually your GitHub repo). Other packages within this project import each other using this path as a prefix. For example, `main.go` imports `github.com/dsluss/resume-website/backend/handlers`.

2. **Declares the minimum Go version** — any Go 1.23.x or later will work.

This project has no external dependencies, so there are no `require` lines. The entire backend is built on Go's standard library.

### `main.go` — The Entry Point

Every executable Go program starts here. Go requires a file with `package main` and a function called `func main()`. When you run `go run .`, Go compiles everything and calls `main()`.

#### Imports

```go
import (
    "fmt"
    "log"
    "net/http"
    "github.com/dsluss/resume-website/backend/handlers"
)
```

Go's imports are explicit — you must list every package you use, or the code will not compile. Standard library packages use short names (`net/http`, `fmt`). Your own packages use the full module path.

- `fmt` — formatted output (`Printf`, `Sprintf`)
- `log` — logging with timestamps, and `log.Fatal` which logs and then exits the program
- `net/http` — Go's built-in HTTP client and server
- `handlers` — our own package defined in `backend/handlers/`

#### The ServeMux — What is a Multiplexer?

```go
mux := http.NewServeMux()
```

A **ServeMux** (short for **server multiplexer**) is a router. The word "multiplex" comes from telecommunications and means "combining multiple signals into one channel." In HTTP, a multiplexer receives all incoming requests on a single port and routes each one to the right function based on the URL.

Think of it like a phone switchboard operator. All calls come in on one line. The operator looks at who the caller is asking for and connects them to the right extension.

Without a mux, every HTTP request to your server would go to the same function. The mux lets you say:
- "If the URL is `/api/health`, run *this* function"
- "If the URL is `/api/resume/bio`, run *that* function"

```go
mux.HandleFunc("GET /api/health", handlers.Health)
mux.HandleFunc("GET /api/resume", handlers.GetResume)
mux.HandleFunc("GET /api/resume/bio", handlers.GetBio)
mux.HandleFunc("GET /api/resume/experience", handlers.GetExperience)
mux.HandleFunc("GET /api/resume/education", handlers.GetEducation)
mux.HandleFunc("GET /api/resume/skills", handlers.GetSkills)
```

Each line **registers a route** — a pairing of a URL pattern to a handler function. The pattern `"GET /api/health"` means "only match HTTP GET requests to the path `/api/health`." Specifying the method in the pattern is a Go 1.22 feature; in older Go code you would check `r.Method` inside the handler manually.

#### The Health Endpoint — Why Does It Exist?

The `/api/health` route is not part of the resume data. It is a **health check** — a standard practice in production systems. Monitoring tools, load balancers, and container orchestrators (like Kubernetes) hit `/health` periodically. If it returns 200, the service is considered alive. If it doesn't respond, the system knows something is wrong and can restart the container or route traffic elsewhere.

Even in a simple project like this, having a health check is a good habit to build.

#### Middleware — Wrapping the Router

```go
handler := corsMiddleware(mux)
log.Fatal(http.ListenAndServe(addr, handler))
```

Notice that `mux` is not passed directly to `ListenAndServe`. Instead it is wrapped with `corsMiddleware`. This is the **middleware pattern**.

**Middleware** is a function that sits between the server and your handlers. It receives every request first, can inspect or modify it, and then decides whether to pass it on to the next handler. It can also do things after the handler runs. Common uses: logging, authentication, rate limiting, and CORS.

The middleware pattern in Go always looks like this:
```go
func someMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // do something before the handler runs
        next.ServeHTTP(w, r)  // call the next handler
        // do something after (optional)
    })
}
```

You can chain multiple middleware together: `authMiddleware(loggingMiddleware(mux))`.

#### CORS — Why the Frontend Needs Permission

```go
func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
        w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusNoContent)
            return
        }

        next.ServeHTTP(w, r)
    })
}
```

**CORS** (Cross-Origin Resource Sharing) is a browser security feature. Browsers enforce a **same-origin policy**: JavaScript running on `localhost:5173` is not normally allowed to make HTTP requests to `localhost:8080`, because they are on different ports, which makes them different "origins."

The browser enforces this by first sending a **preflight request** — an HTTP `OPTIONS` request — to ask the server: "I'm going to make a request from origin X. Is that okay?" The server must respond with the right headers to grant permission, or the browser will block the actual request before it is even sent.

Our middleware adds these permission headers to every response:
- `Access-Control-Allow-Origin` — which origins are allowed to make requests
- `Access-Control-Allow-Methods` — which HTTP methods are allowed
- `Access-Control-Allow-Headers` — which request headers are allowed

When the method is `OPTIONS` (preflight), we return `204 No Content` immediately without calling `next.ServeHTTP` — there's no data to return, just the CORS headers the browser needs.

**Important**: CORS is enforced by the browser, not the server. If you call the API directly (e.g. with `curl` or the REST Client extension), CORS headers are irrelevant — they are only checked by browsers.

#### Starting the Server

```go
addr := ":8080"
fmt.Printf("Resume API listening on http://localhost%s\n", addr)
log.Fatal(http.ListenAndServe(addr, handler))
```

`http.ListenAndServe` starts the HTTP server. The first argument is the address to listen on — `:8080` means "all network interfaces, port 8080." It blocks forever, processing requests in a loop. If it ever returns (e.g. the port is already in use), it returns an error. `log.Fatal` prints that error and exits the program with a non-zero status code.

### `models/resume.go` — Defining Data Shapes

Before you can handle HTTP requests or encode JSON, you need to define what shape your data takes. In Go this is done with **structs**.

```go
type Resume struct {
    Bio        Bio          `json:"bio"`
    Experience []Experience `json:"experience"`
    Education  []Education  `json:"education"`
    Skills     []SkillGroup `json:"skills"`
}
```

A **struct** is a collection of named, typed fields grouped together. If you come from object-oriented languages, think of it as a class with only data (no methods). The fields here map directly to the sections of a resume.

`[]Experience` is a **slice** — Go's dynamically-sized list. A slice of type `T` is written `[]T`. You can add items, remove items, and iterate over them. In JSON, a slice becomes an array `[...]`.

#### Struct Tags — Controlling JSON Field Names

```go
Bio Bio `json:"bio"`
```

The backtick-enclosed string after the type is a **struct tag** — metadata attached to a field. The `encoding/json` package reads the `json:"..."` tag to determine the JSON key name.

Without a tag, Go would use the field name directly: `Bio` → `"Bio"` (capital B). With the tag, it becomes `"bio"` (lowercase), which follows standard JSON conventions. This also lets you use short, descriptive Go field names while having different JSON names if needed.

```go
type Experience struct {
    Company    string   `json:"company"`
    Role       string   `json:"role"`
    StartDate  string   `json:"start_date"`  // Go: StartDate, JSON: start_date
    EndDate    string   `json:"end_date"`
    Location   string   `json:"location"`
    Highlights []string `json:"highlights"`
}
```

Notice `StartDate` in Go becomes `start_date` in JSON. Go convention is CamelCase for exported names; JSON convention is snake_case. The struct tag bridges the two.

### `data/resume.go` — The Actual Content

```go
func SeedResume() models.Resume {
    return models.Resume{
        Bio: models.Bio{
            Name:  "Your Name",
            Title: "Software Engineer",
            ...
        },
    }
}
```

This function constructs a `models.Resume` value using a **struct literal** — you name the struct type and fill in its fields with `{FieldName: value}`. The `models.` prefix is required because `Bio` is defined in the `models` package, not the `data` package.

This is where you put your real resume content. In a production application, this data would typically come from a database. For this learning project, hardcoding it here keeps things simple and removes the need to set up a database.

### `handlers/resume.go` — Responding to Requests

A **handler** is a function that runs in response to an HTTP request. Each API endpoint has exactly one handler. The handler's job is to:
1. Read whatever it needs from the request
2. Do some work (in our case, fetch resume data)
3. Write the response

#### Package-level State and `init()`

```go
var resume models.Resume

func init() {
    resume = data.SeedResume()
}
```

`var` at the package level declares a variable that exists for the entire lifetime of the program — not just for one function call.

`init()` is a special function Go calls automatically before `main()` runs, once per package. Here it loads the resume data once at startup. This means every handler can just read `resume` directly without reloading the data on every request.

You can think of this like a cache: load the data once, reuse it forever. For data that never changes (like a hardcoded resume), this is the right approach.

#### The `writeJSON` Helper

```go
func writeJSON(w http.ResponseWriter, status int, v any) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(v)
}
```

Every handler needs to do the same three things before sending a response:

1. **Set the `Content-Type` header** — tells the client what format the body is in. `application/json` signals that the body is JSON. Without this, the browser might not know how to interpret the response.

2. **Write the status code** — `w.WriteHeader(status)` sends the HTTP status line. This must be called after setting all headers and before writing the body. Once you write the status, headers are sent and cannot be changed.

3. **Encode and send the body** — `json.NewEncoder(w).Encode(v)` takes any Go value (`v any` — the `any` type accepts anything), converts it to JSON, and writes it to the response. It uses the `json:"..."` struct tags defined in `models/resume.go` to decide the field names.

Pulling these three steps into a helper means each handler can be a single line instead of three.

#### The Handlers

```go
func Health(w http.ResponseWriter, r *http.Request) {
    writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}
```

`http.StatusOK` is a constant equal to `200`. Using named constants instead of raw numbers makes the code more readable and less error-prone. Go's `net/http` package defines constants for all standard status codes: `http.StatusNotFound` (404), `http.StatusInternalServerError` (500), etc.

`map[string]string{"status": "ok"}` creates an inline key-value map. Go maps work like dictionaries or hash maps. `map[string]string` means both the keys and values are strings. This encodes to `{"status":"ok"}` in JSON.

```go
func GetBio(w http.ResponseWriter, r *http.Request) {
    writeJSON(w, http.StatusOK, resume.Bio)
}
```

The other handlers are identical in structure, each returning a different field of the `resume` variable. `resume.Bio` is the `Bio` struct; `resume.Experience` is the slice of `Experience` structs, and so on.

The `r *http.Request` parameter gives you access to everything about the incoming request — URL, headers, body, query parameters. None of our handlers use it yet because we are just serving static data. In a real API you might read a user ID from the URL to look up their specific resume.

### `api.http` — Testing the API

The REST Client VS Code extension reads `.http` files. You can click "Send Request" above any block to fire that HTTP request and see the response inline in VS Code. This lets you test your API without a browser or external tools like Postman.

---

## Frontend — React

### What is React?

React is a JavaScript library that lets you build a user interface by composing small, reusable pieces called **components**. A component is a function that returns a description of what to display. When the data behind a component changes, React efficiently updates only the parts of the page that need to change — you do not manually manipulate the HTML.

React uses **JSX** — a syntax extension that looks like HTML inside JavaScript:

```jsx
function Greeting() {
  return <h1>Hello, world</h1>
}
```

This is not real HTML. It gets compiled to plain JavaScript (`React.createElement('h1', null, 'Hello, world')`) before running in the browser. JSX is just a more readable way to describe UI structure.

### `package.json` — The Project Manifest

```json
{
  "name": "resume-frontend",
  "type": "module",
  "scripts": {
    "dev": "vite",
    "build": "vite build",
    "preview": "vite preview"
  },
  "dependencies": { "react": "^18.3.1", "react-dom": "^18.3.1" },
  "devDependencies": { "vite": "^5.4.1", "@vitejs/plugin-react": "^4.3.1" }
}
```

`package.json` is to Node.js what `go.mod` is to Go — it describes the project and its dependencies.

`"type": "module"` tells Node.js to use modern ES Module syntax (`import`/`export`) instead of the older CommonJS syntax (`require`/`module.exports`).

**`scripts`** are shortcuts run with `npm run <name>`:
- `npm run dev` — starts the development server with hot reload
- `npm run build` — compiles everything into static files in `dist/` ready to deploy
- `npm run preview` — serves the built `dist/` folder locally so you can test the production build

**`dependencies`** are packages that ship to the browser and run in production:
- `react` — the core React library (components, hooks, the virtual DOM)
- `react-dom` — the bridge between React and the actual browser DOM. React itself is platform-agnostic; `react-dom` is the part specific to web browsers.

**`devDependencies`** are tools used only during development and building:
- `vite` — the development server and build tool
- `@vitejs/plugin-react` — Vite plugin that adds JSX support and React fast refresh

### `vite.config.js` — The Build Tool

**Vite** (French for "fast") is the development server and build tool. During development it serves your files to the browser and watches for changes. When you save a file, the browser updates in under a second without a full page reload — this is called **Hot Module Replacement (HMR)**.

```js
plugins: [react()]
```

The React plugin enables two things:
- **JSX transformation** — compiles JSX syntax into regular JavaScript
- **React Fast Refresh** — a smarter version of HMR that updates components without losing their current state

```js
server: {
  proxy: {
    '/api': {
      target: 'http://localhost:8080',
      changeOrigin: true,
    },
  },
}
```

This is the **development proxy**, and it is what allows the frontend to talk to the backend without CORS issues in the browser.

Here is the problem it solves: your React app runs on `localhost:5173`. Your Go API runs on `localhost:8080`. When the React app tries to fetch `/api/resume/bio`, the browser sees this as a cross-origin request (different port = different origin) and blocks it, even with the CORS headers on the Go server.

The Vite proxy intercepts any request your React app makes to `/api/*` and forwards it to `http://localhost:8080` from the Vite server itself — not from the browser. The browser only ever sees `localhost:5173`, so there is no cross-origin issue at all.

This is development-only. In production you would configure a real web server (nginx, Caddy, etc.) to do the same forwarding.

### `index.html` — The Shell

```html
<body>
  <div id="root"></div>
  <script type="module" src="/src/main.jsx"></script>
</body>
```

This is the only HTML file in the entire frontend. React is a **Single Page Application (SPA)** — the browser loads this one page, and JavaScript takes over from there. React never causes a full page reload; it updates the DOM directly and manipulates what you see.

The `<div id="root">` is the **mount point** — an empty container where React will inject everything it renders. The `<script>` tag is the entry point that starts the whole React application.

### `src/main.jsx` — Booting React

```jsx
import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.jsx'

createRoot(document.getElementById('root')).render(
  <StrictMode>
    <App />
  </StrictMode>,
)
```

This file runs once and connects React to the HTML page. Reading it line by line:

1. `document.getElementById('root')` — finds that `<div id="root">` in `index.html`
2. `createRoot(...)` — tells React-DOM to take control of that div
3. `.render(<App />)` — renders the `App` component into it

**`StrictMode`** is a development-only wrapper that activates extra warnings and double-invokes certain functions to help you find bugs early. It produces no visible output and has zero effect in production builds.

`import './index.css'` loads the global stylesheet. Vite handles this — it injects the CSS into the page automatically.

### `src/App.jsx` — The Root Component

```jsx
export default function App() {
  return (
    <main className={styles.layout}>
      <Bio />
      <Experience />
      <Education />
      <Skills />
    </main>
  )
}
```

`App` is the **root component** — the top of the component tree. It does nothing except compose the page from four child components.

**`export default`** marks this as the main export of the file. When another file does `import App from './App.jsx'`, it gets this function.

**`className`** is used instead of HTML's `class` because `class` is a reserved keyword in JavaScript. This is one of the main differences between JSX and HTML.

**`<Bio />`** is a self-closing component tag. When React encounters this, it calls the `Bio` function, gets back JSX, and renders that in place of the tag. Components can be nested arbitrarily deep.

`{styles.layout}` is a JavaScript expression inside JSX — curly braces let you escape from JSX back into JavaScript to evaluate any expression. `styles` is an imported CSS Module object; `styles.layout` is the generated class name for `.layout` in `App.module.css`.

### `src/hooks/useResume.js` — Shared Fetch Logic

A **hook** is a special kind of function in React. Hooks let function components do things that previously required class components: managing state, running side effects, sharing logic between components.

All four section components (Bio, Experience, Education, Skills) need to do the same thing:
1. Fetch data from an API endpoint
2. Track whether the data is still loading
3. Track whether an error occurred
4. Return the data when it arrives

Rather than copy-paste this logic into every component, we extract it into a **custom hook**.

```js
export function useResumeSection(path) {
  const [data, setData] = useState(null)
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState(null)

  useEffect(() => { ... }, [path])

  return { data, loading, error }
}
```

By convention, custom hooks start with `use`. This is not just a style choice — React enforces rules about where hooks can be called (only inside components or other hooks), and it identifies hooks by the `use` prefix.

#### `useState` — Making Data Reactive

```js
const [data, setData] = useState(null)
```

`useState` stores a value inside a component that, when changed, tells React to re-render the component with the new value. It returns an array of exactly two things: the current value and a function to update it. We destructure that array immediately with `const [data, setData]`.

- `data` starts as `null` (no data yet)
- When `setData(someValue)` is called, React re-renders and `data` is now `someValue`

This is why React is called "reactive" — you don't manually update the DOM. You update state, and React figures out what changed and updates the DOM for you.

We have three separate state variables because they change independently:
- `loading` starts `true`, becomes `false` when the fetch completes (success or failure)
- `data` starts `null`, gets set when the fetch succeeds
- `error` starts `null`, gets set when the fetch fails

#### `useEffect` — Running Side Effects

```js
useEffect(() => {
  let cancelled = false

  fetch(path)
    .then((res) => {
      if (!res.ok) throw new Error(`HTTP ${res.status}`)
      return res.json()
    })
    .then((json) => { if (!cancelled) setData(json) })
    .catch((err) => { if (!cancelled) setError(err.message) })
    .finally(() => { if (!cancelled) setLoading(false) })

  return () => { cancelled = true }
}, [path])
```

A **side effect** is anything that reaches outside of React — fetching data, writing to `localStorage`, setting up a timer, subscribing to events. React wants to know about these so it can manage them correctly.

`useEffect` takes two arguments:
1. A function containing the side effect
2. A **dependency array** — `[path]` here

The effect runs after the component renders. Whenever any value in the dependency array changes, the effect runs again. With `[path]`, the fetch re-runs if a different URL is passed in.

**The cleanup function**: the function returned from the effect (`return () => { cancelled = true }`) is the **cleanup function**. React calls it when the component unmounts, and before re-running the effect. Here it sets `cancelled = true`.

**Why `cancelled`?** Imagine this scenario: the component renders and starts a fetch. Before the fetch completes, the user navigates away and the component unmounts. The fetch eventually completes and tries to call `setData(json)`. But the component no longer exists — this causes a React error. The `cancelled` flag prevents this: if the cleanup ran before the fetch completed, `cancelled` is `true` and the `.then` callbacks do nothing.

**Promises and `.then()`**: `fetch()` returns a **Promise** — an object representing an asynchronous operation that will complete in the future. You chain `.then()` to run code when the promise resolves (succeeds), `.catch()` for when it rejects (fails), and `.finally()` for code that runs either way.

`res.json()` itself returns another Promise — reading the response body is also asynchronous. So `.then(res => res.json())` gets the raw response, and the next `.then(json => setData(json))` gets the parsed JSON.

### `src/components/Bio.jsx`

```jsx
export default function Bio() {
  const { data, loading, error } = useResumeSection('/api/resume/bio')

  if (loading) return <p className={styles.state}>Loading...</p>
  if (error)   return <p className={styles.state}>Error: {error}</p>

  return (
    <header className={styles.bio}>
      <h1 className={styles.name}>{data.name}</h1>
      ...
    </header>
  )
}
```

The `if (loading) return ...` / `if (error) return ...` pattern is called **conditional rendering** — returning different JSX based on conditions. React components are just functions; returning early with different JSX is completely normal.

On first render, `loading` is `true`, so React shows "Loading...". The `useEffect` in the hook fires after render, starts the fetch, and when the response arrives, calls `setData(json)` and `setLoading(false)`. This triggers a re-render. Now `loading` is `false` and `data` has a value, so the real content renders.

```jsx
{data.links?.map((link) => (
  <a key={link.label} href={link.url} target="_blank" rel="noreferrer">
    {link.label}
  </a>
))}
```

**Optional chaining (`?.`)**: `data.links?.map(...)` means "if `data.links` exists, call `.map()` on it; otherwise return `undefined`." Without the `?`, if `links` were `null` or `undefined`, calling `.map()` would throw an error.

**`.map()` in JSX**: `.map()` transforms an array — here it converts each link object into a JSX `<a>` element. The result is an array of JSX elements, which React renders in sequence.

**`key`**: whenever you render a list with `.map()`, each element needs a unique `key` prop. React uses keys to efficiently track which items were added, removed, or reordered between renders. Without keys, React would have to re-render the entire list on every change.

**`target="_blank" rel="noreferrer"`**: `target="_blank"` opens the link in a new tab. `rel="noreferrer"` is a security best practice — without it, the new tab can access `window.opener` and potentially redirect your page.

### `src/components/Experience.jsx`

```jsx
{data.map((job, i) => (
  <div key={i} className={styles.item}>
    <div className={styles.itemHeader}>
      <h3>{job.role}</h3>
      <p>{job.company} &mdash; {job.location}</p>
      <span>{job.start_date} – {job.end_date}</span>
    </div>
    <ul>
      {job.highlights.map((h, j) => <li key={j}>{h}</li>)}
    </ul>
  </div>
))}
```

`data` here is a JavaScript array (the JSON array of jobs from the API). `.map((job, i) => ...)` passes both the item and its index to the callback. We use `i` as the key since the list is static.

The nested `.map()` inside renders each highlight as an `<li>`. `&mdash;` is an HTML entity for an em dash (—). In JSX you can use HTML entities the same as in HTML.

Notice that `job.start_date` uses snake_case (with underscore) — this matches the Go struct tag `json:"start_date"`, not the Go field name `StartDate`.

### `src/components/Education.jsx`

Structurally identical to Experience but for education records. It imports `Section.module.css` for its styles, the same stylesheet used by Experience. CSS Modules allow two components to share a stylesheet safely — explained below.

### `src/components/Skills.jsx`

```jsx
import styles from './Section.module.css'
import skillStyles from './Skills.module.css'
```

This component imports two CSS Modules and uses both. `styles` is used for the outer section wrapper (shared with Experience and Education); `skillStyles` is used for the tag grid and pill badges specific to Skills.

```jsx
{data.map((group, i) => (
  <div key={i}>
    <h3>{group.category}</h3>
    {group.skills.map((skill) => (
      <span key={skill}>{skill}</span>
    ))}
  </div>
))}
```

Here `key={skill}` uses the skill name itself — since skill names within a group are unique strings, this is better than using an index. React can then track individual skills by name if the list ever changes.

### CSS Modules — Scoped Styles

Files ending in `.module.css` are **CSS Modules**. They look like ordinary CSS but solve the problem of class name collisions.

In a normal CSS file, `.title` is a global class name. If two components both define `.title`, they conflict. In large projects this becomes a maintenance nightmare.

With CSS Modules:

```js
import styles from './Bio.module.css'
// styles.name is actually something like "Bio_name__xK2p1" at runtime
```

During the build, Vite transforms each class name into a unique, auto-generated string. `.name` in `Bio.module.css` and `.name` in `Section.module.css` become different strings and never conflict, even though you wrote them the same way. You get the convenience of simple, readable class names without any risk of global collisions.

---

## How the Two Halves Connect

Here is the complete journey of one piece of data from the Go server to the browser:

```
1. You run `go run .` in backend/
   → Go compiles the code and starts an HTTP server on port 8080
   → The ServeMux is ready to route incoming requests

2. You run `npm run dev` in frontend/
   → Vite starts a dev server on port 5173
   → The proxy is configured: any /api/* request → forward to port 8080

3. You open localhost:5173 in your browser
   → Vite serves index.html
   → The browser executes src/main.jsx
   → React renders <App />, which renders <Bio />, <Experience />, etc.

4. Bio renders for the first time
   → loading=true, data=null
   → React displays "Loading..."
   → After render, useEffect fires

5. useEffect calls fetch('/api/resume/bio')
   → The browser sends GET /api/resume/bio to localhost:5173
   → Vite's proxy intercepts it, forwards to localhost:8080

6. The Go server receives GET /api/resume/bio
   → The ServeMux matches the route to handlers.GetBio
   → GetBio calls writeJSON(w, 200, resume.Bio)
   → Go encodes the Bio struct as JSON using the struct tags
   → The response body: {"name":"Your Name","title":"Software Engineer",...}

7. The response travels back:
   Go (8080) → Vite proxy → browser

8. fetch() resolves in the browser
   → .then(res => res.json()) parses the JSON body into a JavaScript object
   → .then(json => setData(json)) updates React state

9. setData triggers a re-render of Bio
   → loading=false, data={name:"Your Name", title:"Software Engineer",...}
   → React renders the real content: <h1>Your Name</h1>, etc.

10. The user sees your resume.
```

This sequence — server sends JSON, client fetches it, React renders it — is the foundation of nearly every modern web application.
