package main

import (
	"time"

	"github.com/fatih/color"
)

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	ClientsChan     chan string
	Open            bool
}

func (shop *BarberShop) addBarber(barber string) {
	shop.NumberOfBarbers++
	go func() {
		isSleeping := false
		color.Yellow("Barber %s is ready to work!", barber)
		for {
			if len(shop.ClientsChan) == 0 {
				isSleeping = true
				color.Yellow("Barber %s is sleeping", barber)
			}
			client, shopOpen := <-shop.ClientsChan
			if shopOpen {
				if isSleeping {
					color.Yellow("Barber %s is waking up", barber)
					isSleeping = false
				}
				// cut hair
				shop.cutHair(barber, client)
			} else {
				shop.sendBarberhome(barber)
				return
			}
		}
	}()
}

func (shop *BarberShop) cutHair(barber, client string) {
	color.Green("Barber %s is cutting %s's hair", barber, client)
	time.Sleep(shop.HairCutDuration)
	color.Green("Barber %s is done cutting %s's hair", barber, client)
}

func (shop *BarberShop) sendBarberhome(barber string) {
	color.Yellow("Barber %s is going home", barber)
	shop.BarbersDoneChan <- true
}

func (shop *BarberShop) closeShopForDay() {
	color.Cyan("Barbershop is closing for the day")
	close(shop.ClientsChan)
	shop.Open = false

	for a := 1; a <= shop.NumberOfBarbers; a++ {
		<-shop.BarbersDoneChan
	}

	close(shop.BarbersDoneChan)
	color.Cyan("=======================================")
	color.Cyan("Barbershop is closed for the day")
}

func (shop *BarberShop) addClient(client string) {
	color.Green("%s is entering the barbershop", client)

	if shop.Open {
		select {
		case shop.ClientsChan <- client:
			color.Yellow("%s takes a seat in room", client)
		default:
			color.Red("%s is leaving because the barbershop is full", client)
		}
	} else {
		color.Red("%s is leaving because the barbershop is closed", client)
	}
}
