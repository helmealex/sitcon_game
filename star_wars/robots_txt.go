package star_wars

import (
	"fmt"
	"net/http"
)

// robotsTxtHandler serves a virtual robots.txt file.
// It disallows all web crawlers from indexing the site.
func RobotsTxtHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	fmt.Fprintln(w, "User-agent: the explorer")
	fmt.Fprintln(w, "Disallow: /")
	fmt.Fprintln(w, "Allow: L2hvZ3dhcnRz")
}
