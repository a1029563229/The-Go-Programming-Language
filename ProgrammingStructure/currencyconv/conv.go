package currencyconv

func RToH(r RMB) HongKongDollar { return HongKongDollar(r * 1.1134) }
func HToR(h HongKongDollar) RMB { return RMB(h * 0.8981) }
