package main

import (
	"encoding/json"
	"fmt"
)

const str = `
{
	"goods": {
		"830000023": {
			"price": 100,
			"prop": {},
			"name": "1钻石",
			"type": "0"
		},
		"830000016": {
			"price": 600,
			"prop": {},
			"name": "6钻石",
			"type": "0"
		},
		"830000017": {
			"price": 1200,
			"prop": {},
			"name": "12钻石",
			"type": "0"
		},
		"830000024": {
			"price": 1800,
			"prop": {},
			"name": "18钻石",
			"type": "0"
		},
		"830000018": {
			"price": 3000,
			"prop": {},
			"name": "30钻石",
			"type": "0"
		},
		"830000019": {
			"price": 5000,
			"prop": {},
			"name": "50钻石",
			"type": "0"
		},
		"830000020": {
			"price": 10800,
			"prop": {},
			"name": "108钻石",
			"type": "0"
		},
		"830000021": {
			"price": 32800,
			"prop": {},
			"name": "328钻石",
			"type": "0"
		},
		"830000022": {
			"price": 64800,
			"prop": {},
			"name": "648钻石",
			"type": "0"
		}
	},
	"roll": [{
		"prize": "5000\u9c7c\u5e01",
		"icon": "https:\/\/buyu-potal.jiaxianghudong.com\/static\/prop_001.png"
	}, {
		"prize": "\u9752\u94dc\u5f39\u5934*1",
		"icon": "https:\/\/buyu-potal.jiaxianghudong.com\/static\/prop_021.png"
	}, {
		"prize": "20\u6c34\u6676",
		"icon": "https:\/\/buyu-potal.jiaxianghudong.com\/static\/prop_002.png"
	}, {
		"prize": "20000\u9c7c\u5e01",
		"icon": "https:\/\/buyu-potal.jiaxianghudong.com\/static\/prop_001.png"
	}, {
		"prize": "15000\u9c7c\u5e01",
		"icon": "https:\/\/buyu-potal.jiaxianghudong.com\/static\/prop_001.png"
	}, {
		"prize": "60000\u9c7c\u5e01",
		"icon": "https:\/\/buyu-potal.jiaxianghudong.com\/static\/prop_001.png"
	}, {
		"prize": "10000\u9c7c\u5e01",
		"icon": "https:\/\/buyu-potal.jiaxianghudong.com\/static\/prop_001.png"
	}, {
		"prize": "\u9ed1\u94c1\u5f39\u5934*1",
		"icon": "https:\/\/buyu-potal.jiaxianghudong.com\/static\/prop_035.png"
	}, {
		"prize": "10\u6c34\u6676",
		"icon": "https:\/\/buyu-potal.jiaxianghudong.com\/static\/prop_002.png"
	}, {
		"prize": "100000\u9c7c\u5e01",
		"icon": "https:\/\/buyu-potal.jiaxianghudong.com\/static\/prop_001.png"
	}],
	"draw": {
		"rule_format": "<ol><li>1.每日可免费抽奖1次.每充值满%d元可增加1次抽奖机会,未使用的抽奖次数每月月底重置<\/li><li>2.每次抽奖累计%d点幸运值.幸运值满值为%d,满值后再抽一次必得双倍奖励,并重新累计幸运值.<\/li><\/ol>",
		"charge_base": 20,
		"one_draw_add": 10,
		"critical_draw": 100,
		"pay_rule_format": "<ol><li>1.该页面包含账号信息,请勿发送或分享.<\/li><li>2.本活动为限时活动,随时可能停止.<\/li><li>3.每充值%d元,可获得1次免费抽奖次数<\/li><\/ol>"
	},
	"roll_goods": [{
		"id": "2",
		"goodsid": "0",
		"prize": "5000\u9c7c\u5e01",
		"probability": "100",
		"data": "1-5000",
		"crater_time": null,
		"update_time": null,
		"icon": "https:\/\/buyu-potal.jiaxianghudong.com\/static\/prop_001.png"
	}, {
		"id": "3",
		"goodsid": "0",
		"prize": "\u9752\u94dc\u5f39\u5934*1",
		"probability": "2",
		"data": "21-1",
		"crater_time": "2019-01-16 10:25:35",
		"update_time": "2019-01-16 10:25:35",
		"icon": "https:\/\/buyu-potal.jiaxianghudong.com\/static\/prop_021.png"
	}, {
		"id": "4",
		"goodsid": "0",
		"prize": "20\u6c34\u6676",
		"probability": "2",
		"data": "2-20",
		"crater_time": "2019-01-16 10:26:22",
		"update_time": "2019-01-16 10:26:22",
		"icon": "https:\/\/buyu-potal.jiaxianghudong.com\/static\/prop_002.png"
	}, {
		"id": "5",
		"goodsid": "0",
		"prize": "20000\u9c7c\u5e01",
		"probability": "10",
		"data": "1-20000",
		"crater_time": "2019-01-16 10:28:29",
		"update_time": "2019-01-16 10:28:29",
		"icon": "https:\/\/buyu-potal.jiaxianghudong.com\/static\/prop_001.png"
	}, {
		"id": "10",
		"goodsid": "0",
		"prize": "15000\u9c7c\u5e01",
		"probability": "12",
		"data": "1-15000",
		"crater_time": "2019-01-17 18:30:26",
		"update_time": "2019-01-17 18:30:26",
		"icon": "https:\/\/buyu-potal.jiaxianghudong.com\/static\/prop_001.png"
	}, {
		"id": "11",
		"goodsid": "1",
		"prize": "60000\u9c7c\u5e01",
		"probability": "3",
		"data": "1-60000",
		"crater_time": "2019-01-17 19:13:26",
		"update_time": "2019-01-17 19:13:26",
		"icon": "https:\/\/buyu-potal.jiaxianghudong.com\/static\/prop_001.png"
	}, {
		"id": "12",
		"goodsid": "0",
		"prize": "10000\u9c7c\u5e01",
		"probability": "20",
		"data": "1-10000",
		"crater_time": "2019-01-18 10:02:41",
		"update_time": "2019-01-18 10:02:41",
		"icon": "https:\/\/buyu-potal.jiaxianghudong.com\/static\/prop_001.png"
	}, {
		"id": "15",
		"goodsid": "0",
		"prize": "\u9ed1\u94c1\u5f39\u5934*1",
		"probability": "10",
		"data": "35-1",
		"crater_time": "2020-04-16 10:13:00",
		"update_time": "2020-04-16 10:13:00",
		"icon": "https:\/\/buyu-potal.jiaxianghudong.com\/static\/prop_035.png"
	}, {
		"id": "16",
		"goodsid": "0",
		"prize": "10\u6c34\u6676",
		"probability": "20",
		"data": "2-10",
		"crater_time": "2020-04-16 10:15:12",
		"update_time": "2020-04-16 10:15:12",
		"icon": "https:\/\/buyu-potal.jiaxianghudong.com\/static\/prop_002.png"
	}, {
		"id": "18",
		"goodsid": "0",
		"prize": "100000\u9c7c\u5e01",
		"probability": "2",
		"data": "1-100000",
		"crater_time": "2020-04-16 10:34:27",
		"update_time": "2020-04-16 10:34:27",
		"icon": "https:\/\/buyu-potal.jiaxianghudong.com\/static\/prop_001.png"
	}]
}
`

type Wechat struct {
	Goods map[string]struct {
		Price int32    `json:"price"`
		Prop  struct{} `json:"prop"`
		Name  string   `json:"name"`
		Type  string   `json:"type"`
	} `json:"goods"`
	Roll      []RollItem             `json:"roll"`
	Draw      map[string]interface{} `json:"draw"`
	RollGoods []RollGoods            `json:"roll_goods"`
}

type RollItem struct {
	Prize string `json:"prize"`
	Icon  string `json:"icon"`
}

type RollGoods struct {
	Id          string `json:"id"`
	GoodsId     string `json:"goodsid"`
	Prize       string `json:"prize"`
	Probability string `json:"probability"`
	Data        string `json:"data"`
	CreateTime  string `json:"crater_time"`
	UpdateTime  string `json:"update_time"`
	Icon        string `json:"icon"`
}




func main(){
	arr := make([]string, 0)
	str := `["1", "2"]`
	json.Unmarshal([]byte(str), &arr)
	fmt.Printf("%+v", arr)
}
