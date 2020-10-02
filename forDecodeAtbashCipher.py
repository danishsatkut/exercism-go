'''
    Problem Task : Atbash decoder
    Problem Link : https://exercism.io/my/solutions/c5ecf77144d3428fb4f33909a2c8b7cb
'''

def decrypt(message):
	reverse_order = ["Z","Y","X","W","V","","U","T","S","R","Q","P","O","N","M","L","K","J","I","H","G","F","E","D","C","B","A"]
	alpha_order = ["A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","","V","W","X","Y","Z"]
	decoded_msg = ""

	for letter in message:
		if letter == " ":
			decoded_msg += " "
		else:
			for i in range(0,26):
				if letter == alpha_order[i]:
					decoded_msg += reverse_order[i]
	print(decoded_msg)

message = raw_input("enter message to decode in atbash cipher: ")
decrypt(message.upper())