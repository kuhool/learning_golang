package basicLevel

import "C"
import (
	"fmt"
)

func ap(a []int) {
	fmt.Printf("%p\n", a)
	//b := append(a, 10)

	fmt.Printf("%p\n", a)
}

func app(a []int) {
	fmt.Printf("%p\n", a)
	a[0] = 1
	fmt.Printf("%p\n", a)
}

type ProductChangeInfo struct {
	IsChange int      `json:"is_change"`
	ImgUrls  []string `json:"img_urls"`
}
type Yfk struct {
	ProductImgs []string `json:"product_imgs"`
}
type OldYfk struct {
	ProductImgs []string `json:"product_imgs"`
}

type Product struct {
	ImgUrls []string `json:"Img_urls"`
}

func SliceTest4() {
	yfk := Yfk{
		ProductImgs: []string{"A", "B", "D", "F"},
	}
	oldYfk := OldYfk{
		ProductImgs: []string{"A", "B", "C", "F"},
	}
	product := Product{
		ImgUrls: []string{"A", "B", "F", "H"},
	}

	YkfNew := make([]string, 4)
	copy(YkfNew, yfk.ProductImgs)
	YfkOld := make([]string, 4)
	copy(YfkOld, oldYfk.ProductImgs)
	NewProduct := make([]string, 4)
	copy(NewProduct, product.ImgUrls)

	ProductChangeInfoTmp := ProductChangeInfo{}
	ProductImgChangeNew := make([]string, 0, 4)
	//ProductImgChangeNew1 := make([]string, 4)
	for i := 0; i < cap(ProductImgChangeNew); i++ {
		imgTmpNew := YkfNew[i]
		imgTmpOld := YfkOld[i]
		imgTmpProduct := product.ImgUrls[i]
		fmt.Println(i, imgTmpNew, imgTmpOld, imgTmpProduct)
		if len(imgTmpNew) > 0 && len(imgTmpOld) > 0 {
			if imgTmpNew != imgTmpOld {
				ProductChangeInfoTmp.IsChange = 1
				if len(imgTmpNew) > 0 {
					ProductImgChangeNew = append(ProductImgChangeNew, imgTmpNew)
				} else {
					ProductImgChangeNew = append(ProductImgChangeNew, imgTmpProduct)
				}
			} else {
				ProductImgChangeNew = append(ProductImgChangeNew, imgTmpProduct)
			}
		} else {
			if len(imgTmpNew) > 0 || len(imgTmpOld) > 0 {
				if len(imgTmpNew) > 0 {
					ProductImgChangeNew = append(ProductImgChangeNew, imgTmpNew)
				}

				if len(imgTmpOld) > 0 {
					ProductImgChangeNew = append(ProductImgChangeNew, imgTmpNew)
				}

			} else {
				ProductImgChangeNew = append(ProductImgChangeNew, imgTmpProduct)
			}

		}

	}
	fmt.Println(ProductImgChangeNew)

}
func SliceTest3() {
	//var x = nil
	//var x interface{} = nil
	//var x string = nil
	var x error = nil
	fmt.Println(x)
	//str := 'abc' + '123'
	//str := "abc" + "123"
	//str := '123' + "abc"
	//str := fmt.Sprintf("abc%d", 123)

	//fmt.Println(str)
}
func SliceTest1() {
	a := []int{7, 8, 9}
	fmt.Printf("%p\n", a)
	fmt.Printf("%+v\n", a)
	ap(a)
	fmt.Printf("%+v\n", a)
	app(a)
	fmt.Printf("%+v\n", a)
}
