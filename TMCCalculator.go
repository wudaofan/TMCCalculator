package main

import (
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	fmt.Print("Start TMC Calculation!")
	//define variable
	orderNum := "MU5354123456"
	flightNum := "MU5354"
	departureCity := "美国"
	arrivalCity := "武汉"
	flightPrice := "1070"
	username := "解华朋"
	hongbaoNum := "5.88"
	phoneNum := "15558036555"

	part01 := "INSERT INTO `hd_hongbao`.`orders` (`id`, `user_name`, `id_card_number`, `phone_number`, `provider_id`, `service_id`, `order_number`, `order_detail`, `hongbao_given`, `hongbao_used`, `created`, `modified`) VALUES (NULL, '"
	part02 := "', '111', '"
	part03 := "', '1', '1', '"
	part04 := "', '{\\\"flight_number\\\" : \\\""
	part05 := "\\\", \\\"flight_date_departure\\\" : \\\"2017-09-06 08:57\\\", \\\"flight_date_arrival\\\" : \\\"2017-09-06 11:57\\\", \\\"flight_city_departure\\\" : \\\""
	part06 := "\\\", \\\"flight_city_arrival\\\" : \\\""
	part07 := "\\\",\\\"flight_class\\\" : \\\"经济舱\\\", \\\"flight_price\\\" : \\\""
	part08 := "\\\", \\\"flight_tax\\\" : \\\"0\\\"}', '"
	part09 := "', '0', '2017-09-05 14:55:50', '2017-09-05 14:55:50');"

	finalStr := part01 + username + part02 + phoneNum + part03 + orderNum + part04 + flightNum + part05 + departureCity + part06 + arrivalCity + part07 + flightPrice + part08 + hongbaoNum + part09

	//generate orders 订单记录
	d1 := []byte(finalStr)
	err := ioutil.WriteFile("orders.txt", d1, 0644)
	check(err)

	//generate hongbaos 红包发放记录
	d1 = []byte("hongbaos 红包发放记录")
	err = ioutil.WriteFile("hongbaos.txt", d1, 0644)
	check(err)

	//generate hongbao_statistics 红包总数
	d1 = []byte("hongbao_statistics 红包总数")
	err = ioutil.WriteFile("hongbao_statistics.txt", d1, 0644)
	check(err)
}
