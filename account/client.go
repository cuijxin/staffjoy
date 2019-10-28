package account

// NewClient returns a gRPC client for interacting with the account.
// After calling it, run a defer close on the close function
func NewClient() (AccountServiceClient, func() error, error) {

}
