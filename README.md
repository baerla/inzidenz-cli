<h1 align="center">Welcome to inzidenz-cli 👋</h1>
<p>
  <img alt="Version" src="https://img.shields.io/badge/version-0.1.0-blue.svg?cacheSeconds=2592000" />
  <a href="#" target="_blank">
    <img alt="License: AGPL3" src="https://img.shields.io/badge/License-AGPL3-yellow.svg" />
  </a>
</p>

> A cli written in golang which extracts the current incidence of the cities homepage.

## Installation

```sh
go install github.com/baerla/inzidenz-cli@latest
echo 'export COVID_INCIDENCE_CONFIG="$HOME/.inzidenz.json"' >> ~/.zshrc
```

## Usage

```sh
inzidenz-cli add <name> <url of the webpage containing the incidence value>
inzidenz-cli get
inzidenz-cli get <name>
inzidenz-cli get <name> <url of the webpage containing the incidence value>
```
