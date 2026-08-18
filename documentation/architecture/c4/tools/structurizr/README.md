# Structurizr C4 Model Tutorial

This is modified version of the [Structurizr C4 Model Tutorial](https://docs.structurizr.com/dsl/tutorial).

## Things to Note

### Basic
- File name of the diagram must be `workspace.dsl`.[^1]
- When Structurizr is run, it will generate `workspace.json` file and many others in the current directory.
  - From the Structurizr documentation:  
    > The “compiled” version of your workspace source, additionally including diagram layout information from the diagram editor. This is an open JSON data format and is not designed to be edited by hand.[^2]
  - Elements can be moved in the Structurizr web interface and the position will be saved/updated in this file

### Docker
- Because Structurizr generates files, mounted directory must be writable by the user in the container. So either:
  - Change the permissions of the directory on the host to be writable by the user in the container (e.g. `chmod 777 /path/to/mounted/dir`)
  - Change the user in the container to match the owner of the mounted directory (e.g. `docker run --user $(id -u):$(id -g) ...`)

### Structurizr Views
- Each view specifies a single software system to focus and only the elements that have relationships with that software system are displayed in the view by default.
  To show the elements that are not directly related to the focused software system, use the `include` keyword to include them.
- In a container view, only containers that in the specified software system are displayed and others remain black-boxed.
  Even if a relationship is defined between containers in different software systems, it will be depicted as a relationship between a container and a software system, not a container inside it.
  To show the relationships between containers in multiple software systems in a single view, use the `include` keyword to include containers in non-focused software system.[^3]
  Easiest way to include containers in non-focused software system is to use `include element.parent==<software system identifier>` to include all containers in the specified software system.
  One caveat is that `*` in `include` keyword must be the last one in the line, or containers in non-focused software system will not be displayed in the view somehow.
  (e.g. `include c1 c2 *` is valid, but `include * c1 c2` is invalid)

[^1]: https://docs.structurizr.com/local
[^2]: https://docs.structurizr.com/workspaces/file-types#workspacejson
[^3]: https://docs.structurizr.com/dsl/cookbook/container-view-multiple-software-systems/
