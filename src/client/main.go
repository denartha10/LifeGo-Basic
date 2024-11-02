package main

//TODO: fix the null pointer issues that I have that are caused by changing all the quotation services and client infos to concrete types
// They should be pointers

import (
	"fmt"

	"github.com/denartha10/lifeGo-basic/auldfellas"
	"github.com/denartha10/lifeGo-basic/broker"
	"github.com/denartha10/lifeGo-basic/core"
	"github.com/denartha10/lifeGo-basic/dodgygeezers"
	"github.com/denartha10/lifeGo-basic/girlsallowed"
)

var clients = []core.ClientInfo{
	core.NewClient("Niki Collier", core.FEMALE, 49, 1.5494, 80, false, false),
	core.NewClient("Old Geeza", core.MALE, 65, 1.6, 100, true, true),
	core.NewClient("Hannah Montana", core.FEMALE, 21, 1.78, 65, false, false),
	core.NewClient("Rem Collier", core.MALE, 49, 1.8, 120, false, true),
	core.NewClient("Jim Quinn", core.MALE, 55, 1.9, 75, true, false),
	core.NewClient("Donald Duck", core.MALE, 35, 0.45, 1.6, false, false),
}

// Display the client info nicely
func displayProfle(info core.ClientInfo) {
	fmt.Println("|=================================================================================================================|")
	fmt.Println("|=================================================================================================================|")
	fmt.Println("|                                     |                                     |                                     |")
	fmt.Println("| Name: " + fmt.Sprintf("%s", info.Name) + " | Gender: " + fmt.Sprintf("%v", (info.Gender)) + " | Age: " + fmt.Sprintf("%d", info.Age) + " |")
	fmt.Println("| Weight/Height: " + fmt.Sprintf("%fkg %fm", info.Weight, info.Height) + " | Smoker: " + fmt.Sprintf("%t", info.Smoker) + " | Medical Problems: " + fmt.Sprintf("%t", info.MedicalIssues) + " |")
	fmt.Println("|                                     |                                     |                                     |")
	fmt.Println("|=================================================================================================================|")
}

func displayQuotation(quotation core.Quotation) {
	fmt.Println("| Company: " + fmt.Sprintf("%s", quotation.Company) + " | Reference: " + fmt.Sprintf("%s", quotation.Reference) + " | Price: " + fmt.Sprintf("€%f", quotation.Price) + " |")
	fmt.Println("|=================================================================================================================|")
}

func main() {
	// register services
	core.ServiceRegistryBind(core.GIRLS_ALLOWED_SERVICE, girlsallowed.GAQService{})
	core.ServiceRegistryBind(core.AULD_FELLAS_SERVICE, auldfellas.AFQService{})
	core.ServiceRegistryBind(core.DODGY_GEEZERS_SERVICE, dodgygeezers.DGQService{})
	core.ServiceRegistryBind(core.BROKER_SERVICE, broker.LocalBrokerService{})

	// This is the starting point for the application. here we must get a reference
	// to the broker service and then invoke the getQutations method
	service := core.ServiceRegistryLookup(core.BROKER_SERVICE)
	var brokerService core.BrokerService
	if bser, ok := service.(core.BrokerService); ok {
		brokerService = bser
	}

	for _, info := range clients {
		displayProfle(info)

		for _, quotation := range brokerService.GetQuotations(info) {
			displayQuotation(quotation)
		}

		fmt.Print("\n\n")
	}
}
