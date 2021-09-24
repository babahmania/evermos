package repository

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"evermos/constants"
	dto "evermos/dto"
	models "evermos/models"

	"gorm.io/gorm"
)

type ICartRepository interface {
	ListCarts(pagi *dto.Pagination) (res []models.Cart, err error)
	GetBadgeCart(cart models.Cart) (countBadge int64, err error)
	GetDetailCart(cart models.CartDetailData) (res models.CartDetailData, err error)
	CreateCart(cart models.Cart, cartDetail models.CartDetail) (models.Cart, error)
	AddProductCart(cart models.Cart, cartDetail models.CartDetail) (models.Cart, error)
	UpdateProductCart(cart models.Cart, cartDetail models.CartDetail) (models.Cart, error)
	CheckoutCart(cart models.Sales) (models.Sales, error)
	//AlertCartStock
}

type CartRepository struct {
	Conn *gorm.DB
}

func (m *CartRepository) ListCarts(pagi *dto.Pagination) (res []models.Cart, err error) {
	var carts []models.Cart
	var count int64
	m.Conn.Limit(pagi.Limit).Offset(pagi.Offset).Find(&carts)
	m.Conn.Model(&models.Cart{}).Count(&count)
	pagi.Total = count

	return carts, nil
}

func (m *CartRepository) GetBadgeCart(cart models.Cart) (countBadge int64, err error) {
	var cBadge int64 = 0
	carts := models.Cart{}
	m.Conn.First(&carts, "user_id=? and is_active='1'", cart.UserID)
	if carts.CartID <= 0 {
		return cBadge, nil
	}
	cBadge = int64(carts.QtyItem)
	return cBadge, nil
}

func (m *CartRepository) GetDetailCart(cart models.CartDetailData) (res models.CartDetailData, err error) {
	var cartDetail models.CartDetailData
	carts := models.Cart{}
	m.Conn.First(&carts, "user_id=? and cart_id=? and is_active='1'", cart.UserID, cart.CartID)
	if cart.CartID <= 0 {
		return cartDetail, errors.New("cart is not exist")
	}
	cartDetails := models.CartDetail{}
	m.Conn.First(&cartDetails, "cart_id=? and is_active='1'", cart.CartID)
	if cart.CartID <= 0 {
		return cartDetail, errors.New("cart is not exist")
	}
	cartDetail.CartID = cart.CartID
	cartDetail.UserID = cart.UserID
	cartDetail.HeaderData = carts
	cartDetail.OrderData = cartDetails
	return cartDetail, nil
}

func (m *CartRepository) CreateCart(cart models.Cart, cartDetail models.CartDetail) (res models.Cart, err error) {
	//transaction mode still error
	//m.Conn.Begin()
	//m.Conn.Statement.Begin()
	//m.Conn.Statement.Rollback()
	//m.Conn.Statement.Commit()
	//return models.Cart{}, nil
	//check if cart is exist
	//b.First(&user, "id = ?", "1b74413f-f3b8-409f-ac47-e8c062e3472a")
	isExist := m.Conn.First(&cart, "user_id=? and is_active='1'", cart.UserID).RowsAffected
	if isExist > 0 {
		return models.Cart{}, errors.New("cart is exist")
	}
	//fmt.Println(isExist.result.RowsAffected)
	//return models.Cart{}, nil

	result := m.Conn.Create(&cart)
	if result.Error != nil {
		//m.Conn.RollbackTo("sp1")
		//m.Conn.Statement.Rollback()
		return models.Cart{}, result.Error
	} else {
		//m.Conn.SavePoint("sp1")
		cartDetail.CartID = cart.CartID
		cartDetail.CartNo = 1
		result = m.Conn.Create(&cartDetail)
		if result.Error != nil {
			//m.Conn.RollbackTo("sp1")
			//m.Conn.Statement.Rollback()
			return models.Cart{}, result.Error
		}
	}
	//m.Conn.Commit()
	return cart, nil
}

