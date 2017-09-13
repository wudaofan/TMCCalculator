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
	userID := "37"        //用户ID
	orderNum := "CZ9253"  //订单号
	flightNum := "CZ9253" //航班号

	//departureCity := "上海"
	departureCitys := []string{
		"上海",
		"杭州",
		"郑州",
		"海口",
		"杭州",
		"海口",
		"海口",
		"海口",
		"海口",
		"西安",
		"三亚",
		"杭州",
		"杭州",
		"武汉",
		"喀什",
		"喀什",
		"喀什",
		"喀什",
		"喀什",
		"喀什",
		"海口",
		"厦门",
		"厦门",
		"大连",
		"大连",
		"西安",
		"西安",
		"西安",
		"温州",
		"温州",
		"杭州",
		"太原",
		"杭州",
		"宁波",
	}

	//arrivalCity := "西安"
	arrivalCitys := []string{
		"沈阳",
		"深圳",
		"杭州",
		"杭州",
		"海口",
		"郑州",
		"郑州",
		"郑州",
		"郑州",
		"三亚",
		"西安",
		"武汉",
		"西安",
		"三亚",
		"乌鲁木齐",
		"乌鲁木齐",
		"乌鲁木齐",
		"乌鲁木齐",
		"乌鲁木齐",
		"乌鲁木齐",
		"南昌",
		"日照",
		"日照",
		"宁波",
		"宁波",
		"杭州",
		"杭州",
		"杭州",
		"西安",
		"西安",
		"成都",
		"广州",
		"沈阳",
		"成都",
	}

	//flightPrice := "1610"
	flightPrices := []string{
		"1460",
		"1510",
		"1610",
		"1910",
		"1720",
		"1870",
		"1870",
		"1870",
		"1870",
		"2130",
		"2210",
		"750",
		"1100",
		"970",
		"1470",
		"1470",
		"1470",
		"1470",
		"1470",
		"1470",
		"630",
		"1700",
		"1700",
		"1600",
		"1600",
		"1570",
		"1570",
		"1570",
		"1410",
		"1410",
		"850",
		"1030",
		"1330",
		"1560",
	}

	//username := "赵楠"
	hongbaoNum := "5.88"      //随机红包数量
	phoneNum := "15657178960" //代理商手机号

	userNames := []string{
		"范巧云",
		"王天亮",
		"孙瑛芳",
		"陈永潮",
		"张龙馨",
		"刘艳淑",
		"马静",
		"孙轻敏",
		"张占芝",
		"赵晔",
		"黎群英",
		"肖敏芳",
		"李艳梅",
		"甘永林",
		"董蕾",
		"高建学",
		"陆维良",
		"盛红飞",
		"朱庆华",
		"朱芸",
		"邹子乐",
		"曾广荣",
		"孙勇",
		"姜琳",
		"徐正虎",
		"黄林根",
		"李蓉蓉",
		"黄弋博",
		"蔡玲玲",
		"潘萍萍",
		"汪春玉",
		"李斌",
		"马婧林",
		"方培力",
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
		part09 := "', '0', '2017-09-14 08:00:00', '2017-09-14 08:00:00');"

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
		hongbaosSuijiStr := "INSERT INTO `hd_hongbao`.`hongbaos` (`id`, `user_id`, `service_id`, `provider_id`, `type`, `hongbao_count`, `hongbao_money`, `description`, `created`, `modified`) VALUES (NULL, '" + userID + "', '1', '1', '1', '5.88', '5.88','2', '2017-09-14 08:00:05', '2017-09-14 08:00:05');"

		f, err := os.OpenFile("suijihongbaos.txt", os.O_APPEND, 0666)
		n, err := f.WriteString(hongbaosSuijiStr)
		f.Close()
		check(err)
		print(strconv.Itoa(n) + ",")
	}

	//平台补贴红包
	/**/
	compensateHongbaos := []string{
		"100",
		"102.5",
		"109.5",
		"130",
		"117.5",
		"127.5",
		"127.5",
		"127.5",
		"127.5",
		"145",
		"150",
		"52.5",
		"75",
		"67.5",
		"100",
		"100",
		"100",
		"100",
		"100",
		"100",
		"42.5",
		"115",
		"115",
		"110",
		"110",
		"107.5",
		"107.5",
		"107.5",
		"97.5",
		"97.5",
		"57.5",
		"70",
		"90",
		"107.5",
	}

	for i := 0; i < len(compensateHongbaos); i++ {
		hongbaosStr := "INSERT INTO `hd_hongbao`.`hongbaos` (`id`, `user_id`, `service_id`, `provider_id`, `type`, `hongbao_count`, `hongbao_money`, `description`, `created`, `modified`) VALUES (NULL, '" + userID + "', '1', '1', '1', '" + compensateHongbaos[i] + "', '" + compensateHongbaos[i] + "','4', '2017-09-13 14:00:10', '2017-09-13 14:00:10');"

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
