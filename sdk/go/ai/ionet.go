package ai

import "strings"

// ionetModelPrefix marks a model as IO Intelligence (io.net)-routed. It is
// stripped before the request goes out; see stripIonetPrefix.
const ionetModelPrefix = "ionet/"

// stripIonetPrefix removes the routing-only "ionet/" prefix from a model
// name. IO Intelligence serves bare Hugging Face-style `org/name` ids, so
// the marker must not reach the wire.
func stripIonetPrefix(model string) string {
	if len(model) >= len(ionetModelPrefix) &&
		strings.EqualFold(model[:len(ionetModelPrefix)], ionetModelPrefix) {
		return model[len(ionetModelPrefix):]
	}
	return model
}