func (m *CartRepository) AddProductCart(cart models.Cart, cartDetail models.CartDetail) (res models.Cart, err error) {
	//transaction mode still error
	//m.Conn.Begin()
	//m.Conn.Statement.Begin()
	//m.Conn.Statement.Rollback()
	//m.Conn.Statement.Commit()
	//return models.Cart{}, nil
	//check if cart is exist
	oldCart := models.Cart{}
	m.Conn.First(&oldCart, "user_id=? and cart_id=? and is_active='1'", cart.UserID, cart.CartID)
	if oldCart.CartID <= 0 {
		return models.Cart{}, errors.New("cart is not exist")
	}

	var cartNo uint = 1
	var lastNo int64
	m.Conn.Table("cart_details").Where("cart_id=? and inv_id=? and is_active='1'", cart.CartID, cartDetail.InvID).Count(&lastNo)
	if lastNo > 0 {
		return models.Cart{}, errors.New("product item is already exist")
	}
	m.Conn.Table("cart_details").Where("cart_id=? and is_active='1'", cart.CartID).Count(&lastNo)
	//fmt.Println(lastNo)
	//return models.Cart{}, errors.New("lagi debug")
	cartNo = uint(lastNo) + 1
	cartDetail.CartID = cart.CartID
	cartDetail.CartNo = cartNo

	cart.QtyItem = cartDetail.QtyOrder + oldCart.QtyItem
	cart.AmountPrice = cart.AmountPrice + oldCart.AmountPrice
	cart.AmountDisc = cart.AmountDisc + oldCart.AmountDisc
	cart.AmountTotal = cart.AmountTotal + oldCart.AmountTotal
	cart.UpdatedAt = time.Now()
	result := m.Conn.Create(&cartDetail)
	if result.Error != nil {
		//m.Conn.RollbackTo("sp1")
		//m.Conn.Statement.Rollback()
		return models.Cart{}, result.Error
	} else {
		//m.Conn.SavePoint("sp1")
		result = m.Conn.Model(cart).Where("cart_id=? and user_id=?", cart.CartID, cart.UserID).Updates(&cart)
		if result.RowsAffected <= 0 {
			return models.Cart{}, result.Error
		}
	}
	//m.Conn.Commit()
	return cart, nil
}

func (m *CartRepository) UpdateProductCart(cart models.Cart, cartDetail models.CartDetail) (res models.Cart, err error) {
	//transaction mode still error
	//m.Conn.Begin()
	//m.Conn.Statement.Begin()
	//m.Conn.Statement.Rollback()
	//m.Conn.Statement.Commit()
	//return models.Cart{}, nil
	//check if cart is exist
	oldCart := models.Cart{}
	m.Conn.First(&oldCart, "user_id=? and cart_id=? and is_active='1'", cart.UserID, cart.CartID)
	if oldCart.CartID <= 0 {
		return models.Cart{}, errors.New("cart is not exist")
	}
	type totalOldAmount struct {
		TotalQtyOrder       int
		TotalAmountPrice    int
		TotalAmountDiscount int
		TotalAmount         int
	}
	var totalOldAmountData totalOldAmount
	m.Conn.Table("cart_details").Select("sum(qty_order) as TotalQtyOrder, sum(qty_order*amount_price) as TotalAmountPrice, sum(qty_order*amount_disc) as TotalAmountDiscount, sum(qty_order*amount_price)-sum(qty_order*amount_disc) as TotalAmount").Where("cart_id=? and inv_id<>? and is_active='1'", cart.CartID, cartDetail.InvID).Group("cart_id").Scan(&totalOldAmountData)
	if totalOldAmountData.TotalAmount <= 0 {
		return models.Cart{}, errors.New("cart detail this item product is not exist")
	}
	cart.QtyItem = uint(totalOldAmountData.TotalQtyOrder) + cartDetail.QtyOrder
	cart.AmountPrice = uint(totalOldAmountData.TotalAmountPrice) + +(cartDetail.QtyOrder * cartDetail.AmountPrice)
	cart.AmountDisc = uint(totalOldAmountData.TotalAmountDiscount) + (cartDetail.QtyOrder * cartDetail.AmountDisc)
	cart.AmountTotal = cart.AmountPrice - cart.AmountDisc
	/*
		cart.UpdatedAt = time.Now()
		var cartDetailNew = models.CartDetail{}
		cartDetailNew.CartID = cartDetail.CartID
		cartDetailNew.InvID = cartDetail.InvID
		cartDetailNew.Name = cartDetail.Name
		cartDetailNew.SupplierID = cartDetail.SupplierID
		cartDetailNew.CartNo = cartDetail.CartNo
		cartDetailNew.QtyOrder = cartDetail.QtyOrder
		cartDetailNew.AmountDisc = cartDetail.AmountDisc
		cartDetailNew.AmountPrice = cartDetail.AmountPrice
		cartDetailNew.CreatedAt = cartDetail.CreatedAt
	*/
	cartDetail.UpdatedAt = time.Now()
	//fmt.Println(cartDetailNew.AmountDisc)
	//return models.Cart{}, errors.New("lagi debug")
	result := m.Conn.Table("cart_details").Where("cart_id=? and inv_id=? and is_active='1'", cart.CartID, cartDetail.InvID).Updates(&cartDetail)
	if result.Error != nil {
		//m.Conn.RollbackTo("sp1")
		//m.Conn.Statement.Rollback()
		return models.Cart{}, result.Error
	} else {
		//m.Conn.SavePoint("sp1")
		result = m.Conn.Model(cart).Where("cart_id=? and user_id=?", cart.CartID, cart.UserID).Updates(&cart)
		if result.RowsAffected <= 0 {
			return models.Cart{}, result.Error
		}
	}
	//m.Conn.Commit()
	return cart, nil
}

