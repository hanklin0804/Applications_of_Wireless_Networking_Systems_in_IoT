import socket

msgFromClient       = "Hello"
serverAddressPort   = ("140.118.122.155", 5056)
localIP     = "140.118.122.155"
localPort   = 5057
bufferSize  = 1024



# Create a UDP socket at client side
UDPClientSocket = socket.socket(family=socket.AF_INET, type=socket.SOCK_DGRAM)

# Bind to address and ip
UDPClientSocket.bind((localIP, localPort))

# Send to server using created UDP socket
UDPClientSocket.sendto(msgFromClient.encode('utf-8'), serverAddressPort)

msgFromServer = UDPClientSocket.recvfrom(bufferSize)

msg = "Message from Server {}".format(msgFromServer[0])
print(msg)