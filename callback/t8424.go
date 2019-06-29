package callback

import (
	"errors"

	"github.com/sangx2/ebest/model"
	"github.com/sangx2/ebest/wrapper"
)

// T8424 업종전체조회
type T8424 struct {
	InBlock  model.T8424InBlock
	OutBlock []model.T8424OutBlock

	TPS, LPP int

	ReceiveDataChan                chan wrapper.XaQueryReceiveData
	ReceiveMessageChan             chan wrapper.XaQueryReceiveMessage
	ReceiveChartRealDataChan       chan wrapper.XaQueryReceiveChartRealData
	ReceiveChartSearchRealDataChan chan wrapper.XaQueryReceiveSearchRealData
}

func NewT8424() *T8424 {
	return &T8424{
		TPS: 1, LPP: 200,
		ReceiveDataChan:                make(chan wrapper.XaQueryReceiveData, 1),
		ReceiveMessageChan:             make(chan wrapper.XaQueryReceiveMessage, 1),
		ReceiveChartRealDataChan:       make(chan wrapper.XaQueryReceiveChartRealData, 1),
		ReceiveChartSearchRealDataChan: make(chan wrapper.XaQueryReceiveSearchRealData, 1),
	}
}

func (t T8424) GetTPS() int {
	return t.TPS
}

func (t T8424) GetLPP() int {
	return t.LPP
}

func (t T8424) GetReceiveDataChan() chan wrapper.XaQueryReceiveData {
	return t.ReceiveDataChan
}

func (t T8424) GetReceiveMessageChan() chan wrapper.XaQueryReceiveMessage {
	return t.ReceiveMessageChan
}

func (t T8424) GetReceiveChartRealDataChan() chan wrapper.XaQueryReceiveChartRealData {
	return t.ReceiveChartRealDataChan
}

func (t T8424) GetReceiveChartSearchRealDataChan() chan wrapper.XaQueryReceiveSearchRealData {
	return t.ReceiveChartSearchRealDataChan
}

func (t *T8424) SetFieldData(e *wrapper.Ebest, resPath string, inBlock1 interface{}, inBlock2 interface{}) error {
	e.ResFileName(resPath + "t8424.res")

	i, ok := inBlock1.(model.T8424InBlock)
	if !ok {
		return errors.New("Invalid inBlock1")
	}

	t.InBlock = i

	e.SetFieldData("t8424InBlock", "gubun1", 0, t.InBlock.Gubun1)

	return nil
}

func (t T8424) GetOutBlock() (interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) {
	return t.OutBlock, nil, nil, nil, nil, nil
}

func (t *T8424) ReceivedData(e *wrapper.Ebest, x wrapper.XaQueryReceiveData) {
	TRcount := e.GetBlockCount("t8424OutBlock")

	t.OutBlock = nil

	for i := 0; i < int(TRcount); i++ {
		tr := model.T8424OutBlock{
			Hname:  e.GetFieldData("t8424OutBlock", "hname", i),
			Upcode: e.GetFieldData("t8424OutBlock", "upcode", i)}
		t.OutBlock = append(t.OutBlock, tr)
	}

	t.ReceiveDataChan <- x
}

func (t T8424) ReceivedMessage(e *wrapper.Ebest, x wrapper.XaQueryReceiveMessage) {
	t.ReceiveMessageChan <- x
}

func (t T8424) ReceivedChartRealData(e *wrapper.Ebest, x wrapper.XaQueryReceiveChartRealData) {
	t.ReceiveChartRealDataChan <- x
}

func (t T8424) ReceivedSearchRealData(e *wrapper.Ebest, x wrapper.XaQueryReceiveSearchRealData) {
	t.ReceiveChartSearchRealDataChan <- x
}