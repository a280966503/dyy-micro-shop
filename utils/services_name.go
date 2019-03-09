package utils

const (
	SERVICE_MANAGER_WEB 		="com.dyy.manager.web"
	SERVICE_MANAGER_SERVICE 	="com.dyy.manager.service"
)

var servicePort = map[string]string{
	SERVICE_MANAGER_WEB		:			":8081",
}

func ServicePort(name string)string  {
	s := servicePort[name]

	return s

}