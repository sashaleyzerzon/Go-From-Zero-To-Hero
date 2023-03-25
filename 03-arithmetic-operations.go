package main

func main03() {
	// Execrise 01
	println("Exersice 01")

	var nPackages int = 100
	println("I have", nPackages, "packages to deliver")

	// Execrcise 02
	var nDeliveredPackages int = 20
	println("I have", nPackages-nDeliveredPackages, "left to deliver")

	// Exercise 03
	var nCustomerNum int = 4
	println("I have", nCustomerNum, "to spread the deliveries between. \nEach customer will receive", (nPackages-nDeliveredPackages)/nCustomerNum, "packages")

}
