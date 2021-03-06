package res

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
)

// T8424 업종전체조회
type T8424InBlock struct {
	Gubun1 string `json:"구분1"`
}

type T8424OutBlock struct {
	gorm.Model `json:"-"`

	Hname  string `json:"업종명"`
	Upcode string `json:"업종코드"`
}

func (t T8424OutBlock) ToJSON() string {
	b, _ := json.Marshal(t)
	return string(b)
}
