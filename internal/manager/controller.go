package manager

import (
	controllers "smartparking/internal/controller"
	"smartparking/internal/interfaces/controller"
	"sync"
)

type controllerImpl struct {
	authControllerInit         sync.Once
	authController             controller.Auth
	carControllerInit          sync.Once
	carController              controller.Car
	clientControllerInit       sync.Once
	clientController           controller.Client
	coreControllerInit         sync.Once
	coreController             controller.Core
	entryHistoryControllerInit sync.Once
	entryHistoryController     controller.EntryHistory
	imageControllerInit        sync.Once
	imageController            controller.Image
	parkingPlaceControllerInit sync.Once
	parkingPlaceController     controller.ParkingPlace
	parkingZoneControllerInit  sync.Once
	parkingZoneController      controller.ParkingZone
	personalControllerInit     sync.Once
	personalController         controller.Personal
	taxControllerInit          sync.Once
	taxController              controller.Tax
	ticketControllerInit       sync.Once
	ticketController           controller.Ticket
}

func (cm *controllerImpl) Auth() controller.Auth {
	cm.authControllerInit.Do(func() {
		if cm.authController == nil {
			cm.authController = controllers.NewAuthController(MManager)
		}
	})
	return cm.authController
}

func (cm *controllerImpl) Car() controller.Car {
	cm.carControllerInit.Do(func() {
		if cm.carController == nil {
			cm.carController = controllers.NewCarController(MManager)
		}
	})
	return cm.carController
}

func (cm *controllerImpl) Client() controller.Client {
	cm.clientControllerInit.Do(func() {
		if cm.clientController == nil {
			cm.clientController = controllers.NewClientController(MManager)
		}
	})
	return cm.clientController
}

func (cm *controllerImpl) Core() controller.Core {
	cm.coreControllerInit.Do(func() {
		if cm.coreController == nil {
			cm.coreController = controllers.NewCoreController(MManager)
		}
	})
	return cm.coreController
}

func (cm *controllerImpl) EntryHistory() controller.EntryHistory {
	cm.entryHistoryControllerInit.Do(func() {
		if cm.entryHistoryController == nil {
			cm.entryHistoryController = controllers.NewEntryHistoryController(MManager)
		}
	})
	return cm.entryHistoryController
}

func (cm *controllerImpl) Image() controller.Image {
	cm.imageControllerInit.Do(func() {
		if cm.imageController == nil {
			cm.imageController = controllers.NewImageController(MManager)
		}
	})
	return cm.imageController
}

func (cm *controllerImpl) ParkingPlace() controller.ParkingPlace {
	cm.parkingPlaceControllerInit.Do(func() {
		if cm.parkingPlaceController == nil {
			cm.parkingPlaceController = controllers.NewParkingPlaceController(MManager)
		}
	})
	return cm.parkingPlaceController
}

func (cm *controllerImpl) ParkingZone() controller.ParkingZone {
	cm.parkingZoneControllerInit.Do(func() {
		if cm.parkingZoneController == nil {
			cm.parkingZoneController = controllers.NewParkingZoneController(MManager)
		}
	})
	return cm.parkingZoneController
}

func (cm *controllerImpl) Personal() controller.Personal {
	cm.personalControllerInit.Do(func() {
		if cm.personalController == nil {
			cm.personalController = controllers.NewPersonalController(MManager)
		}
	})
	return cm.personalController
}

func (cm *controllerImpl) Tax() controller.Tax {
	cm.taxControllerInit.Do(func() {
		if cm.taxController == nil {
			cm.taxController = controllers.NewTaxController(MManager)
		}
	})
	return cm.taxController
}

func (cm *controllerImpl) Ticket() controller.Ticket {
	cm.ticketControllerInit.Do(func() {
		if cm.ticketController == nil {
			cm.ticketController = controllers.NewTicketController(MManager)
		}
	})
	return cm.ticketController
}
