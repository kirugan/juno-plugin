## Example of a plugin for Juno

Workflow:

1. Write your code in `myplugin.go`
2. Run `make build`
3. Pass produced `myplugin.so` to Juno like `juno --shared-plugin=<path_to_myplugin.so>`