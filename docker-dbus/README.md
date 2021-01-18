Dockered DBus
=============

This directory contains a DBus setup which is used for testing and examining the DBus integration of ``nukibridgectl``.

    $ export DBUS_SESSION_BUS_ADDRESS=tcp:host=10.211.55.6,bind=*,port=55556,family=ipv4
    $ docker-compose up -d
    $ go run cmd/nukibridgectl/main.go bridge watch en0

