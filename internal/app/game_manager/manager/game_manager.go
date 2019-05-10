package manager

import (
	"encoding/json"
	"github.com/mingz2013/game-micro-server-go/internal/pkg/msg"
)

//```
//{
//"gameMap": {
//	"gameId123": {
//		"processIds": [1, 2, 3]
//}
//}
//
//
//"processMap": {
//	"processId123": {
//		"gameId": 123,
//		"tableIds": [1, 2, 3]
//	}
//}
//
//
//"tableMap": {
//
//	"tableid123": {
//		"gameId": 123,
//		"processId": 123,
//	}
//}
//
//
//}
//
//```

type GameManager struct {
	// 这里固化到redis，加锁操作
	gameMap    map[string]map[string]interface{} // {"gameId234": {"processIds": [1, 2, 3]}}
	ProcessMap map[string]map[string]interface{} // {"processId123": {"gameId": 1, "", "tableIds": ["1", "2"]}}
	tableMap   map[string]map[string]interface{} // {"tableId123":{"gameId", 123, "processId": 123}}
}

func (m *GameManager) Init() {
	m.gameMap = make(map[string]map[string]interface{})
	m.ProcessMap = make(map[string]map[string]interface{})
	m.tableMap = make(map[string]map[string]interface{})

}

func NewGameManager() *GameManager {
	m := GameManager{}
	m.Init()
	return &m
}

func (m *GameManager) CreateTable(userId int, gameId int) (tableId string, err error) {
	//  找到一个相同gameId的桌子进程，请求创建一张桌子，创建成功后，登记相关消息，如果桌子数量不够了，需要新建进程

	// 加锁

	processId := m.findOneProcess(gameId)

	channel := m.ProcessMap[processId]["channel"].(string)
	message := msg.NewMsg()
	message.SetCmd("table")
	message.SetParams(map[string]interface{}{
		"action": "create",
	})
	data, _ := json.Marshal(message)
	retData := m.Query(channel, data)

	var retMsg msg.Msg
	json.Unmarshal(retData, retMsg)

	if retMsg["retcode"] == 0 {
		tableId = retMsg["tableId"].(string)

		m.ProcessMap[processId]["gameId"] = gameId
		m.ProcessMap[processId]["tableIds"] = append(m.ProcessMap[processId]["tableIds"].([]string), tableId)

		m.tableMap[tableId] = map[string]interface{}{
			"gameId":    gameId,
			"processId": processId,
		}

	} else {
		err = retMsg["retcode"].(error)
	}

	return
}

func (m *GameManager) DelTable() {

}

func (m *GameManager) findOneProcess(gameId int) (processId string) {
	// 检查此gameId下的每个进程，看哪个进程的桌子数量没满且最多，选哪个
	return
}

func (m *GameManager) Send(channel string, message []byte) {

}

func (m *GameManager) Query(channel string, message []byte) (retMsg []byte) {
	return
}
