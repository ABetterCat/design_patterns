package simple_factory

// 创建ES返回结构工厂
func ResESStruct(tag string) EsDataRes {
	if tag == "false" {
		return &LogstashRes{}
	}
	return &LogstashGYRes{}
}

// 定义基础接口
type BaseEsRes interface {
	ReError() *LogstashError
	GetStatus(int)
}

// 定义基础结构体
type LogstashResBase struct {
	Error  LogstashError `json:"error"`
	Status int           `json:"status"`
}

// 实现基础接口
func (l *LogstashResBase) ReError() *LogstashError {
	return &l.Error
}

func (l *LogstashResBase) GetStatus(status int) {
	l.Status = status
	return
}

// 定义延展接口
type EsDataRes interface {
	BaseEsRes
	ReData() interface{}
	ReTotal() int
	ReAggregations() Aggregations
	GetName() string
}

// 定义延展结构体
type LogstashRes struct {
	LogstashResBase
	Data LogstashLog `json:"data"`
}

type LogstashGYRes struct {
	LogstashResBase
	Data LogstashLogGY `json:"data"`
}

// 实现延展接口
func (l *LogstashRes) GetName() string {
	return "北京"
}
func (l *LogstashGYRes) GetName() string {
	return "贵阳"
}

func (l *LogstashRes) ReData() interface{} {
	return &l.Data
}
func (l *LogstashGYRes) ReData() interface{} {
	return &l.Data
}

func (l *LogstashRes) ReTotal() int {
	return l.Data.Hits.Total
}
func (l *LogstashGYRes) ReTotal() int {
	return l.Data.Hits.Total.Value
}

//func (l *LogstashGYRes) GetStatus(status int) {
//	l.Status = status
//	return
//}

func (l *LogstashRes) ReAggregations() Aggregations {
	return l.Data.Aggregations
}
func (l *LogstashGYRes) ReAggregations() Aggregations {
	return l.Data.Aggregations
}

type LogstashLogGY struct {
	Shards struct {
		Failed     int `json:"failed"`
		Skipped    int `json:"skipped"`
		Successful int `json:"successful"`
		Total      int `json:"total"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		Hits []LogstashHits `json:"hits"`
	}
	Aggregations Aggregations `json:"aggregations"`
}

type LogstashError struct {
	Error struct {
		Type   string `json:"type"`
		Reason string `json:"reason"`
		Line   int    `json:"line"`
		Col    int    `json:"col"`
	}
}

type LogstashLog struct {
	Shards struct {
		Failed     int `json:"failed"`
		Skipped    int `json:"skipped"`
		Successful int `json:"successful"`
		Total      int `json:"total"`
	} `json:"_shards"`
	Hits struct {
		Total int            `json:"total"`
		Hits  []LogstashHits `json:"hits"`
	}
	Aggregations Aggregations `json:"aggregations"`
}
type Aggregations struct {
	Models Models `json:"models"`
}

type Models struct {
	DocCountErrorUpperBound int       `json:"doc_count_error_upper_bound"`
	SumOtherDocCount        int       `json:"sum_other_doc_count"`
	Buckets                 []Buckets `json:"buckets"`
}

//struct {
//Models struct {
//DocCountErrorUpperBound float64       `json:"doc_count_error_upper_bound"`
//SumOtherDocCount        float64       `json:"sum_other_doc_count"`
//Buckets                 []Buckets `json:"buckets"`
//} `json:"models"`
//} `json:"aggregations"`

type Buckets struct {
	Key      string `json:"key"`
	DocCount int    `json:"doc_count"`
}

type LogstashHits struct {
	Id     string         `json:"_id"`
	Index  string         `json:"_index"`
	Score  float64        `json:"_score"`
	Source LogstashSource `json:"_source"`
	Type   string         `json:"_type"`
}

type LogstashSource struct {
	RequestCookie string `json:"request_cookie"`
	//downstream_User-Info
	DownstreamUserInfo string `json:"downstream_User-Info"`
	Time               string `json:"time"`
	Timestamp          string `json:"@timestamp"`
	RequestMethod      string `json:"RequestMethod"`
	RequestPath        string `json:"RequestPath"`
	RequestHost        string `json:"RequestHost"`
}
