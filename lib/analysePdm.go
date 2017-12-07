package WPRLibs

/* 功能说明：
 *  负责处理PDM文件的解析
 *
 *
 */
import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

/* 功能说明：
 *  将传入的pdm文件按照原始格式解析出来
 *
 *
 */
func decodePDMFile(filename string) (PRootModel PdmRootModel, ulog UpdatePdmFileLog, err error) {
	// 初始化
	ulog = UpdatePdmFileLog{}
	ulog.UpdateTime = time.Now()

	// 打开文件
	file, err := os.Open(filename)
	Check(err)
	defer file.Close()
	beego.Debug("filename: ", filename)

	ulog.UpdateFileName = file.Name()
	ulog.UpdateFileStatus, err = file.Stat()
	Check(err)

	// 读取文件内容
	data, err := ioutil.ReadAll(file)
	Check(err)

	// 解析xml文件内容
	PRootModel = PdmRootModel{}
	err = xml.Unmarshal(data, &PRootModel)
	Check(err)

	return PRootModel, ulog, nil
}

/* 功能说明：
 *  针对Pdm解析后格式进行一些基本处理
 *  将原始格式转换为展示输出的格式
 *
 */
func (this *PdmRootModel) ExchangeToOutput(ulog *UpdatePdmFileLog) (out OutputTables, err error) {
	out = OutputTables{}
	userMap := make(map[string]PdmUserInfo)

	// 加载用户列表，处理Table信息需要用到
	for _, u := range this.Root.Children.Model.Users.Users {
		userMap[u.Id] = u
	}

	// TODO: 加载外键关系，处理Table信息需要用到
	// for _, r := range this.Root.Children.Model.References.References {
	// }

	// 解析基本的Table结构
	ulog.UpdateTabNum = len(this.Root.Children.Model.Tables.Tables)
	for _, t := range this.Root.Children.Model.Tables.Tables {
		o := OutputTable{}
		keyMap := make(map[string][]PdmKey)
		idxMap := make(map[string][]PdmIndex)

		// 处理表的基本信息
		o.TableCode = t.Code
		o.TableName = t.Name
		o.TableComment = t.Comment

		// 处理表Keys
		// 将原结构Key=>[]col的结构映射为一个map[col][]key_id
		for _, k := range t.Keys.Keys {
			for _, kc := range k.KeyColumns.KeyColumns {
				keyMap[kc.Ref] = append(keyMap[kc.Ref], k)
			}
		}

		// 处理表Indexes
		// 将原结构Idx=>[]col的结构映射为一个map[col][]idx_id
		for _, i := range t.Indexes.Indexes {
			for _, ic := range i.IndexColumns.IndexColumns {
				for _, icc := range ic.IdxColumn.Columns {
					idxMap[icc.Ref] = append(idxMap[icc.Ref], i)
				}
			}
		}

		// 处理表Owner
		o.TableOwner = userMap[t.Owner.User.Ref].Name

		// 处理表Columns
		ulog.UpdateTabColNum += len(t.Columns.Columns)
		for _, c := range t.Columns.Columns {
			outCol := OutputColumn{}
			outCol.ColumnId = c.Id
			outCol.ColumnCode = c.Code
			outCol.ColumnName = c.Name
			outCol.ColumnComment = c.Comment
			outCol.ColumnDataType = c.DataType
			outCol.ColumnDefaultValue = c.DefaultValue
			if c.Length != "" {
				outCol.ColumnLength, err = strconv.Atoi(c.Length)
				Check(err)
			}
			outCol.ColumnMandatory = (c.ColumnMandatory == "1")

			// TODO: 附属扩展字段需要解析
			outCol.ColumnExtendedAttributesText = c.ExtendedAttributesText

			// 设置Key信息
			_, ok := keyMap[c.Id]
			if ok {
				for _, a := range keyMap[c.Id] {
					if t.PrimaryKey.Key == a.Id {
						outCol.ColumnIsPrimaryKey = true
						outCol.ColumnkeyCode = append(outCol.ColumnkeyCode, "PK-"+a.Code)
					} else {
						outCol.ColumnkeyCode = append(outCol.ColumnkeyCode, "AK-"+a.Code)
					}
				}
			}

			// 设置IDX信息
			_, ok = idxMap[c.Id]
			if ok {
				for _, idx := range idxMap[c.Id] {
					outCol.ColumnIndexCode = append(outCol.ColumnIndexCode, idx.Code)
				}
			}

			// TODO: 设置FK信息

			// 字段信息写入表内
			o.TableColumns = append(o.TableColumns, outCol)
		}
		out.Tables = append(out.Tables, o)
	}
	return out, nil
}
