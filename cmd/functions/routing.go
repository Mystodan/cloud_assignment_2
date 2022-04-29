package server

import (
	consts "assignment-2/constants"
	"assignment-2/endpoints/cases"
	"assignment-2/endpoints/notifications"
	"assignment-2/endpoints/policy"
	"assignment-2/endpoints/status"
	"net/http"
	"strings"
)

func routeHandler(constant string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(strings.TrimSuffix(constant, "/"), handler)
	http.HandleFunc(constant, handler)
}
func RouteNotifications() {
	routeHandler(consts.NOTIFICATIONS_PATH, notifications.HandlerNotifications)
	notifications.Url = consts.NOTIFICATIONS_PATH
}
func RouteCases() {
	routeHandler(consts.CASES_PATH, cases.HandlerCases)
	cases.Url = consts.CASES_PATH
}
func RoutePolicy() {
	routeHandler(consts.POLICY_PATH, policy.HandlerPolicy)
	policy.Url = consts.POLICY_PATH
}
func RouteStatus() {
	routeHandler(consts.STATUS_PATH, status.HandlerStatus)
	status.Url = consts.STATUS_PATH
}
