# mycli

A Node.js CLI application boilerplate built with TypeScript.

## Requirements

- Node.js >= 20.0.0
- pnpm

## Installation

```bash
pnpm install
pnpm build
```

## Usage

```bash
# Show help
mycli --help

# Show version
mycli --version

# Greet a user
mycli greet <name>

# Greet in uppercase
mycli greet --uppercase <name>
```

## Development

```bash
# Build
pnpm build

# Watch mode
pnpm dev

# Run directly
pnpm start

# Lint
pnpm lint

# Type check
pnpm typecheck

# Clean build output
pnpm clean
```

## Project Structure

```
src/
  bin/cli.ts            # Entry point
  commands/
    index.ts            # Program setup
    greet.ts            # Example subcommand
  lib/
    runner.ts           # Main runner
    exit.ts             # Graceful shutdown handling
    logger.ts           # Structured logger
```

## License

MIT
