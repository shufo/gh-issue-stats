# gh-issue-stats

📊 Export GitHub issue statistics via the [GitHub CLI](https://github.com/cli/cli)

A `gh` CLI extension that lets you export issue stats aggregated by labels with various format (CSV, JSON, TSV).

![Screencast](https://github.com/user-attachments/assets/27d4f6ba-1cfa-4c85-b9a8-3402248247b0)

## Requirements

- [GitHub CLI](https://github.com/cli/cli)

## Installation

```bash
$ gh extension install shufo/gh-issue-stats
```

## Output Example

Stats from [GitHub CLI](https://github.com/cli/cli) issues.

```bash
$ gh issue-stats cli/cli
╭─────────────────────┬──────┬────────┬───────┬─────────┬──────────────────────────────┬─────────────────────────────╮
│ Label               │ Open │ Closed │ Total │ Open %  │ Average Time To Close (Days) │ Median Time To Close (Days) │
├─────────────────────┼──────┼────────┼───────┼─────────┼──────────────────────────────┼─────────────────────────────┤
│ enhancement         │  431 │   1438 │  1869 │ 23.06%  │ 92                           │ 6                           │
│ bug                 │  206 │   1621 │  1827 │ 11.28%  │ 38                           │ 2                           │
│ needs-triage        │  137 │    801 │   938 │ 14.61%  │ 5                            │ 0                           │
│ help wanted         │  143 │    470 │   613 │ 23.33%  │ 120                          │ 31                          │
│ needs-user-input    │   30 │    467 │   497 │ 6.04%   │ 36                           │ 5                           │
│ p3                  │  100 │    297 │   397 │ 25.19%  │ 103                          │ 16                          │
│ core                │  159 │    182 │   341 │ 46.63%  │ 274                          │ 161                         │
│ *unlabeled*         │    0 │    322 │   322 │ 0.00%   │ 10                           │ 0                           │
│ gh-pr               │  165 │     62 │   227 │ 72.69%  │ 140                          │ 23                          │
│ p2                  │   37 │    159 │   196 │ 18.88%  │ 109                          │ 20                          │
│ feedback            │   12 │    156 │   168 │ 7.14%   │ 21                           │ 1                           │
│ docs                │   38 │    102 │   140 │ 27.14%  │ 111                          │ 21                          │
│ platform            │   39 │     51 │    90 │ 43.33%  │ 113                          │ 22                          │
│ gh-repo             │   55 │     32 │    87 │ 63.22%  │ 65                           │ 26                          │
│ blocked             │   45 │     32 │    77 │ 58.44%  │ 173                          │ 69                          │
│ discuss             │   34 │     42 │    76 │ 44.74%  │ 172                          │ 21                          │
│ packaging           │    8 │     61 │    69 │ 11.59%  │ 184                          │ 36                          │
│ needs-design        │   11 │     55 │    66 │ 16.67%  │ 319                          │ 271                         │
│ gh-auth             │   32 │     30 │    62 │ 51.61%  │ 118                          │ 27                          │
│ gh-issue            │   41 │     18 │    59 │ 69.49%  │ 129                          │ 29                          │
│ good first issue    │    2 │     51 │    53 │ 3.77%   │ 74                           │ 11                          │
│ tech-debt           │   20 │     26 │    46 │ 43.48%  │ 77                           │ 24                          │
│ codespaces          │   19 │     27 │    46 │ 41.30%  │ 110                          │ 48                          │
│ gh-run              │   27 │     18 │    45 │ 60.00%  │ 106                          │ 45                          │
│ p1                  │    3 │     41 │    44 │ 6.82%   │ 4                            │ 1                           │
│ gh-release          │   25 │     15 │    40 │ 62.50%  │ 102                          │ 21                          │
│ needs-investigation │    8 │     31 │    39 │ 20.51%  │ 151                          │ 28                          │
│ extension-idea      │    9 │     28 │    37 │ 24.32%  │ 307                          │ 309                         │
│ windows             │    4 │     32 │    36 │ 11.11%  │ 236                          │ 104                         │
│ gh-codespace        │   20 │     14 │    34 │ 58.82%  │ 124                          │ 47                          │
│ gh-attestation      │    9 │     24 │    33 │ 27.27%  │ 12                           │ 7                           │
│ gh-api              │   22 │     11 │    33 │ 66.67%  │ 173                          │ 42                          │
│ actions             │    8 │     24 │    32 │ 25.00%  │ 117                          │ 33                          │
│ gh-extension        │   20 │     10 │    30 │ 66.67%  │ 185                          │ 46                          │
│ auth                │    7 │     21 │    28 │ 25.00%  │ 160                          │ 21                          │
│ invalid             │    0 │     22 │    22 │ 0.00%   │ 0                            │ 0                           │
│ gh-project          │    6 │     13 │    19 │ 31.58%  │ 70                           │ 18                          │
│ gh-workflow         │   13 │      4 │    17 │ 76.47%  │ 54                           │ 45                          │
│ gh-search           │   12 │      3 │    15 │ 80.00%  │ 22                           │ 11                          │
│ accessibility       │    9 │      3 │    12 │ 75.00%  │ 141                          │ 15                          │
│ extensions          │    4 │      8 │    12 │ 33.33%  │ 124                          │ 60                          │
│ gh-gist             │    8 │      3 │    11 │ 72.73%  │ 6                            │ 5                           │
│ gh-browse           │    5 │      5 │    10 │ 50.00%  │ 323                          │ 163                         │
│ config              │    0 │     10 │    10 │ 0.00%   │ 242                          │ 173                         │
│ gh-variable         │    1 │      8 │     9 │ 11.11%  │ 19                           │ 13                          │
│ gh-config           │    8 │      1 │     9 │ 88.89%  │ 36                           │ 36                          │
│ gh-secret           │    5 │      3 │     8 │ 62.50%  │ 48                           │ 60                          │
│ gh-status           │    5 │      0 │     5 │ 100.00% │ 0                            │ 0                           │
│ gh-cache            │    1 │      4 │     5 │ 20.00%  │ 5                            │ 3                           │
│ duplicate           │    0 │      4 │     4 │ 0.00%   │ 0                            │ 0                           │
│ cli.github.com      │    2 │      1 │     3 │ 66.67%  │ 1                            │ 1                           │
│ gh-org              │    3 │      0 │     3 │ 100.00% │ 0                            │ 0                           │
│ gh-help             │    3 │      0 │     3 │ 100.00% │ 0                            │ 0                           │
│ gh-gpg-key          │    1 │      1 │     2 │ 50.00%  │ 4                            │ 4                           │
│ gh-alias            │    2 │      0 │     2 │ 100.00% │ 0                            │ 0                           │
│ gh-completion       │    2 │      0 │     2 │ 100.00% │ 0                            │ 0                           │
│ gh-ssh-key          │    2 │      0 │     2 │ 100.00% │ 0                            │ 0                           │
│ project             │    0 │      1 │     1 │ 0.00%   │ 25                           │ 25                          │
│ gh-reference        │    0 │      1 │     1 │ 0.00%   │ 99                           │ 99                          │
│ gh-ruleset          │    1 │      0 │     1 │ 100.00% │ 0                            │ 0                           │
│ gh-label            │    0 │      1 │     1 │ 0.00%   │ 175                          │ 175                         │
├─────────────────────┼──────┼────────┼───────┼─────────┼──────────────────────────────┼─────────────────────────────┤
│ Total               │  691 │   3812 │  4503 │ 15.35%  │ 56                           │ 2                           │
╰─────────────────────┴──────┴────────┴───────┴─────────┴──────────────────────────────┴─────────────────────────────╯
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

- Persist aggregated results to file

```bash
$ gh issue-stats -s stats.json
```

- Persist raw source data to file

```bash
$ gh issue-stats -o issues.json
```

- Verbose output

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
