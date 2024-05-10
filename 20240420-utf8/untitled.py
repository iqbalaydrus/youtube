# convert binary ke arab/hijaiyah
# 0625 i
# 0642 q
# 0628 b
# 0627 a
# 0644 l

def unicode_to_binary(u: str) -> str:
	first_byte = "110"
	second_byte = "10"
	binary = bin(int(u, 16))[2:]
	first_byte += binary[:5]
	second_byte += binary[5:]
	return first_byte + second_byte

def binary_to_str(b: str) -> str:
	return bytes.fromhex(hex(int(b, 2))[2:]).decode()

print(binary_to_str(unicode_to_binary("0625")))
print(binary_to_str(unicode_to_binary("0642")))
print(binary_to_str(unicode_to_binary("0628")))
print(binary_to_str(unicode_to_binary("0627")))
print(binary_to_str(unicode_to_binary("0644")))