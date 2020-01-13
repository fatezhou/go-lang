package test

import (
	"encoding/json"
	"flag"
	"fmt"
	"sync"
	"time"
	"zoyee-tool/internal/pkg/utils"
	http "zoyee-tool/internal/pkg/http"
)

type test struct{
	UserCount int
	UserStaySecond int
	Url string
	userData *sync.Map
	wg *sync.WaitGroup
}

type user struct{
	UserId string
	EnterTime int64
	RecvTime int64
	InTime int64
	Succ bool
}

type ResData struct {
	Code int32 `json:"code"`
}

type Res struct{
	Data ResData `json:"data"`
}

var Test = &test{}

func (t *test)ParseCmd(){
	flag.IntVar(&t.UserCount, "c", 1, "user count")
	flag.IntVar(&t.UserStaySecond, "t", 60, "user stay seconds")
	flag.StringVar(&t.Url, "url", "queue1", "")
	flag.Parse()
}

func (t *test)OneDo(user string, usr *user){
	//t.wg.Done()
	urlQ := "https://" + t.Url + ".youyueworld.com/apis/queue"
	urlU := "https://" + t.Url + ".youyueworld.com/apis/update_queue_time_by_token"
 	c := http.HttpClient{}
	str := fmt.Sprintf(`{"token":"%s"}`, usr.UserId)
	usr.EnterTime = time.Now().Unix()
	resp := c.Post(urlQ, str, nil)
	fmt.Printf(resp)
	usr.RecvTime = time.Now().Unix()
	res := Res{}
	for ;;{
		if json.Unmarshal([]byte(resp), &res) == nil{
			//fmt.Printf("%+v", res)
			if res.Data.Code == 1{
				usr.InTime = time.Now().Unix()
				usr.Succ = true
				break
			}else{
				resp = c.Post(urlU, str, nil)
			}
		}else{
			break
		}
		if time.Now().Unix()  - usr.RecvTime >= int64(t.UserStaySecond){
			usr.Succ = true
			break
		}
		time.Sleep(60 * time.Second)
	}
	fmt.Printf("[%+v]\n", usr)
	t.wg.Done()
}

func (t *test)Run(){
	t.userData = &sync.Map{}
	t.wg = &sync.WaitGroup{}
	rand := utils.Rand{}
	t.wg.Add(t.UserCount)
	for i := 0; i < t.UserCount; i++{
		strUser := string(rand.RandString(32, utils.KC_RAND_KIND_ALL))
		fmt.Printf("%s\n", strUser)
		usr := &user{
			UserId: strUser,
		}
		t.userData.Store(strUser, usr)

		go func() {
			t.OneDo(strUser, usr)
		}()
	}
	t.wg.Wait()
	t.Print()
}

func (t *test)Print(){
	Succ := int(0)
	Fail := int(0)
	SumWait := int(0)
	t.userData.Range(func(key, value interface{}) bool {
		ptr := value.(*user)
		if ptr.Succ{
			Succ++
			SumWait += int(ptr.EnterTime - ptr.RecvTime)
		}else{
			Fail++
		}
		return true
	})

	if Succ != 0{
		fmt.Printf("Succ:[%d]\nFail[%d]\nWaitTime.Avg[%d]", Succ, Fail, SumWait / Succ)
	}else{
		fmt.Printf("No one succ!")
	}
}