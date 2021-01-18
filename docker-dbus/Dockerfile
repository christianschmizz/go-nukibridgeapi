FROM debian:buster

RUN apt-get update && apt-get install -y dbus

COPY session-local.conf /etc/dbus-1/session-local.conf
EXPOSE 55556

ENTRYPOINT ["dbus-daemon", "--config-file=/usr/share/dbus-1/session.conf", "--nofork", "--print-address", "--nopidfile"]
