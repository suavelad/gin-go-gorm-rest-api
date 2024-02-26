package config

import (
	"os"
)

var PORT string = os.Getenv("PORT")
var DOMAIN_HOST string = os.Getenv("DOMAIN_HOST")
