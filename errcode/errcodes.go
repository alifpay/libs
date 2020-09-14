package errcode

// Error codes
const (
	TransactionNotExists          = 104
	TransactionAlreadyExists      = 107
	SuccessDuplicate              = 108
	Success                       = 200
	InProcessing                  = 201
	AcceptedToProcessing          = 202
	Rejected                      = 203
	CheckRequestSuccessful        = 301
	AccountFound                  = 302
	ServiceTemporarilyUnavailable = 303
	DoNotTryAgain                 = 304
	InsufficientFundsOnTheAccount = 305
	BadRequest                    = 400
	NotFound                      = 404
	AmountOutOfRange              = 405
	CheckRequestFailed            = 406
	InternalServerError           = 500
	ServiceUnavailable            = 503
	UnknownError                  = 520
	ProviderConnectionFailed      = 521
	BadProviderResponse           = 523
)
