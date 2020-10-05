import math

pi = math.pi #3.141592653589793 just change it if you want precious value such like 3.14
#pi = 3.14
#pi = 22/7

#luas lingkaran

#L = π × d²/4 = π × r²
def luas(r):
	r = r * r 
	luas = r * pi
	return luas


#keliling 
#K = π × d = 2 × π × r


def keliling(r):
	keliling = 2 * pi * r 
	return keliling
