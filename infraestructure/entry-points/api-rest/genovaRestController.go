package apirest

import (
	"github.com/jearcila/hex-architecture/core/ports"
	useCase "github.com/jearcila/hex-architecture/core/useCase"
)

// Init genova service
type GenovaRestController struct{}

func CreateInstance() GenovaRestController {
	return GenovaRestController{}
}

func (g GenovaRestController) AuthorizationRest(acquirer ports.GenovaServiceInt) useCase.G2Handler {
	return useCase.Authorization(acquirer)
}

func (g GenovaRestController) CaptureRest(acquirer ports.GenovaServiceInt) useCase.G2Handler {
	return useCase.Capture(acquirer)
}

func (g GenovaRestController) PurchaseRest(acquirer ports.GenovaServiceInt) useCase.G2Handler {
	return useCase.Purchase(acquirer)
}

func (g GenovaRestController) RefundRest(acquirer ports.GenovaServiceInt) useCase.G2Handler {
	return useCase.Refund(acquirer)
}
