---
syntax: bash
tags: [cli-tools]
---

### fzf preview window
```bash
fzf --preview "command to run {1}"
```


### Adding preview options preview options
```bash
fzf --preview "command" --preview-window=(right | left),( width in % )
```


### Example with tldr
```bash
tldr --list | fzf --preview "tldr {1} -C" --preview-window=right,70% | xargs tldr
```


### Note you can get the argument piped to with {1}
