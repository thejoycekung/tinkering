# modch

`modch` is a command-line utility to convert `chmod` octals into symbolic permissions, or those symbols into octals.

To build: `go build modch`
To run: `./modch [octal]` or `./modch [symbolic perms]`

## Examples

```
$ ./modch 550
Number 550 received! Converting to symbolic format ...
This number represents these permissions:  r-xr-x---
```

```
$ ./modch r--rwxr--
Symbol r--rwxr-- received! Converting to octal format ...
These permissions can be represented as this octal number:  474
```

## To-do
[ ] Module tests  
[ ] Verbose mode
