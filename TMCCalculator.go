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

	fmt.Print("Start TMC Calculation!")
	//define variable
	userID := "37"
	orderNum := "MU9578"
	flightNum := "MU9578"
	//departureCity := "上海"
	departureCitys := []string{
		"温州",
		"鄂尔多斯",
		"温州",
		"杭州",
		"广州",
		"成都",
		"厦门",
		"厦门",
		"广州",
		"广州",
		"广州",
		"西安",
		"西安",
		"西安",
		"西安",
		"广州",
		"厦门",
		"果洛",
		"广州",
		"满洲里",
		"丽江",
		"丽江",
		"浦东",
		"广州",
		"成都",
		"海口",
	}

	//arrivalCity := "西安"
	arrivalCitys := []string{
		"西安",
		"上海",
		"西安",
		"赤峰",
		"南通",
		"北京",
		"郑州",
		"郑州",
		"南通",
		"济南",
		"济南",
		"杭州",
		"杭州",
		"杭州",
		"杭州",
		"上海",
		"广州",
		"西宁",
		"青岛",
		"武汉",
		"杭州",
		"杭州",
		"成都",
		"太原",
		"南京",
		"北京",
	}

	//flightPrice := "1610"
	flightPrices := []string{
		"1160",
		"1910",
		"1160",
		"2230",
		"2320",
		"1210",
		"1770",
		"1170",
		"2320",
		"2310",
		"2310",
		"1640",
		"1640",
		"1640",
		"1640",
		"2440",
		"1710",
		"1610",
		"1900",
		"2950",
		"1220",
		"1610",
		"1900",
		"1750",
		"1530",
		"2570",
	}

	//username := "赵楠"
	hongbaoNum := "5.88"
	phoneNum := "15657178960"

	userNames := []string{
		"潘萍萍",
		"吴巧玲",
		"蔡玲玲",
		"江亦爽",
		"唐国裕",
		"敬子洁",
		"赵锴豪",
		"赵锴豪",
		"梁洁宁",
		"焦亚西",
		"汪海燕",
		"张春雷",
		"谢子瑞",
		"张子晗",
		"唐子坤",
		"张家宁",
		"龚旭东",
		"洪梦权",
		"谢荣祥",
		"姚可海",
		"胡湘丹",
		"胡湘丹",
		"陈念宇",
		"吴文斌",
		"孙帅",
		"张立庆",
	}
	//fmt.Println(userNames)

	fmt.Println("length = " + strconv.Itoa(len(userNames)))

	//订单记录
	//sum := 0
	for i := 0; i < len(userNames); i++ {
		//sum += i
		//fmt.Print(userNames[i] + ",")

		part01 := "INSERT INTO `hd_hongbao`.`orders` (`id`, `user_name`, `id_card_number`, `phone_number`, `provider_id`, `service_id`, `order_number`, `order_detail`, `hongbao_given`, `hongbao_used`, `created`, `modified`) VALUES (NULL, '"
		part02 := "', '111', '"
		part03 := "', '1', '1', '"
		part04 := "', '{\\\"flight_number\\\" : \\\""
		part05 := "\\\", \\\"flight_date_departure\\\" : \\\"2017-09-06 08:57\\\", \\\"flight_date_arrival\\\" : \\\"2017-09-06 11:57\\\", \\\"flight_city_departure\\\" : \\\""
		part06 := "\\\", \\\"flight_city_arrival\\\" : \\\""
		part07 := "\\\",\\\"flight_class\\\" : \\\"经济舱\\\", \\\"flight_price\\\" : \\\""
		part08 := "\\\", \\\"flight_tax\\\" : \\\"0\\\"}', '"
		part09 := "', '0', '2017-09-13 09:00:00', '2017-09-13 09:00:00');"

		finalStr := part01 + userNames[i] + part02 + phoneNum + part03 + orderNum + part04 + flightNum + part05 + departureCitys[i] + part06 + arrivalCitys[i] + part07 + flightPrices[i] + part08 + hongbaoNum + part09

		//generate orders 订单记录
		/*
			d1 := []byte(finalStr)
			err := ioutil.WriteFile("orders.txt", d1, 0644)
			check(err)
		*/
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
	/*
		for i := 0; i < len(userNames); i++ {
			hongbaosStr := "INSERT INTO `hd_hongbao`.`hongbaos` (`id`, `user_id`, `service_id`, `provider_id`, `type`, `hongbao_count`, `hongbao_money`, `description`, `created`, `modified`) VALUES (NULL, '" + userID + "', '1', '1', '1', '5.88', '5.88','2', '2017-09-13 11:00:00', '2017-09-13 11:00:00');"

			f, err := os.OpenFile("hongbaos.txt", os.O_APPEND, 0666)
			n, err := f.WriteString(hongbaosStr)
			f.Close()
			check(err)
			print(strconv.Itoa(n) + ",")
		}
	*/

	//平台补贴红包
	/* */
	compensateHongbaos := []string{
		"80",
		"130",
		"80",
		"152.5",
		"160",
		"82.5",
		"120",
		"120",
		"160",
		"365",
		"365",
		"260",
		"260",
		"260",
		"260",
		"167.5",
		"117.5",
		"110",
		"130",
		"202.5",
		"110",
		"110",
		"130",
		"120",
		"105",
		"175",
	}

	for i := 0; i < len(compensateHongbaos); i++ {
		hongbaosStr := "INSERT INTO `hd_hongbao`.`hongbaos` (`id`, `user_id`, `service_id`, `provider_id`, `type`, `hongbao_count`, `hongbao_money`, `description`, `created`, `modified`) VALUES (NULL, '" + userID + "', '1', '1', '1', '" + compensateHongbaos[i] + "', '" + compensateHongbaos[i] + "','4', '2017-09-13 11:00:00', '2017-09-13 11:00:00');"

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
