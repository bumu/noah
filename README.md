# Noah APIs

Noah APIs is using a microservice architecture, just for testing and learning.

## APIs Framework

```mermaid
sequenceDiagram
    actor User as Participant
    participant Makefile
    participant RunGo as Run.go
    participant Bootstrap
    participant App
    participant Handlers
    participant Repo
    participant Schema as Database

    Makefile ->> RunGo : run
    RunGo ->> Bootstrap : Bootstrap
    Bootstrap ->> App : App FxOptions
    App ->> Handlers : Handlers
    Handlers ->> Repo : Repo
    Repo ->> Schema : Schema
```

## Install bufbuild/buf/buf

```bash
brew install bufbuild/buf/buf
```

## Generate

```bash
buf generate
```

## Install go-task

```bash
brew install go-task
```
