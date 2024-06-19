---
syntax: bash
tags: [ macos, cli-tools ]
---

### To delete the Apple_APFS, 
### Note: Only delete the 2nd one first one is used by macos
diskutil apfs deleteContainer disk4

### delete other containers, Replace <no> with disk number
```bash
diskutil eraseVolume free free disk0s<no>
```

### Resizing the main macos containers, Replace <no> with disk number of macos

```bash
diskutil apfs resizeContainer disk0s<no> 0
```


