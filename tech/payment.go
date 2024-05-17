package tech

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
)

func (h *Handler) PayOrder(console *Console) {
	var key string
	fmt.Println("Введите ключ")
	fmt.Scanln(&key)

	err := h.payment.ProcessOrderPayment(context.Background(), key)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "PayOrder",
		}).Error(err.Error())
		ErrorResponse(err)
		return
	}

	fmt.Println("Заказ успешно оплачен")
}
