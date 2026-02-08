# Personal Shared Libs project
This personal projects aims to simplify my life in the future when initiating a new repositories

**A collection of reusable Go libraries** maintained by Richard
This repository contains shared packages and utilities to support multiple projects without duplication.

## 📦 What's Inside
This repo is structured into self-contained Go modules that provide standalone functionality:

- **`config`** — reusable configuration loading and parsing utilities  
- **`multiapikey`** — (example name) utilities for managing multiple API key sets  
(*Add/expand descriptions here as the modules evolve.*)

## 🚀 Motivation

Many codebases end up duplicating common utilities. This repo centralizes shared logic so:

- Core logic is written **once**
- Projects can re-use and maintain consistent behaviors
- You avoid copy-paste anti-patterns

> Shared Go libraries encourage cleaner, DRY code across your personal and team projects.

## 🧱 Architecture

Each package follows Go modules best practices:
- Independently usable
- Properly versioned
- Standards-compliant imports
- Simple and testable

## 🛠 Getting Started

Import any sub-package via Go modules:

```bash
go get github.com/richardwjy/shared-libs/<sub-package>@latest
