# go-nukibridgeapi

This project aims to make Nuki's Bridge API accessible. Either through the 
provided library or as a program from the command-line.

![Develop](https://github.com/christianschmizz/go-nukibridgeapi/workflows/Build/badge.svg?branch=develop)

Therefore it is introducing a command-line tool called `nukibridgectl` which 
implements access to the basic functions of your Nuki bridge for now.

# Activate the Bridge API

Before you can access the bridge's API you have to activate it. *It's not active
by default.*

## Check API status

If you are unsure about whether the bridge's API is already activated or not 
you can check with curl:

    $ curl http://<bridge-address>
    
It will come up with `HTTP 400 Bad Request` when active. Otherwise you'll see 
something like `Connection refused`. 

## Activate the API

To activate the bridge API you have to bring you bridge into maintenance mode 
go to "Manage Bridge" in the Nuki app. During the Bridge's setup procedure you 
can configure the settings of your Wi-Fi as well as the developer mode. To use 
the API you have to *enable the developer mode*.

The token you need for accessing the API is shown only once when activating 
the developer mode. So, make a note of it.

# Discover bridges

Note: The discovery of bridges does not require a valid token.

For discovering bridges you need a working internet-connection, and you have to 
execute `nukibridgectl` from the LAN the bridge is in. Afterwards the `discover`
command should report your bridge.

    $ nukibridgectl discover
    ID        IP             Port Updated
    123456789 192.168.178.1  8080 2020-10-01 20:00:00 +0000 UTC

# Help

To get a quick overview of the capabilities `nukibridgectl` offers I recommend
having a look at the help:

    $ nukibridgectl --help

Which might look like this:

    Work seamlessly with your Nuki Bridge from the command line.
    
    Usage:
      nukibridgectl <command> <subcommand> [flags]
      nukibridgectl [command]
    
    Available Commands:
      bridge      Bridge commands
      discover    Discover bridges
      help        Help about any command
    
    Flags:
          --config string   config file (default is $HOME/.nukibridgectl.yaml)
      -h, --help            help for nukibridgectl
    
    Use "nukibridgectl [command] --help" for more information about a command.

# Configuration

## Config file

For ease of use I would recommend the use of a settings-file named 
`.nukibridgectl.yaml` in one of the subsequent locations:

- `/etc/nukibridgectl/`
- Your home directory
- Working dir

The file should look like this:

    host: "192.168.178.1:8080"
    token: abcde6

To use a custom name or location for your configuration file:

    $ nukibridgectl --config /my/config/feelfree.yaml ...

## Params

If you want to pass on your configuration at the command-line you can do so, too:

    $ nukibridgectl bridge --host 192.168.178.1:8080 --token abcde6 <command>

# Examples

## List devices

After you set up your configuration you can query the bridge for available devices: 

    $ nukibridgectl bridge list
    Type ID        Name        Battery Firmware Version
    0    123456789 Wohnungstür 68%     2.9.10
    2    123456789 Haustür     0%      1.6.4

A type of 0 means a Smartlock, a type of 2 an Opener.

## Simple Locking / Unlocking of a Nuki Smartlock

To just lock/unlock a smart lock (type: 0) you type:

    $ nukibridgectl bridge lock 0 <deviceID>
    $ nukibridgectl bridge unlock 0 <deviceID>

## Applying more elaborate lock actions

See [nuki Bridge's API docs on lock action](https://developer.nuki.io/page/nuki-bridge-http-api-1-12/4#heading--lock-actions) for 
a list of available actions and their purpose. 

    # Execute actions at a smart lock (type: 0) via lockAction
    $ nukibridgectl bridge lockAction 0 <deviceID> <lock|unlock|unlatch|lockandgo|lockandgowithunlatch>

    # Execute actions at an opener (type: 2) via lockAction
    $ nukibridgectl bridge lockAction 2 <deviceID> <rto_on|rto_off|esa|cm_on|cm_off>

## Query the state

See [nuki Bridge's API docs on lock states](https://developer.nuki.io/page/nuki-bridge-http-api-1-12/4#heading--lock-states) for 
further details on the available information.

Short reminder from the Nuki docs:

> Warning: /lockstate gets the current state directly from the device and so should not be used for constant polling to avoid draining the batteries too fast. /list can be used to get regular updates on the state, as is it cached on the bridge.

    $ ./nukibridgectl bridge lockState 2 123456789

# Tests

## Integration tests

To run the tests on a physical bridge at your LAN:

    % make integration-test BRIDGE_HOST=<ip:port> BRIDGE_TOKEN=<token>

# Debugging things

You noticed some unexpected behaviour or just want to know whats going on
behind the scenes? You can enable debug logging by setting DEBUG at your env.

    $ DEBUG=1 nukibridgectl bridge --host 192.168.178.1:8080 --token abcde6 info

# DBus Bridge

`nukibridgectl` can also forward all events of your Nuki bridge to a DBus 
instance. This is useful if you want to open up your Nuki events to a broader 
audience, e.g. 3rd party applications.

Therefor `nukibridgectl` will launch a local webserver and registers itself for
callbacks.

## Fire up a local DBus instance for testing

    $ MY_SESSION_BUS_SOCKET=/tmp/dbus/$USER.session.usock
    $ dbus-daemon --session --nofork --address unix:path=$MY_SESSION_BUS_SOCKET

## Run nukibridgectl

    $ DBUS_SESSION_BUS_ADDRESS=unix:path=/tmp/dbus/$USER.session.usock DEBUG=1 go run cmd/nukibridgectl/main.go bridge watch en0

    2020-11-21T15:43:58+01:00 DBG   bridge.CallbackData{
            ... // 5 identical fields
            BatteryCritical:       false,
            KeypadBatteryCritical: false,
    -       DoorsensorState:       3,
    +       DoorsensorState:       2,
    -       DoorsensorStateName:   "door opened",
    +       DoorsensorStateName:   "door closed",
            RingactionTimestamp:   s"0001-01-01 00:00:00 +0000 UTC",
            RingactionState:       false,
      }
    
    2020-11-21T15:44:04+01:00 DBG   bridge.CallbackData{
            ... // 5 identical fields
            BatteryCritical:       false,
            KeypadBatteryCritical: false,
    -       DoorsensorState:       2,
    +       DoorsensorState:       3,
    -       DoorsensorStateName:   "door closed",
    +       DoorsensorStateName:   "door opened",
            RingactionTimestamp:   s"0001-01-01 00:00:00 +0000 UTC",
            RingactionState:       false,
      }

## Monitoring

    $ dbus-monitor --address unix:path=/tmp/$USER.session.usock
    
    signal time=1605969838.072019 sender=:1.0 -> destination=(null destination) serial=7 path=/nuki/bridge; interface=nuki.bridge; member=Event
       struct {
          int32 597123456
          int32 0
          int32 2
          int32 3
          string "unlocked"
          boolean false
          boolean false
          int32 3
          string "door opened"
          boolean false
       }
    
    signal time=1605969844.810805 sender=:1.0 -> destination=(null destination) serial=8 path=/nuki/bridge; interface=nuki.bridge; member=Event
       struct {
          int32 597123456
          int32 0
          int32 2
          int32 3
          string "unlocked"
          boolean false
          boolean false
          int32 2
          string "door closed"
          boolean false
       }
    
    
    
