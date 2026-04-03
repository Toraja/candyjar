# Windows Tips

### Find out which process is using a file/directory

There are 2 ways to do that.

#### handle.exe

Usage (for 64bit PC):

```shell
handle64.exe <filename>
```

#### Built-in Resource Monitor

1. Open `Resource Monitor`
1. Select `CPU` tab
1. Enter file name into the textbox on `Associated Handles` and press "Enter"
   -> Processes using the file will be displayed.

Reference: [filesystems - Find out which process is locking a file or folder in Windows - Super User](https://stackoverflow.com/questions/3565218/how-to-know-what-process-is-using-a-given-file)

### Hide installed programs from Programs and Features (and Apps & Features)

See [this website](http://woshub.com/how-to-hide-installed-programs-from-programs-and-features/)

### Prevent apps from pinning themselves automatically

1. Open file `$env:USERPROFILE\AppData\Local\Microsoft\Windows\Shell\LayoutModification.xml`  
   If this does not work, modify `C:\Users\Default\AppData\Local\Microsoft\Windows\Shell\LayoutModification.xml`
1. Find `CustomTaskbarLayoutCollection` tag
1. Remove the entries inside `TaskbarPinList`

### Change icon of program / shortcut to any image

1. Copy and paste image into `paint` program.
1. Save the image as bitmap (BMP).
1. Rename and change the extension of the image to `ico`.
1. Select the image as icon.

In some case, shortcut icon for some program (eg. Microsoft Edge, Google Chrome) which are pinned to start menu somehow remains unchanged even when you change it via `properties`.
A workaround is to start such program via powershell.

*Example*
```powershell
C:\Windows\System32\WindowsPowerShell\v1.0\powershell.exe -WindowStyle Hidden -Command "Start-Process ${env:ProgramFiles(x86)}\Microsoft\Edge\Application\msedge.exe -Args www.bing.com"
```
