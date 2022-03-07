## Usage

Wait-for is reimplantation of https://github.com/vishnubob/wait-for-it - a pure bash script that will wait on the availability of a host and TCP port.

It is useful for synchronizing the spin-up of interdependent services, such as linked docker containers.
Since it is a pure bash script, it does not have any external dependencies.

```text
wait-for --help
wait-for [options] Run command with args after the test finishes
  -help
        Print usage
  -host string
        Hostname to check
  -port string
        Port
  -quite
        Disable logging
  -timeout int
        Timeout in seconds (default 15)
```

## Examples

For example, let's test to see if we can access port 80 on `www.google.com`,
and if it is available, run "ps".

```text
./wait-for -host google.com -port 80 -timeout 5 ps -ax 92923
2022/03/07 19:17:23 :::::::::::::::::::::::::::::: STDOUT ::::::::::::::::::::::::::::::
2022/03/07 19:17:23 Running command: ["ps" "-ax" "92923"]
2022/03/07 19:17:23   PID TTY           TIME CMD
92923 ??         0:00.04 /System/Library/Frameworks/Metal.framework/Versions/A/XPCServices/MTLCompilerService.xpc/Contents/MacOS/MTLCompilerService
2022/03/07 19:17:23 :::::::::::::::::::::::::::::: STDOUT ::::::::::::::::::::::::::::::
```

Failed test:

```text
./wait-for -host google.com -port 81 -timeout 5 ps -ax 92923
2022/03/07 19:18:21 Timeout occurred after waiting 5s
```

Failed command:

```text
./wait-for -host google.com -port 80 -timeout 5 ps -ax 000000
2022/03/07 19:19:56 :::::::::::::::::::::::::::::: STDERR ::::::::::::::::::::::::::::::
2022/03/07 19:19:56 Failed to run [ps -ax 000000]: exit status 1
2022/03/07 19:19:56 Command output was:
  PID TTY           TIME CMD
2022/03/07 19:19:56 :::::::::::::::::::::::::::::: STDERR ::::::::::::::::::::::::::::::
```

