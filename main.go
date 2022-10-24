package main

import (
	"StoreInventoryManagementSystem/dao"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"log"
	"net/mail"

	//	"fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
	"os"
    "time"
	"StoreInventoryManagementSystem/service"
)
var Email string;
func Call() { //回调函数
    ticker := time.NewTicker(time.Minute)
	for _ = range ticker.C{
		due, err := service.SearchPastDue()
		if err !=nil{
			log.Println(err)
		}else{
			if due != ""{
				service.Alert(due,Email)
			}
		}
	}
	ticker.Stop()
}
func Alarm(){
	due, err := service.SearchPastDue()
	log.Println("发送邮箱:"+Email)
	if err !=nil{
		log.Println(err)
	}else{
		if due != ""{
			service.Alert(due,Email)
		}
	}
}
func main() {
	os.Setenv("FYNE_FONT","STXINWEI.TTF")
	a := app.New()
	a.Settings().SetTheme(theme.DarkTheme())
    w := a.NewWindow("商店存货管理系统")
	w.SetMaster()
	temp := a.NewWindow("商店存货管理系统")
	entry := widget.NewEntry()
	entry.SetPlaceHolder("比如1213605030@qq.com")
	newString := binding.NewString()
	label := widget.NewLabelWithData(newString)
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "请输入报警邮箱(自己邮箱即可)", Widget: entry}},
			OnSubmit: func() {
				Email = entry.Text
				_, err := mail.ParseAddress(Email) //判断邮箱是否合理
				if err != nil{
					box := container.NewVBox(
						container.New(layout.NewGridLayout(3),widget.NewLabel(""),widget.NewLabel("邮箱格式错误"),widget.NewLabel("")),
						&widget.Button{Text: "关闭模式窗口"},
						)
					up := widget.NewPopUp(box, temp.Canvas())
					up.ShowAtPosition(fyne.NewPos(
						// 计算窗口的中央位置
						temp.Canvas().Size().Width/2- up.MinSize().Width/2,
						temp.Canvas().Size().Height/2-up.MinSize().Height/2))
					box.Objects[1].(*widget.Button).OnTapped = func() {
						up.Hide()
					}
				}else{
					get, _ := newString.Get()
					newString.Set("报警邮箱为:"+Email+"\n\n"+get)
					temp.Close()
				}
			},
			}
			temp.SetContent(form)
	temp.Resize(fyne.NewSize(600,100))
	w.Resize(fyne.NewSize(1200,200))
	go Call()
	SepiceInput := widget.NewEntry()
	SepiceInput.SetPlaceHolder("请输入商品种类(水果)")
	NameInput := widget.NewEntry()
	NameInput.SetPlaceHolder("请输入商品名字(香蕉)")
	PurChaseTimeInput := widget.NewEntry()
	PurChaseTimeInput.SetPlaceHolder("请输入进货日期(2006-01-02)")
	PastDueInput := widget.NewEntry()
	PastDueInput.SetPlaceHolder("请输入过期日期(2006-01-02)")
	SSepiceInput := widget.NewEntry()
	SSepiceInput.SetPlaceHolder("请输入商品种类(水果)")
	SNameInput := widget.NewEntry()
	SNameInput.SetPlaceHolder("请输入商品名字(香蕉)")
	SPurChaseTimeInput := widget.NewEntry()
	SPurChaseTimeInput.SetPlaceHolder("请输入进货日期(2006-01-02)")
	SPastDueInput := widget.NewEntry()
	SPastDueInput.SetPlaceHolder("请输入过期日期(2006-01-02)")
	USepiceInput := widget.NewEntry()
	USepiceInput.SetPlaceHolder("请输入商品种类(水果)")
	UNameInput := widget.NewEntry()
	UNameInput.SetPlaceHolder("请输入商品名字(香蕉)")
	UPurChaseTimeInput := widget.NewEntry()
	UPurChaseTimeInput.SetPlaceHolder("请输入进货日期(2006-01-02)")
	UPastDueInput := widget.NewEntry()
	UPastDueInput.SetPlaceHolder("请输入过期日期(2006-01-02)")

	sum, _ ,_:= service.Query(dao.Object{})
	var origin dao.Object
	newString.Set(sum)
	Purchasebutton := widget.NewButtonWithIcon("进货",theme.ContentCopyIcon() ,func() {
		PurChaseTime := PurChaseTimeInput.Text
		_, err := time.Parse("2006-01-02", PurChaseTime)
		if err !=nil{
			log.Println("出货日期格式错误")
			return
		}
		PastDueTime := PastDueInput.Text
		_, err = time.Parse("2006-01-02", PastDueTime)
		if err != nil{
			log.Println("过期日期格式错误")
			return
		}
		if SepiceInput.Text == ""{
			log.Println("种类不能为空")
		}else if NameInput.Text == ""{
			log.Println("名字不能为空")
		}else {
			object := dao.Object{
				SepiceInput.Text,
				NameInput.Text,
				PurChaseTimeInput.Text,
				PastDueInput.Text,
			}
			log.Println("进货了:"+object.ToShow())
			service.PurChase(object)
			query, _, _ := service.Query(dao.Object{})
			newString.Set("进货操作:"+object.ToShow()+"\n"+query)
			origin = object
			USepiceInput.SetText(origin.Specie)
			UNameInput.SetText(origin.Name)
			UPurChaseTimeInput.SetText(origin.PurChaseTime)
			UPastDueInput.SetText(origin.PastDue)
		}
	})
	PastDueButton := widget.NewButtonWithIcon("出货",theme.ContentRemoveIcon(),func(){
		andDelete, err := service.OpenAndDelete()
		query, _, _ := service.Query(dao.Object{})
		if err != nil{
			log.Println(err)
		}else{
			newString.Set(andDelete+query)
		}
	})
	UpdateButton := widget.NewButtonWithIcon("更新",theme.ContentRedoIcon(),func(){
		update, err,o := service.Update(origin, dao.Object{
			USepiceInput.Text,
			UNameInput.Text,
			UPurChaseTimeInput.Text,
			UPastDueInput.Text,
		})
		if err!=nil{
			log.Println(err)
		}else{
			newString.Set(update)
		}
		origin = o
		USepiceInput.SetText(origin.Specie)
		UNameInput.SetText(origin.Name)
		UPurChaseTimeInput.SetText(origin.PurChaseTime)
		UPastDueInput.SetText(origin.PastDue)
	})
	AlarmButton := widget.NewButtonWithIcon("过期报警",theme.ErrorIcon(),func(){
		Alarm()
	})
	ExportButton := widget.NewButtonWithIcon("导出Excel",theme.DocumentPrintIcon(),func(){
		service.WriteToExcel()
	})
	SearchButton := widget.NewButtonWithIcon("查询(输入其中一个或多个)",theme.SearchIcon(),func(){
		if SPurChaseTimeInput.Text != ""{
			SPurChaseTime := SPurChaseTimeInput.Text
			_, err := time.Parse("2006-01-02", SPurChaseTime)
			if err !=nil{
				log.Println("出货日期格式错误")
				return
			}
		}
		if SPastDueInput.Text != ""{
			SPastDueTime := SPastDueInput.Text
			_, err := time.Parse("2006-01-02", SPastDueTime)
			if err !=nil{
				log.Println("出货日期格式错误")
				return
			}
		}
		object := dao.Object{
			Specie:  SSepiceInput.Text,
			Name: SNameInput.Text,
			PurChaseTime: SPurChaseTimeInput.Text,
			PastDue: SPastDueInput.Text,
		}
		query, err, o := service.Query(object)
		if err != nil{
			log.Println(err)
		}
		newString.Set(query)
		origin = o
		USepiceInput.SetText(origin.Specie)
		UNameInput.SetText(origin.Name)
		UPurChaseTimeInput.SetText(origin.PurChaseTime)
		UPastDueInput.SetText(origin.PastDue)
	})
	combo := widget.NewSelect([]string{"按照种类排序", "按照名字排序","按照进货时间排序","按照过期时间排序"}, func(value string) {
	})
	combo.SetSelectedIndex(0)
	SortButton := widget.NewButtonWithIcon("排序",theme.StorageIcon(),func(){
		sort, err := service.Sort(combo.Selected)
		if err!=nil{
			log.Panicln(err)
		}else{
			newString.Set(sort)
		}
	})
	f1 := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), label, layout.NewSpacer())
	f2 := container.New(layout.NewGridLayout(5),widget.NewLabel(""),widget.NewLabel("商品种类"),widget.NewLabel("商品名字"),widget.NewLabel("进货时间"),widget.NewLabel("过期时间"))
	f3 := container.New(layout.NewGridLayout(5), Purchasebutton,SepiceInput,NameInput,PurChaseTimeInput,PastDueInput)
	f4 := container.New(layout.NewGridLayoutWithColumns(5), SearchButton,SSepiceInput,SNameInput,SPurChaseTimeInput,SPastDueInput)
	f5 := container.New(layout.NewGridLayoutWithColumns(5),UpdateButton,USepiceInput,UNameInput,UPurChaseTimeInput,UPastDueInput)
	f6:= container.New(layout.NewGridLayout(5), PastDueButton,AlarmButton,SortButton,combo,ExportButton)
	w.SetContent(container.New(layout.NewVBoxLayout(), f1,f2,f3,f4,f5,f6))
	w.Show()
	temp.ShowAndRun()

}