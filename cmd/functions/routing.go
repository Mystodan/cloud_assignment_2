package server

import (
	Apis "assignment-2/constants"
	"assignment-2/endpoints/cases"
	"assignment-2/endpoints/notifications"
	"assignment-2/endpoints/policy"
	"assignment-2/endpoints/status"
	stub "assignment-2/endpoints/stubbing"
	"net/http"
	"strings"
)

func routeHandler(constant string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(strings.TrimSuffix(constant, "/"), handler)
	http.HandleFunc(constant, handler)
}
func RouteNotifications() {
	routeHandler(Apis.NOTIFICATIONS_PATH, notifications.HandlerNotifications)
	notifications.Url = Apis.NOTIFICATIONS_PATH
}
func RouteCases() {
	routeHandler(Apis.CASES_PATH, cases.HandlerCases)
	cases.Url = Apis.CASES_PATH
}
func RoutePolicy() {
	routeHandler(Apis.POLICY_PATH, policy.HandlerPolicy)
	policy.Url = Apis.POLICY_PATH
}
func RouteStatus() {
	routeHandler(Apis.STATUS_PATH, status.HandlerStatus)
	status.Url = Apis.STATUS_PATH
}

//Deprecated
func RouteStubbing() {
	routeHandler(Apis.STUBBING, stub.HandlerStub)
	stub.Url = Apis.STUBBING
}
