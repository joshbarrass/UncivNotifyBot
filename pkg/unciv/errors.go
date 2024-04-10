package unciv

import "errors"

var ErrBadStatusCode = errors.New("response has bad status code")
var ErrCouldNotFindPlayer = errors.New("could not find player")
