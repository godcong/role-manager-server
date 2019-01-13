package model

import (
	"github.com/mongodb/mongo-go-driver/bson"
)

// Frame ...
type Frame struct {
	URL       string  `json:"url,omitempty"`
	Offset    int     `json:"offset"`
	Rate      float64 `json:"rate,omitempty"`
	SfaceData []struct {
		Faces []struct {
			ID   string  `json:"id,omitempty"`
			Name string  `json:"name,omitempty"`
			Rate float64 `json:"rate,omitempty"`
		} `json:"faces,omitempty"`
		H int `json:"h,omitempty"`
		W int `json:"w,omitempty"`
		X int `json:"x,omitempty"`
		Y int `json:"y,omitempty"`
	} `json:"sfaceData,omitempty"`
}

// Result ...
type Result struct {
	Frames     []Frame `json:"frames"`
	Label      string  `json:"label"`
	Rate       float64 `json:"rate"`
	Scene      string  `json:"scene"`
	Suggestion string  `json:"suggestion"`
}

// ResultData ...
type ResultData struct {
	Code int `json:"code"`
	Data []struct {
		Code   int    `json:"code"`
		DataID string `json:"dataId"`
		Extras struct {
		} `json:"extras"`
		Msg     string   `json:"msg"`
		Results []Result `json:"results"`
		TaskID  string   `json:"taskId"`
		URL     string   `json:"url"`
	} `json:"data"`
	Msg       string `json:"msg"`
	RequestID string `json:"requestId"`
}

// MediaCensor ...
type MediaCensor struct {
	Model `bson:",inline"`
	//MediaID    primitive.ObjectID `bson:"media_id"`
	Verify     string        `bson:"verify"` //人工验证
	RequestKey string        `bson:"request_key"`
	ResultData []*ResultData `bson:"result_data,omitempty"`
}

func (m *MediaCensor) _Name() string {
	return "media_censor"
}

// CreateIfNotExist ...
func (m *MediaCensor) CreateIfNotExist() error {
	return CreateIfNotExist(m)
}

// Create ...
func (m *MediaCensor) Create() error {
	return InsertOne(m)
}

// Update ...
func (m *MediaCensor) Update() error {
	return UpdateOne(m)
}

// Delete ...
func (m *MediaCensor) Delete() error {
	return DeleteByID(m)
}

// Find ...
func (m *MediaCensor) Find() error {
	return FindByID(m)
}

// FindByKey ...
func (m *MediaCensor) FindByKey() error {
	return FindOne(m, bson.M{
		"request_key": m.RequestKey,
	})
}

// Media ...
func (m *MediaCensor) Media() (*Media, error) {
	media := NewMedia()
	media.CensorID = m.ID
	err := media.Find()
	return media, err
}

// NewMediaCensor ...
func NewMediaCensor() *MediaCensor {
	return &MediaCensor{
		Model: model(),
	}
}
