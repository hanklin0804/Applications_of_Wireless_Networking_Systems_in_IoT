import socket

localIP     = "140.118.122.155"
localPort   = 5057



# Create a datagram socket
UDPClientSocket = socket.socket(family=socket.AF_INET, type=socket.SOCK_DGRAM)

# Bind to address and ip
UDPClientSocket.bind((localIP, localPort))

print("UDP Client up and listening")


# Listen for incoming datagrams
n = 0
while(n<=10000):
    # Sending a msg to Proxy
    n += 1
    msg = 'hello '+ str(n)
    UDPClientSocket.sendto(msg.encode(), ("140.118.122.155", 5056))

    clientMsg = "Message from Sever:{}".format(UDPClientSocket.recvfrom(1024)[0])
    clientIP  = "Sever IP Address:{}".format(UDPClientSocket.recvfrom(1024)[1])
    print(clientMsg)
    print(clientIP)

