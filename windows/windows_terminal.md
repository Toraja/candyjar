## Configuration
### General
#### Change default to ubuntu
By setting default profile  
1. Run `uuidgen` and take a note of the output
1. Replace the value of `defaultProfile` in `setting.json` (located at the top).

By changing profile order  
1. In the `profiles` section of `setting.json`, sort profiles as you like.  
  The profile at the top will be used when the terminal is opened.  
  Also, the order of these profiles corresponds to the keyboard shortcut to open
  shell such as `Ctrl+Shift+1`.  

### Ubuntu
Go to ubuntu section `profiles -> list`

#### Change starting directory
1. add `"startingDirectory": "//wsl$/Ubuntu/<linux path to directory>"` to the
   section.  
   eg. `"startingDirectory": "//wsl$/Ubuntu/home/user"`

#### Change font
1. add `"fontFace": "Consolas"` to the section

#### Change cursor shape
1. add `"cursorShape": "filledBox"` to the section


### Keybind
Paste value of `actions` below into `setting.json`.
```json
{
    "actions":
    [
        {
            "command": "unbound",
            "keys": "ctrl+comma"
        },
        {
            "command": "unbound",
            "keys": "win+sc(41)"
        },
        {
            "command":
            {
                "action": "copy",
                "singleLine": false
            },
            "keys": "ctrl+shift+c"
        },
        {
            "command": "find",
            "keys": "ctrl+shift+f"
        },
        {
            "command": "openTabRenamer",
            "keys": "ctrl+shift+r"
        },
        {
            "command": "paste",
            "keys": "ctrl+shift+v"
        },
        {
            "command":
            {
                "action": "splitPane",
                "split": "auto",
                "splitMode": "duplicate"
            },
            "keys": "alt+shift+d"
        }
    ],
}
```
