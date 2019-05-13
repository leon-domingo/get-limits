# Small binary which returns the limits a battery should be charged to and discharged to for the current day

Usage is very simple:

```shell
get-limits [both|top|bottom]
```

For example:

```shell
get-limits both
```

A list of arbitrary values ranging from 10,20,25,30 (bottom) to 90,95,98 (top) is defined in order to return the limits for the current day.
