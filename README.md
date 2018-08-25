# zulu

A command line utility to make it easier to get and parse UNIX timestamps (of
different precision). The need for this tool came out of me frequently needing
to work with UNIX timestamps in either second or millisecond precision.

## License

`zulu` is released under the MIT License. See the [LICENSE
file](https://github.com/theckman/zulu/blob/master/LICENSE) for more details.

## Installing

At the time of writing this utility requires you have a Go toolchain available
to build it from source. I'll work on getting binary releases up soon.

```
go get github.com/theckman/zulu/cmd/z
```

Assuming your `$GOBIN` location (default `$GOPATH/bin`) is in your `$PATH`, you
should now have a binary named `z` available for use.

## Usage

By default `z` prints the current time in UTC and localtime, as well as with
UNIX timestamps of different precisions:

```
$ z
Local time:		2018-08-27 08:00:00 PDT (7 hours behind UTC)

UTC time:		2018-08-27 15:00:00 UTC
Epoch timestamp:	1535382000
Millisecond timestamp:	1535382000000
Nanosecond timestamp:	1535382000000000000
```

You can also have it provide you with only the formatted timestamp of your choosing:

```
$ z --unix
1535382000
```

In addition to that, you can pass it a timestamp to have it parse and print the
datetime at that timestamp. The value of the timestamp can be any of the
formatter flags (like `--unix` above). It defaults to a mode where it tries to
guess whether it's a `--unix` or `--unix-milliseconds`:

```
$ z 1535381999
Local time:		2018-08-27 07:59:59 PDT (7 hours behind UTC)

UTC time:		2018-08-27 14:59:59 UTC [was 1s ago]
Epoch timestamp:	1535381999
Millisecond timestamp:	1535381999000
Nanosecond timestamp:	1535381999000000000
```

### Local Timezone

Zulu defaults to using your local timezone when printing times, however by
setting a `TZ` environment variable you can override this. It expects the same
format as the UNIX `TZ` EnvVar (e.g., `America/New_York`). My system is in the
`America/Los_Angeles` timezone, but by setting the `TZ` environment variable you
can see it changes my local zone:

```
$ TZ=America/New_York z
Local time:		2018-08-27 11:00:00 EDT (4 hours behind UTC)

UTC time:		2018-08-27 15:00:00 UTC
Epoch timestamp:	1535382000
Millisecond timestamp:	1535382000000
Nanosecond timestamp:	1535382000000000000
```

### Timestamp Format

Whether you're passing in a timestamp to be parsed, or whether you want to print
the current timestamp, there are a few Formatter flags available to give zulu a
hint of which type to parse.

If you provide a formatter flag, but no timestamp, it prints the current time in
that format. If you provide a formatter flag and a timestamp, it attempts to
parse the timestamp.

```
$ z --rfc-3339 '2018-08-27 15:00:00Z'
Local time:		2018-08-27 08:00:00 PDT (7 hours behind UTC)

UTC time:		2018-08-27 15:00:00 UTC [was 1s ago]
Epoch timestamp:	1535382000
Millisecond timestamp:	1535382000000
Nanosecond timestamp:	1535382000000000000
```

See the help output (`-h/--help`) for the list of all of the Formatter flags.

### Help Output

```
Usage:
  z [OPTIONS] [TIMESTAMP]

Application Options:
  -V, --version               print version and exit

Timestamp Formatters:
  -u, --unix                  time format is the UNIX Epoch (seconds since 00:00:00 UTC January 1, 1970)
  -m, --unix-milliseconds     time format is the UNIX Epoch to the nearest second with millisecond precision
  -n, --unix-nanoseconds      time format is the UNIX Epoch to the nearest second with nanosecond precision
  -r, --rfc-3339              time format is the RFC 3339 format in UTC timezone, with the T replaced by a space
  -i, --iso-8601              time format is the same as the -r/--rfc-3339 format, except it includes the 'T' between
                              the date and time
  -t, --internal-time-format  time format is the same as is printed by default (e.g. 2018-08-27 15:00:00 UTC)
  -z, --zulu                  same as -r/--rfc-3339 but with no timezone specifier

Help Options:
  -h, --help                  Show this help message

Arguments:
  TIMESTAMP:                  An optional timestamp to parse and print the datetime of. The expected format depends
                              on which Timestamp Formatter flag you use. The default is a UNIX Epoch parser that
                              attempts to assume the desired precision (seconds, milliseconds, nanoseconds)
```
