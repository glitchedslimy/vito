# 💊 VITO

**VITO** it's a CLI ultra-vitamined designed to accelerate your workflow. It generates project structures (boilerplates) inteligently and expands with **Pills** (plugins) to adapt to what you currently need.

[![Go Report Card](https://goreportcard.com/badge/github.com/glitchedslimy/vito)](https://goreportcard.com/report/github.com/glitchedslimy/vito)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

---

## ✨ Features
- **🚀 Fast boilerplates:** Create base projects from curated templates or with _AI🤖_ _(AI is currently on Alpha stages)_.
- **💊 Pills for everyone**: An ecosystem of plugins to add functionalities to VITO, with various vitamines from Lua to Go.
- **🎨 Elegant TUI**: Modern and clean interactive terminal.

## 🛠️ Installation
**_If you have Go installed:_**
```bash
go install github.com/glitchedslimy/vito@latest
```
_(Distributed binaries for macOS, Linux and Windows will be in the future)_

## 🚀 Quickstart
**Create a new project**
```bash
vito create [project-name]
```

**See help**
```bash
vito --help
```

## 🔌 The "Pills" concept
_VITO_ is not only a generator; is a **host** for functionalities. A **Pill** can be anything: a personalized linter, an automated deploy launcher or a secret manager locally.

_To create your own pill, use: `vito create --template pill-base`_

## 🏗️ Roadmap
The roadmap can be seen in the official [ROADMAP](./ROADMAP.md).

## 🤝 Contributions
¿You have a new vitamine or want to improve the core? ¡PRs are welcome!
1. Fork the project
2. Create your feature branch with `git checkout -b feature/new-implementation-name`.
3. Commit your changes `git commit -m 'feat(): added this new thing'`.
4. Push the branch `git push origin feature/new-implementation-name`.
5. Open a Pull Request

---

Devloped with 💖 by [glitchedslimy](https://github.com/glitchedslimy)
