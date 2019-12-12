package tempconv

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func CToK(c Celsius) Kelvins { return Kelvins(c - AbsoluteC) }

func KtoC(k Kelvins) Celsius { return Celsius(k + Kelvins(AbsoluteC)) }
