package WPRLibs

import "testing"

func Test_decodePDMFile(t *testing.T) {
	var isSuccess bool = false
	fList, err := ListCurrDir("../data/pdm", ".pdm")
	Check(err)

	for _, f := range fList {
		var pRootModel PdmRootModel
		var pdmLog UpdatePdmFileLog
		if f.FName == "WPR_PhysicalDataModel_Test.pdm" {
			t.Logf("准备处理文件： %s \n", f.FPath)
			pRootModel, pdmLog, err = decodePDMFile(f.FPath)
			Check(err)

			// t.Log("pRootModel:", pRootModel)
			t.Log("pdmLog:", pdmLog)

			// 对解析数据进行效验
			switch {
			case pRootModel.Root.Children.Model.Id != "o2":
				t.Error("Model.Id:", pRootModel.Root.Children.Model.Id)
				fallthrough
			case pRootModel.Root.Children.Model.Name != "WPR_PhysicalDataModel_Test":
				t.Error("Model.Name:", pRootModel.Root.Children.Model.Name)
				fallthrough
			case len(pRootModel.Root.Children.Model.Users.Users) != 2:
				t.Error("Users counter:", len(pRootModel.Root.Children.Model.Users.Users))
				fallthrough
			case len(pRootModel.Root.Children.Model.Users.Users) != 2:
				t.Error("Users counter:", len(pRootModel.Root.Children.Model.Users.Users))
				fallthrough
			case pRootModel.Root.Children.Model.Users.Users[0].Name != "Test_User_Name_1":
				t.Error("Users Name:", pRootModel.Root.Children.Model.Users.Users[0].Name)
				fallthrough
			case len(pRootModel.Root.Children.Model.Tables.Tables) != 3:
				t.Error("Table counter:", len(pRootModel.Root.Children.Model.Tables.Tables))
				fallthrough
			case len(pRootModel.Root.Children.Model.Tables.Tables[0].Columns.Columns) != 6:
				t.Error("Table[0] Column counter:", len(pRootModel.Root.Children.Model.Tables.Tables[0].Columns.Columns))
				fallthrough
			case len(pRootModel.Root.Children.Model.Tables.Tables[0].Keys.Keys) != 2:
				t.Error("Table[0] Keys counter:", len(pRootModel.Root.Children.Model.Tables.Tables[0].Keys.Keys))
				fallthrough
			case len(pRootModel.Root.Children.Model.Tables.Tables[0].Indexes.Indexes) != 1:
				t.Error("Table[0] Keys counter:", len(pRootModel.Root.Children.Model.Tables.Tables[0].Indexes.Indexes))
				fallthrough
			default:
				isSuccess = true
			}

			break

		} else {
			t.Logf("准备处理文件： %s \n", f.FPath)
			t.Log("pRootModel:", pRootModel)
			t.Log("pdmLog:", pdmLog)
			isSuccess = false
		}
	}

	if !isSuccess {
		t.Error("解析错误！")
	}

}

func Test_ExchangeToOutput(t *testing.T) {
	var isSuccess bool = false
	fList, err := ListCurrDir("../data/pdm", ".pdm")
	Check(err)

	for _, f := range fList {
		var pRootModel PdmRootModel
		var pdmLog UpdatePdmFileLog
		var outTab OutputTables
		if f.FName == "WPR_PhysicalDataModel_Test.pdm" {
			t.Logf("准备处理文件： %s \n", f.FPath)
			pRootModel, pdmLog, err = decodePDMFile(f.FPath)
			Check(err)

			// t.Log("pRootModel:", pRootModel)
			t.Log("pdmLogIn:", pdmLog)

			// 对数据进行转换
			outTab, err = pRootModel.ExchangeToOutput(&pdmLog)
			Check(err)
			t.Log("outTab:", outTab.Tables)
			t.Log("pdmLogOut:", pdmLog)

			// 对输出数据进行效验
			switch {
			case pdmLog.UpdateTabColNum != 15:
				t.Error("pdmLog.UpdateTabColNum:", pdmLog.UpdateTabColNum)
				fallthrough
			case len(outTab.Tables) != 3:
				t.Error("len(outTab.Tables):", len(outTab.Tables))
				fallthrough
			default:
				isSuccess = true
			}

			break
		} else {
			t.Log("pdmLog:", pdmLog)
			isSuccess = false
		}
	}

	if !isSuccess {
		t.Error("转换错误！")
	}
}
