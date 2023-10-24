package scripts

type WechatReadingNotes struct {
	NotesFilePath string
	MemoTime      string
	Visibility    string
	Tags          []string
	Url           string
	Token         string
}

// 每一条 memo 的结构
type ImportMemo struct {
	Content    string `json:"content"`
	Visibility string `json:"visibility"`
	CreatedTs  int64  `json:"createdTs"`
	UpdateTs   int64  `json:"updateTs"`
	DisplayTs  int64  `json:"displayTs"`
}

type ImportTag struct {
	Name string `json:"name"`
}

func NewWechatReadingNotes(notesFilePath string, memoTime string, visibility string, tags []string, url string, token string) *WechatReadingNotes {
	return &WechatReadingNotes{NotesFilePath: notesFilePath, MemoTime: memoTime, Visibility: visibility, Tags: tags, Url: url, Token: token}
}

func NewImportMemo(content string, visibility string, Time int64) *ImportMemo {
	return &ImportMemo{Content: content, Visibility: visibility, CreatedTs: Time, UpdateTs: Time, DisplayTs: Time}
}

func NewImportTag(name string) *ImportTag {
	return &ImportTag{Name: name}
}
