# Small binary which returns the limits a battery should be charged/discharged to for the given day (or today, by default)

Usage is very simple:

```shell
get-limits [both|top|bottom] <YYYYMMDD>
```

For example:

```shell
get-limits both 20190530
```

If no date is supplied then the current date will be used:

```shell
get-limits both
```

A list of arbitrary values ranging from 10,20,25,30 (bottom) to 90,95,98 (top) is defined in order to return the limits for the given day.
