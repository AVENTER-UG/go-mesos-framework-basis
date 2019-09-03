# go-mesos-framework-basis

Dies ist die Basis für Mesos Frameworks.

## Framework starten

```Bash

export FRAMEWORK_USER="root"
export FRAMEWORK_NAME="test_framework"
export MESOS_PRINCIPAL="<mesos_principal>"
export MESOS_USERNAME="<mesos_user>"
export MESOS_PASSWORD="<mesos_password>"
export MESOS_MASTER="<mesos_master_server>:5050"


go run init.go app.go
```

Dies startet das Framework. Es wird sich an den Mesos Master anmelden. Nach wenigen Sekunden kann man "test_framework" als Eintrag in der Mesos UI sehen. Gleichzeitig öffnet das Framework einen Port auf 10000 auf der Maschine auf dem das Framework gestartet wurde.

## Task Starten

Das Basis Framework kann zur Demonstation einen Task starten. Dies erfolgt über folgenden Aufruf.

```Bash
curl -X POST 127.0.0.1:10000/test/\?cmd\=python%20-m%20SimpleHTTPServer%209033
```

Auf einem Mesos Agent wird man nun einen entsprechenden Prozess erkennen können.
