# Timer

App for running scripts in given intervals and showing output of last run of each script.

## Installation

```
go install github.com/klapacz/timer@latest
```

## Running

```
# read config from ~/.config/timer.yaml
timer
# read config from provided path
timer ./path/to/config.yaml
```

## Config example

```yaml
separator: " | "
cmds:
  - cmd: date "+%S"
    interval: 5
  - cmd: date "+%S"
    interval: 1
```

Output:

```
44 | 44
44 | 45
44 | 46
44 | 47
44 | 48
49 | 48
49 | 49
49 | 50
49 | 51
49 | 52
49 | 53
54 | 53
54 | 54
```
