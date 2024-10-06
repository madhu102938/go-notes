package main

type celsius float32
type fahrenheit float32
type kelvin float32

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

func (k kelvin) fahrenheit() fahrenheit {
	temp := k.celsius()
	return temp.fahrenheit()
}

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0)/5.0 + 32.0)
}

func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

func (f fahrenheit) celsius() celsius {
	return celsius(5.0 * (f - 32.0) / 9.0)
}

func (f fahrenheit) kelvin() kelvin {
	temp := f.celsius()
	return temp.kelvin()
}
func main() {
    
}
