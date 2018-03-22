def is_palindrome(number) :
	string = "{0}".format(number)
	
	for i in range(len(string) / 2) :
		if(string[i] != string[len(string) - 1 - i]) :
			return False
	return True 

def largest_palindrome(largest) :
	lp = 0
	for i in range(0, largest + 1) :
		for j in range(0, largest + 1) :
			if i*j > lp and is_palindrome(i * j):
				lp = i * j
	return lp


lp = largest_palindrome(999)
print("largest palindrome: {0}".format(lp))