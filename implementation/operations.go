package implementation

import (
	"errors"
	"github.com/aditya/ProjectCatalog/models"
	"sort"
	"strconv"
	"time"
)

type Inmemoryimplement struct{
	ProductArr []models.Product

}

type TimeStampStructure struct{
	prodId int64
	timeInSec int64
}

var mp map[TimeStampStructure]int64

func Initializer(){
	mp=make(map[TimeStampStructure]int64)
	return
}

func min(a int64, b int64) int64{
	if a<=b{
		return a
	}

	return b

}

func (h *Inmemoryimplement) CreateProduct(product models.Product) error{

	for _, prod:=range h.ProductArr{
		if prod.Name==product.Name{
			return errors.New("Product already exist.")
		}
	}
	product.ID=strconv.Itoa(len(h.ProductArr)+1)
	h.ProductArr = append(h.ProductArr,product)

	return nil

}

func (h *Inmemoryimplement) ShowProduct() []models.Product{
	return h.ProductArr
}

func (h *Inmemoryimplement) ShowProductById(productId string) models.Product{
	for _,product:=range h.ProductArr{
		if product.ID==productId{
			return product
		}
	}
	return models.Product{}
}

func (h *Inmemoryimplement) UpdateProduct(updatedProduct models.Product,productName string) error{
	for i,item:=range h.ProductArr{
		if item.Name==productName{
			updatedProduct.ID = strconv.Itoa(i + 1)
			if item.Price != updatedProduct.Price {
				item.Price = updatedProduct.Price
			}
			if item.Description != updatedProduct.Description {
				item.Description = updatedProduct.Description
			}
			h.ProductArr[i] = updatedProduct
			return nil
		}
	}
	return errors.New("Product do not exist.")
}

func (h *Inmemoryimplement) BuyProduct(purchaseProduct models.Product, productName string) error{
	for i,item:=range h.ProductArr{
		if item.Name==productName{
			currCount, _ := strconv.ParseInt(item.Quantity, 10, 0)
			reqCount, _ := strconv.ParseInt(purchaseProduct.Quantity, 10, 0)
			if currCount-reqCount < 0 {
				return errors.New("Product not available.")
			}

			idint, _ :=strconv.ParseInt(item.ID,10,0)
			timestamp:=time.Now().Unix()

			local:=TimeStampStructure{
				prodId: idint,
				timeInSec: timestamp,
			}
			mp[local]=reqCount

			currCount-=reqCount

			h.ProductArr[i].Quantity=strconv.Itoa(int(currCount))
			return nil
		}
	}
	return errors.New("Item not available.")
}

func (h *Inmemoryimplement) TopProduct() []string{
	currTime:=time.Now().Unix()
	prodQuantity:=make([]int64,len(h.ProductArr)+1)

	for k,v :=range mp{
		if currTime-k.timeInSec<=3600{
			prodQuantity[k.prodId]+=v
		}
	}

	productsTop:=make([][2]int64,0)

	for ind,val:=range prodQuantity{
		if val>0 {
			arr := [2]int64{val, int64(ind)}
			productsTop = append(productsTop, arr)
		}
	}

	sort.Slice(productsTop, func(i, j int) bool {
		return productsTop[i][0] > productsTop[j][0]
	})

	listProducts:=make([]string,0)
	l:=len(productsTop)
	var i int64
	for i=0;i<min(5,int64(l));i++{
		ind:=productsTop[i][1]-1
		listProducts=append(listProducts,h.ProductArr[ind].Name)
	}

	return listProducts



}

