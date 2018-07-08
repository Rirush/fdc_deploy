# FDC Deploy Protocol commands

## PING
Checks connection with server.
### Arguments:
None.
### Returns:
None.

## ERROR
Used by server. Indicates that an error happened while server was processing a command.
### Arguments:
- **1:** Error code
- **2:** Error description
### Returns:
N/A

## LIST\_APPS
Requests all available applications to current user.
### Arguments:
None.
### Returns:
- **Repeated:**
- - **1:** Application name

## DEPLOY\_NEW
Creates and deploys new application on server.
### Arguments:
- **1:** Unique application name
- **2:** Application repository URL (must contain `fdc.yaml` file in its root and must be reachable from server (ssh key added on repository host, or HTTP(s) with no authorization))
### Returns:
- **1:** Deployed application version

## DEPLOY\_UPDATE
Deploys new version of existing on server application. Latest version on repository will be deployed.
### Arguments:
- **1:** Application name
### Returns:
- **1:** Deployed application version

## DEPLOY\_ROLLBACK
Rolls back to previous deploy. Server must keep previous deploy.
### Arguments:
- **1:** Application name
### Returns:
- **1:** Deployed application version

## DEPLOY\_DELETE
Deletes all containers and saved code from server.
### Arguments:
- **1:** Application name
### Returns:
None.

## APP\_VERSION
Requests selected app version.
### Arguments:
- **1:** Application name
### Returns:
- **1:** Application version

## APP\_STOP
Stops all containers of specified application.
### Arguments:
- **1:** Application name
### Returns:
None.

## APP\_START
Starts all containers of specified application.
### Arguments:
- **1:** Application name
### Returns:
None.

## APP\_RESTART
Restarts all containers of specified application.
### Arguments:
- **1:** Application name
### Returns:
None.

## APP\_STATUS
Requests application containers status.
### Arguments:
- **1:** Application name
### Returns:
- **1:** Application status (one of `running` or `stopped`)
