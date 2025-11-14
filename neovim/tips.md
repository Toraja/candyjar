# TIPS

#### Project-wide string replace
1. Use a searching tool (like Telescope's live_grep) to search for the string you want to replace, then add the matches to quickfix list.
2. Use the :cfdo command to perform the substitution across all files listed in the Quickfix list.  
    Append `| bd` if you want to free resources.  
    ```vim
    :cfdo %s/oldString/newString/gc | update
    ```
