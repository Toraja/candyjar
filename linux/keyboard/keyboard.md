# Linux Keyboard

## Customise Keyboard Layout
1. Create `~/.Xmodmap` file
2. Put config like below
3. Use `xev` to see what keycode a key generates

```xmodmap
# syntax
keycode <keycode> = (Keysyms) Key Shift+Key mode_switch+Key mode_switch+Shift+Key AltGr+Key AltGr+Shift+Key

# eg.
keycode 133 = Super_L NoSymbol Super_L
keycode 105 = Overlay1_Enable NoSymbol Overlay1_Enable
```
