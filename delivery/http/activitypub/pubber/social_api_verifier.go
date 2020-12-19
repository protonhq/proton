package pubber

import (
	"net/http"
	"net/url"
)

// socialAPIVerifier will verify incoming requests from clients and is meant to
// encapsulate authentication functionality by standards such as OAuth (RFC
// 6749).
type socialAPIVerifier struct {
}

// Verify will determine the authenticated user for the given request,
// returning false if verification fails. If the request is entirely
// missing the required fields in order to authenticate, this function
// must return nil and false for all values to permit attempting
// validation by HTTP Signatures. If there was an internal error
// determining the authentication of the request, it is returned.
//
// Return values are interpreted as follows:
//     (userFoo, true,  true,  <nil>) => userFoo passed authentication and is authorized
//     (<any>,   true,  false, <nil>) => a user passed authentication but failed authorization (Permission denied)
//     (<any>,   false, false, <nil>) => authentication failed: deny access (Bad request)
//     (<nil>,   false, true,  <nil>) => authentication failed: must pass HTTP Signature verification or will be Permission Denied
//     (<nil>,   true,  true,  <nil>) => "I don't care, try to validate using HTTP Signatures. If that still doesn't work, permit raw requests access anyway."
//     (<any>,   <any>, <any>, error) => an internal error occurred during validation
//
// Be very careful that the 'authenticatedUser' value is non-nil when
// returning 'authn' and 'authz' values of true, or else the library
// will use the most permissive logic instead of the most restrictive as
// outlined above.
func (o *socialAPIVerifier) Verify(r *http.Request) (authenticatedUser *url.URL, authn, authz bool, err error) {
	return nil, true, true, nil
}

// VerifyForOutbox is the same as Verify, except that the request must
// authenticate the owner of the provided outbox IRI.
//
// Return values are interpreted as follows:
//     (true,  true,   <nil>) => user for the outbox passed authentication and is authorized
//     (true,  false,  <nil>) => a user passed authentication but failed authorization for this outbox (Permission denied)
//     (false, true,   <nil>) => authentication failed: must pass HTTP Signature verification or will be Permission Denied
//     (false, false,  <nil>) => authentication failed: deny access (Bad request)
//     (<any>, <any>,  error) => an internal error occurred during validation
func (o *socialAPIVerifier) VerifyForOutbox(r *http.Request, outbox *url.URL) (authn, authz bool, err error) {
	_, authn, authz, err = o.Verify(r)
	return
}
