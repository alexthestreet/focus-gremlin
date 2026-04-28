# Focus Gremlin

Focus Gremlin is a Linux-only Go CLI that runs a background daemon and opens a terminal-based check-in prompt on a schedule.

The daemon owns scheduling, runtime state, and persistence. The prompt is a short-lived Bubble Tea TUI launched in a new terminal window.

## Commands

- `focus-gremlin daemon`
- `focus-gremlin prompt`
- `focus-gremlin status`
- `focus-gremlin config init`
- `focus-gremlin config show`

## Build

```bash
go build -o focus-gremlin ./cmd/focus-gremlin
```

## Quick Start

```bash
./focus-gremlin config init
./focus-gremlin config show
./focus-gremlin daemon
```

To run a prompt directly in the current terminal:

```bash
./focus-gremlin prompt
```

## Paths

- Config: `~/.config/focus-gremlin/config.json`
- Data: `~/.local/share/focus-gremlin/history.db`
- Runtime: `$XDG_RUNTIME_DIR/focus-gremlin/`

Runtime files:

- `daemon.lock`
- `state.json`

## Example Config

```json
{
  "interval_minutes": 60,
  "active_start": "09:00",
  "active_end": "17:00",
  "snooze_minutes": 10,
  "prompt_timeout_seconds": 0,
  "statuses": [
    "on_track",
    "off_track",
    "deep_in_the_void",
    "snooze"
  ],
  "terminal_command": [
    "x-terminal-emulator",
    "-e"
  ]
}
```

## Terminal Launcher

The daemon launches `focus-gremlin prompt` by prepending the configured `terminal_command`.

Example values:

```json
[
  "x-terminal-emulator",
  "-e"
]
```

```json
[
  "kitty",
  "-e"
]
```

```json
[
  "alacritty",
  "-e"
]
```

## Status

```bash
./focus-gremlin status
```

This reports the resolved config path, data path, runtime directory, and whether the lock and state files currently exist.

## Notes

- Linux only
- Built for terminal and tiling-window-manager workflows
- Responses are stored locally in SQLite
