package entity

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddHot(t *testing.T) {
	//CreateTable()
	hot := &Hot{Name: "test", Context: "{key:value}", DataType: "test", EnableShow: 1, Rss: ""}
	fmt.Println(reflect.TypeOf(hot), reflect.TypeOf(&hot))
	hot.AddHot()

}

func TestQueryHotbyId(t *testing.T) {
	hot := &Hot{}
	hot.QueryHotbyId(1)
	fmt.Println(hot)
}
