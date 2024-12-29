# gh-issue-stats

📊 Export GitHub issue statistics via the [GitHub CLI](https://github.com/cli/cli)

A `gh` CLI extension that lets you export issue stats with various format (CSV, JSON, TSV).

![Screencast](https://github.com/user-attachments/assets/27d4f6ba-1cfa-4c85-b9a8-3402248247b0)

## Requirements

- [GitHub CLI](https://github.com/cli/cli)

## Installation

```bash
$ gh extension install shufo/gh-issue-stats
```

## Usage

- Basic usage

```bash
$ gh issue-stats
```

- Specific repository

```bash
$ gh issue-stats owner/repo
```

- Change output format. (default: table. Supports `json`, `csv` and `tsv`)

```bash
$ gh issue-stats --format json
$ gh issue-stats owner/repo --format csv
$ gh issue-stats owner/repo --format tsv
```

- Debug

```bash
$ gh issue-stats --debug
```

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

## Development

```bash
# build extension
$ make build
# or go-task build to watch source
# Install extension locally
$ gh extension install .
# Run 
$ gh issue-stats
```

## Testing

```bash
$ make test
# or go-task test
```

## LICENSE

MIT
