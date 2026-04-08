# Task Scheduler

## Tips

### Running PowerShell script

The script to run must be passed as arguments rather than selecting it as `Program/script`.  
Note that `-WindowStyle Hidden` does not completely hide the window and it still flashes the window.  

_Program/script_
```
C:\Windows\System32\WindowsPowerShell\v1.0\powershell.exe
```

_Add arguments_
```
-WindowStyle Hidden -ExecutionPolicy Bypass -File "C:\path\to\script.ps1"
```
