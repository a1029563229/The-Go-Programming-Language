package tempconv

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func CToK(c Celsius) Kelvins { return Kelvins(c - AbsoluteC) }

func KtoC(k Kelvins) Celsius { return Celsius(k + Kelvins(AbsoluteC)) }

func IToM(i Inch) Meter { return Meter(i * 0.0254) }

func MToI(m Meter) Inch { return Inch(m / 0.0254) }

func IToK(i Ib) Kilogram { return Kilogram(i * 0.4535) }

func KToI(k Kilogram) Ib { return Ib(k / 0.4535) }
