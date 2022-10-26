import socket

localIP     = "140.118.122.155"
localPort   = 5056



# Create a datagram socket
UDPProxySocket = socket.socket(family=socket.AF_INET, type=socket.SOCK_DGRAM)

# Bind to address and ip
UDPProxySocket.bind((localIP, localPort))

print("UDP Proxy up and listening")

# Listen for incoming datagrams
while(True):
    clientMsg , clientIP = UDPProxySocket.recvfrom(1024)

    print("Message from Client: ",clientMsg.decode())
    print("Client IP Address: ",clientIP)


    # Sending a msg to server
    UDPProxySocket.sendto(clientMsg, ("140.118.122.155", 5055))