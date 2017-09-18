package main

import (
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	fmt.Println("Start TMC Calculation!")

	//define variable
	userID := "41"                      //用户ID
	hongbaoNum := "5.88"                //随机红包数量
	phoneNum := "18605710572"           //代理商手机号
	createDate := "2017-09-13 20:00:00" //创建时间
	orderNum := "AZ9253"                //订单号
	flightNum := "AZ9253"               //航班号

	//departureCity := "上海"
	departureCitys := []string{
		"海口",
		"杭州",
		"厦门",
		"扬州",
		"成都",
		"成都",
		"杭州",
		"杭州",
		"杭州",
		"杭州",
		"太原",
		"太原",
		"南京",
		"南京",
		"杭州",
		"杭州",
	}

	//arrivalCity := "西安"
	arrivalCitys := []string{
		"南昌 ",
		"北京",
		"浦东",
		"深圳",
		"杭州",
		"杭州",
		"拉萨",
		"拉萨",
		"拉萨",
		"石家庄",
		"杭州",
		"杭州",
		"长春",
		"长春",
		"深圳",
		"深圳",
	}

	//flightPrice := "1610"
	flightPrices := []string{
		"630",
		"1520",
		"670",
		"1180",
		"1550",
		"1550",
		"1920",
		"1920",
		"1920",
		"360",
		"1350",
		"1350",
		"1290",
		"1290",
		"1170",
		"1170",
	}

	//username := "赵楠"
	userNames := []string{
		"周国兰",
		"刘小娇",
		"熊德煌",
		"贡旭东",
		"毛丹",
		"王璐",
		"卢巧文",
		"高爱华",
		"王桂芬",
		"金燕和",
		"杨冬英",
		"陈家芳",
		"潘菘芋",
		"宣槐",
		"叶遇春",
		"叶伟生",
	}
	//fmt.Println(userNames)

	fmt.Println("length = " + strconv.Itoa(len(userNames)))

	//generate orders 订单记录
	/*
		d1 := []byte(finalStr)
		err := ioutil.WriteFile("orders.txt", d1, 0644)
		check(err)
	*/

	for i := 0; i < len(userNames); i++ {

		//fmt.Print(userNames[i] + ",")
		part01 := "INSERT INTO `hd_hongbao`.`orders` (`id`, `user_name`, `id_card_number`, `phone_number`, `provider_id`, `service_id`, `order_number`, `order_detail`, `hongbao_given`, `hongbao_used`, `created`, `modified`) VALUES (NULL, '"
		part02 := "', '111', '"
		part03 := "', '1', '1', '"
		part04 := "', '{\\\"flight_number\\\" : \\\""
		part05 := "\\\", \\\"flight_date_departure\\\" : \\\"2017-09-06 08:57\\\", \\\"flight_date_arrival\\\" : \\\"2017-09-06 11:57\\\", \\\"flight_city_departure\\\" : \\\""
		part06 := "\\\", \\\"flight_city_arrival\\\" : \\\""
		part07 := "\\\",\\\"flight_class\\\" : \\\"经济舱\\\", \\\"flight_price\\\" : \\\""
		part08 := "\\\", \\\"flight_tax\\\" : \\\"0\\\"}', '"
		part09 := "', '0', '" + createDate + "', '" + createDate + "');"

		finalStr := part01 + userNames[i] + part02 + phoneNum + part03 + orderNum + part04 + flightNum + part05 + departureCitys[i] + part06 + arrivalCitys[i] + part07 + flightPrices[i] + part08 + hongbaoNum + part09

		f, err := os.OpenFile("orders.txt", os.O_APPEND, 0666)
		n, err := f.WriteString(finalStr)
		f.Close()
		check(err)
		print(strconv.Itoa(n) + ",")

	}

	//generate hongbaos 红包发放记录
	/*
		d1 = []byte("hongbaos 红包发放记录")
		err = ioutil.WriteFile("hongbaos.txt", d1, 0644)
		check(err)
	*/

	//随机红包
	/* */
	for i := 0; i < len(userNames); i++ {
		hongbaosSuijiStr := "INSERT INTO `hd_hongbao`.`hongbaos` (`id`, `user_id`, `service_id`, `provider_id`, `type`, `hongbao_count`, `hongbao_money`, `description`, `created`, `modified`) VALUES (NULL, '" + userID + "', '1', '1', '1', '" + hongbaoNum + "', '" + hongbaoNum + "','2', '" + createDate + "', '" + createDate + "');"

		f, err := os.OpenFile("suijihongbaos.txt", os.O_APPEND, 0666)
		n, err := f.WriteString(hongbaosSuijiStr)
		f.Close()
		check(err)
		print(strconv.Itoa(n) + ",")
	}

	//平台补贴红包
	/**/
	compensateHongbaos := []string{
		"42.5",
		"105",
		"45",
		"80",
		"105",
		"105",
		"130",
		"130",
		"130",
		"25",
		"92.5",
		"92.5",
		"87.5",
		"87.5",
		"80",
		"80",
	}

	for i := 0; i < len(compensateHongbaos); i++ {
		hongbaosStr := "INSERT INTO `hd_hongbao`.`hongbaos` (`id`, `user_id`, `service_id`, `provider_id`, `type`, `hongbao_count`, `hongbao_money`, `description`, `created`, `modified`) VALUES (NULL, '" + userID + "', '1', '1', '1', '" + compensateHongbaos[i] + "', '" + compensateHongbaos[i] + "','4', '" + createDate + "', '" + createDate + "');"

		f, err := os.OpenFile("hongbaos.txt", os.O_APPEND, 0666)
		n, err := f.WriteString(hongbaosStr)
		f.Close()
		check(err)
		print(strconv.Itoa(n) + ",")
	}

	//generate hongbao_statistics 红包总数
	/*
		d1 = []byte("hongbao_statistics 红包总数")
		err = ioutil.WriteFile("hongbao_statistics.txt", d1, 0644)
		check(err)
	*/
}
