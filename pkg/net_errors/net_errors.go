package net_errors

func IoError() string { return "i/o timeout" }
func BindSocketError() string {
	return "bind: An operation on a socket could not be performed because the system lacked sufficient buffer space or because a queue was full."
}
