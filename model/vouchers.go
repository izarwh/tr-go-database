package model

type Voucher struct {
	id    int
	code  string
	value float32
}

func (v *Voucher) GetVoucher() (*int, *string, *float32) {
	return &v.id, &v.code, &v.value
}

func (v *Voucher) GetVoucherId() *int {
	return &v.id
}

func (v *Voucher) GetVoucherCode() *string {
	return &v.code
}

func (v *Voucher) GetVoucherValue() *float32 {
	return &v.value
}
