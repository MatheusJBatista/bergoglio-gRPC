package generateCode

import "fmt"

type PromotionalCodeSerialAndOrder struct {
	SerialNumberInit  int32
	SerialNumberFinal int32
	QuantityPerSerie  int32
}

func PadLeft(str, pad string, lenght int) string {
	for {
		str = pad + str
		if len(str) > lenght {
			return str[0:lenght]
		}
	}
}

func Generate(request *PromotionalCodeSerialAndOrder) []string {
	promotionalCodes := []string{}
	lenghtSerial := len(fmt.Sprint(request.SerialNumberFinal))
	lenghtOrder := len(fmt.Sprint(request.QuantityPerSerie)) - 1
	for serial := request.SerialNumberInit; serial <= request.SerialNumberFinal; serial++ {
		for order := 0; order < int(request.QuantityPerSerie); order++ {
			code := PadLeft(fmt.Sprint(serial), "0", lenghtSerial) + PadLeft(fmt.Sprint(order), "0", lenghtOrder)
			promotionalCodes = append(promotionalCodes, code)
		}
	}
	return promotionalCodes
}
