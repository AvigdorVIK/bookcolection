package handler
 
import (
	"bookcolection/model"
	"fmt"
)

func CreateRecord(name, receng, info string)error{
	record := model.Book{
		Name: name,
		Receng: receng,
        Info: info,
	}
    fmt.Println("zapis stvoreno")
	return nil

}

func ShowRecord()error{
	for _, recor := range record{
		fmt.Printf("Nazva knygi:", %s, record.Name)
		fmt.Printf("Vidguk do knygi:", %s, record.Receng)
		fmt.Printf("Dodatkova informacia:", %s, record.Info)
		fmt.Printf("______________________________________")
	}
	return nil
}

func DeleteRecord(name string)error{


}

