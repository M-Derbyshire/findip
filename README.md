# findip

`findip` is a command-line tool that can be used to determine the current IP of a device on your network.

Let's say that you had a device running on your network, and that it hosted a web app for you to use when you were connected. You could configure the device to always have the same IP, and then use that in the URL that you connect to.

But what if you aren't allowed to configure the network devices in this way? Or what if the device that hosts the service is changed at some point?

`findip` can be used to scan a range of IP addresses (v4) to see which is hosting a certain service. You provide the range of IP addresses, a port, an endpoint, and the text you want to search for. The program will then send HTTP requests (in batches) to the IPs in the range, until it finds one that has the given search text in its response body. Once it finds the correct IP address, that address is outputted. You could then use this address in a script, to trigger your browser opening the correct URL.

## Usage

### Flags
- **`-from` (string)** - The IP (v4) address that we want to start scanning from (default: 192.168.0.1)
- **`-to` (string)** - The IP (v4) address that we want to end scanning on (default: 192.168.0.100). The scan will include this IP
- **`-port` (positive integer)** - The port that we want to send HTTP requests to (default: 80)
- **`-endpoint` (string)** - The endpoint that we want to send HTTP requests to (default: empty string)
- **`-protocol` (string)** - The protocol to use in HTTP requests (default: http). E.g. http or https.
- **`-batchsize` (positive integer)** - This determines how many HTTP requests should be sent at once (default: 32). So, if you had a range of 200 IP addresses, and you set this to 50, then the program would send 50 requests at once, wait for them to all complete, and then move on to the next 50 (doing this 4 times).
- **`-searchtext` (string)** - The text that we want to find in a HTTP response body (default: "hello world!"). This could be an identifier that a device is hosting at a specific port/endpoint, or a known unique string in one of its standard responses.

### Example Browser-Triggering Script

This is an example of a Bash script that finds a device, and then opens a Chrome tab for the service hosted on that device.

```
#!/bin/bash

# Run the command and capture the found IP address
ip=$(findip -port 8080 -searchtext "my search text" -from 192.168.0.1 -to 192.168.0.50)
status=$?

# If there was an error, we stop
if [[ $status -ne 0 ]]; then
    exit "$status"
fi

# Now trigger Chrome with the desired URL
url="http://$ip:8081"
google-chrome --new-tab "$url"
```


## Development

### Testing

There are 3 MAKE commands related to the project's tests:
- `update-test-snapshots`
- `test`
- `e2e-test`

Unfortunately, due to their size, snapshots have not been commited to this repository. Therefore, before working on this project, you will want to run the `make update-test-snapshots` command.

Once the snapshots are generated, you can run the project's unit tests with `make test`.

To run the project's End-to-End tests, you can use `make e2e-test`.