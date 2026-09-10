# Task Scheduler

## Import Task

When importing a task that was exported from another computer, the user account that runs the task might be SID-based and not exist on the current computer.
You can get the current user by running `whoami /user` on PowerShell or Command Prompt, but setting the user account in Task Scheduler GUI might not work.
In such case, you can edit the XML file of the task and change the `UserId` to the current user name.

```xml
<Task ...>
  ...
  <Principals>
    <Principal id="Author">
      <UserId>USER NAME</UserId>
      ...
    </Principal>
  </Principals>
  ...
</Task>
```

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