func (m *CartRepository) CheckoutCart(cart models.Sales) (res models.Sales, err error) {
	//transaction mode still error
	//m.Conn.Begin()
	//m.Conn.Statement.Begin()
	//m.Conn.Statement.Rollback()
	//m.Conn.Statement.Commit()
	//return models.Cart{}, nil
	//check if cart is exist
	carts := models.Cart{}
	m.Conn.First(&carts, "user_id=? and cart_id=? and is_active='1'", cart.UserID, cart.SalesID)
	if carts.CartID <= 0 {
		return models.Sales{}, errors.New("cart is not exist")
	}
	cartDetails := models.CartDetail{}
	m.Conn.First(&cartDetails, "cart_id=? and is_active='1'", cart.SalesID)
	if cart.SalesID <= 0 {
		return models.Sales{}, errors.New("cart is not exist")
	}
	msgInventory := ""
	//check last stock product each item
	type Result struct {
		Name     string
		QtyStock int
		QtyOrder int
	}
	//var resultData Result
	var resultDatas []Result
	sql := "SELECT b.name, b.qty_stock as QtyStock, a.qty_order as QtyOrder " +
		"FROM cart_details a, products b " +
		"WHERE a.inv_id = b.inv_id AND a.supplier_id=b.supplier_id AND b.is_active='1' AND " +
		"a.is_active='1' and b.qty_stock<=a.qty_order AND a.cart_id = " + strconv.Itoa(int(cart.SalesID))
	m.Conn.Raw(sql).Scan(&resultDatas)

	//msgInventory detail item not enough
	//stock not enough
	if len(resultDatas) > 0 {
		msgInventory = strconv.Itoa(len(resultDatas)) + " item"
		return models.Sales{}, errors.New("Stock Last Inventory not enough, " + msgInventory)
	}

	//return models.Sales{}, errors.New("lg debug")

	//generate models
	sales := models.Sales{}
	sales.SalesDate = time.Now()
	sales.InvoiceNo = strconv.Itoa(int(cart.SalesID))
	sales.UserAddress = "Jl."
	sales.StatusOrder = "0"
	sales.UserID = carts.UserID
	sales.QtyItem = carts.QtyItem
	sales.AmountPrice = carts.AmountPrice
	sales.AmountDisc = carts.AmountDisc
	sales.AmountExpedition = cart.AmountExpedition
	sales.AmountTotal = carts.AmountTotal + sales.AmountExpedition
	//test save sales order
	result := m.Conn.Create(&sales)
	if result.Error != nil {
		//m.Conn.RollbackTo("sp1")
		//m.Conn.Statement.Rollback()
		return models.Sales{}, result.Error
	} else {
		//m.Conn.SavePoint("sp1")
		sql = "INSERT INTO sales_details( " +
			"sales_id, sales_no, inv_id, " +
			"	supplier_id, NAME, noted, " +
			"	qty_order, amount_price, amount_disc, amount_total, " +
			"	is_promo, is_active, created_at, updated_at, deleted_at " +
			")  " +
			"SELECT " + strconv.Itoa(int(cart.SalesID)) + " AS sales_id, cart_no  AS sales_no, inv_id, " +
			"	supplier_id, NAME, noted, " +
			"	qty_order, amount_price, amount_disc, amount_total, " +
			"	is_promo, is_active, created_at, updated_at, deleted_at " +
			"FROM cart_details " +
			"WHERE cart_id = " + strconv.Itoa(int(cart.SalesID)) + " AND is_active='1'"
		results := m.Conn.Exec(sql)
		if results.Error != nil {
			//m.Conn.RollbackTo("sp1")
			//m.Conn.Statement.Rollback()
			return models.Sales{}, errors.New("create sales order details error: " + fmt.Sprint(results.Error))
		}
		results = m.Conn.Table("carts").Where("cart_id=? and user_id=?", cart.SalesID, cart.UserID).Update("is_active", constants.STATUS.STATUS_CART_CHECKOUT)
		if results.RowsAffected <= 0 {
			return models.Sales{}, errors.New("checkout process error")
		}
	}
	//m.Conn.Commit()
	return sales, nil
}
