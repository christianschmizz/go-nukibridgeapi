# go-nukibridgeapi

This project aims to make Nuki's Bridge API accessible from the command-line.

Therefor it is introducing a command-line tool called `nukibridgectl` which 
implements access to the basic functions of your Nuki bridge for now.
 
# Activate the Bridge API

Before you can access the bridge's API you have to activate it. *It's not active
by default.*

## Check API status

If you are unsure about whether the bridge's API is already activated or not 
you can check with curl:

    $ curl http://<bridge-ip>
    
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

## Params

If you want to pass on your configuration at the command-line you can do so, too:

    $ nukibridgectl bridge --host 192.168.178.1:8080 --token <command>
