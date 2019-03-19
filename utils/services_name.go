package utils

const (
	SERVICE_MANAGER_WEB 		="com.dyy.manager.web"
	SERVICE_MANAGER_SERVICE 	="com.dyy.manager.service"

	SERVICE_SHOP_WEB 		="com.dyy.shop.web"
	SERVICE_SHOP_SERVICE 	="com.dyy.shop.service"

	SERVICE_PORTAL_WEB 		="com.dyy.portal.web"

	SERVICE_SEARCH_WEB 		="com.dyy.search.web"

	SERVICE_USER_WEB 		="com.dyy.user.web"
	SERVICE_CART_WEB 		="com.dyy.cart.web"
	SERVICE_CONTENT_SERVICE		="com.dyy.content.service"

	)

var servicePort = map[string]string{
	SERVICE_MANAGER_WEB		:			":8081",
	SERVICE_SHOP_WEB		:			":8082",
	SERVICE_PORTAL_WEB		:			":8083",
	SERVICE_SEARCH_WEB		:			":8084",
	SERVICE_USER_WEB		:			":8085",
	SERVICE_CART_WEB		:			":8086",
}

func ServicePort(name string)string  {
	s := servicePort[name]

	return s

}