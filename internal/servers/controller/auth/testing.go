package auth

import (
	"context"

	authpb "github.com/hashicorp/boundary/internal/gen/controller/auth"
	"github.com/hashicorp/boundary/internal/requests"
	"github.com/hashicorp/boundary/internal/servers/controller/common"
)

// DisabledAuthTestContext is meant for testing, and uses a context that has
// auth checking entirely disabled. Supported options: WithScopeId an WithUserId
// are used directly; WithKms is passed through into the verifier context.
func DisabledAuthTestContext(iamRepoFn common.IamRepoFactory, scopeId string, opt ...Option) context.Context {
	reqInfo := authpb.RequestInfo{DisableAuthEntirely: true}
	opts := getOpts(opt...)
	reqInfo.ScopeIdOverride = opts.withScopeId
	if reqInfo.ScopeIdOverride == "" {
		reqInfo.ScopeIdOverride = scopeId
	}
	reqInfo.UserIdOverride = opts.withUserId
	if reqInfo.UserIdOverride == "" {
		reqInfo.UserIdOverride = "u_auth"
	}
	requestContext := context.WithValue(context.Background(), requests.ContextRequestInformationKey, &requests.RequestContext{})
	return NewVerifierContext(requestContext, iamRepoFn, nil, nil, opts.withKms, &reqInfo)
}
