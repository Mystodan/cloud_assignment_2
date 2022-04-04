package main

import (
	consts "assignment-2/constants"
	funcs "assignment-2/functions"
)

/**
 *	The main function.
 *	Creates a link/routes to the correct handlers
 */

func main() {
	funcs.TimerStart(funcs.Timer)
	port := funcs.SetPort(consts.DEFAULT_PORT)

	funcs.SetListener(port)
}
