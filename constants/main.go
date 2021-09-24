package constants

const AUTH = "user:"

const APP_NAME = "go_app"

const SEND_EMAIL_Q = "send_email"

var STATUS = struct {
	ACTIVE                        string
	INACTIVE                      string
	DELETED                       string
	STATUS_CART_CHECKOUT          string
	STATUS_SALES_PAYMENT          string
	STATUS_SALES_PAYMENT_CONFIRM  string
	STATUS_SALES_SUPPLIER_CONFIRM string
	STATUS_SALES_SUPPLIER_PACKING string
	STATUS_SALES_SUPPLIER_SEND    string
	STATUS_SALES_CUSTOMER_RECEIVE string
	STATUS_SALES_CANCEL           string
}{
	INACTIVE:             "0",
	ACTIVE:               "1",
	DELETED:              "2",
	STATUS_CART_CHECKOUT: "3",

	STATUS_SALES_PAYMENT:          "1",
	STATUS_SALES_PAYMENT_CONFIRM:  "2",
	STATUS_SALES_SUPPLIER_CONFIRM: "3",
	STATUS_SALES_SUPPLIER_PACKING: "4",
	STATUS_SALES_SUPPLIER_SEND:    "5",
	STATUS_SALES_CUSTOMER_RECEIVE: "6",
	STATUS_SALES_CANCEL:           "9",
}
