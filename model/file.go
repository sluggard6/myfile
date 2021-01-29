package model

//Policy 存储策略
type Policy struct {
	Model
	Type PolicyType
	size int64
	Path string
	Sha  string
}

//PolicyType 策略类型
type PolicyType string

const (
	//MyFile 自带的文件存储方案
	MyFile PolicyType = "myfile"
)

//File 文件
type File struct {
	Model
	Name     string `json:"name"`
	Ext      string `json:"ext"`
	FolderID uint   `json:"-"`
	Size     uint   `json:"size"`
	PolicyID uint   `json:"-"`
	Policy   Policy `gorm:"foreignKey:PolicyID" json:"-"`
}

//GetFilesByFolderID 查询目录下的所有文件
func (file *File) GetFilesByFolderID() (files *[]File, err error) {
	files = &[]File{}
	err = db.Where("folder_id=?", file.FolderID).Find(files).Error
	return
}
