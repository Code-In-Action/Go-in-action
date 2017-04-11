package matchers

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/goinaction/code/chapter2/sample/search"
)
func init() {
	var matcher rssMatcher
	search.Register("rss", matcher)
}
