package fdc_protocol

func PingRequest() Package {
	return Package{
		Command:   COMMAND_PING,
		Arguments: make([]string, 0),
	}
}

func ErrorAnswer(code, desc string) Package {
	return Package{
		Command:   COMMAND_ERROR,
		Arguments: []string{code, desc},
	}
}

func ListAppsRequest() Package {
	return Package{
		Command:   COMMAND_LIST_APPS,
		Arguments: make([]string, 0),
	}
}

type Application struct {
	Name string
	URL  string
}

func ListAppsAnswer(apps ...Application) Package {
	p := Package{
		Command:   COMMAND_LIST_APPS,
		Arguments: make([]string, 0),
	}
	for _, app := range apps {
		p.Arguments = append(p.Arguments, app.Name, app.URL)
	}
	return p
}

func DeployNewRequest(name, url string) Package {
	return Package{
		Command:   COMMAND_DEPLOY_NEW,
		Arguments: []string{name, url},
	}
}

func DeployNewAnswer(version string) Package {
	return Package{
		Command:   COMMAND_DEPLOY_NEW,
		Arguments: []string{version},
	}
}

func DeployUpdateRequest(name string) Package {
	return Package{
		Command:   COMMAND_DEPLOY_UPDATE,
		Arguments: []string{name},
	}
}

func DeployUpdateAnswer(version string) Package {
	return Package{
		Command:   COMMAND_DEPLOY_UPDATE,
		Arguments: []string{version},
	}
}

func DeployRollbackRequest(name string) Package {
	return Package{
		Command:   COMMAND_DEPLOY_ROLLBACK,
		Arguments: []string{name},
	}
}

func DeployRollbackAnswer(version string) Package {
	return Package{
		Command:   COMMAND_DEPLOY_ROLLBACK,
		Arguments: []string{version},
	}
}

func DeployDeleteRequest(name string) Package {
	return Package{
		Command:   COMMAND_DEPLOY_DELETE,
		Arguments: []string{name},
	}
}

func DeployDeleteAnswer() Package {
	return Package{
		Command:   COMMAND_DEPLOY_DELETE,
		Arguments: make([]string, 0),
	}
}

func AppVersionRequest(name string) Package {
	return Package{
		Command:   COMMAND_APP_VERSION,
		Arguments: []string{name},
	}
}

func AppVersionAnswer(version string) Package {
	return Package{
		Command:   COMMAND_APP_VERSION,
		Arguments: []string{version},
	}
}

func AppStopRequest(name string) Package {
	return Package{
		Command:   COMMAND_APP_STOP,
		Arguments: []string{name},
	}
}

func AppStopAnswer() Package {
	return Package{
		Command:   COMMAND_APP_STOP,
		Arguments: make([]string, 0),
	}
}

func AppStartRequest(name string) Package {
	return Package{
		Command:   COMMAND_APP_START,
		Arguments: []string{name},
	}
}

func AppStartAnswer() Package {
	return Package{
		Command:   COMMAND_APP_START,
		Arguments: make([]string, 0),
	}
}

func AppRestartRequest(name string) Package {
	return Package{
		Command:   COMMAND_APP_RESTART,
		Arguments: []string{name},
	}
}

func AppRestartAnswer() Package {
	return Package{
		Command:   COMMAND_APP_RESTART,
		Arguments: make([]string, 0),
	}
}

func AppStatusRequest(name string) Package {
	return Package{
		Command:   COMMAND_APP_STATUS,
		Arguments: []string{name},
	}
}

func AppStatusAnswer(status string) Package {
	return Package{
		Command:   COMMAND_APP_STATUS,
		Arguments: []string{status},
	}
}
