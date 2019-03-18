package utils

const (
	SERVICE_MANAGER_WEB 		="com.dyy.manager.web"
	SERVICE_MANAGER_SERVICE 	="com.dyy.manager.service"

	SERVICE_SHOP_WEB 		="com.dyy.shop.web"
	SERVICE_SHOP_SERVICE 	="com.dyy.shop.service"

	SERVICE_PORTAL_WEB 		="com.dyy.portal.web"
	SERVICE_PORTAL_SERVICE 	="com.dyy.portal.service"
)

var servicePort = map[string]string{
	SERVICE_MANAGER_WEB		:			":8081",
	SERVICE_SHOP_WEB		:			":8082",
	SERVICE_PORTAL_WEB		:			":8083",
}

func ServicePort(name string)string  {
	s := servicePort[name]

	return s

}