package currencyconv

import "fmt"

type RMB float64
type HongKongDollar float64

func (r RMB) String() string            { return fmt.Sprintf("￥%g", r) }
func (h HongKongDollar) String() string { return fmt.Sprintf("$%g", h) }
