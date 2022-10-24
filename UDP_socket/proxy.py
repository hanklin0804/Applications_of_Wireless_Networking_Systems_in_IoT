import socket

localIP     = "140.118.122.155"
localPort   = 5056
bufferSize  = 1024


# Create a datagram socket
UDPProxySocket = socket.socket(family=socket.AF_INET, type=socket.SOCK_DGRAM)

# Bind to address and ip
UDPProxySocket.bind((localIP, localPort))

print("UDP Proxy up and listening")

# Listen for incoming datagrams
while(True):
    bytesAddressPair = UDPProxySocket.recvfrom(bufferSize)
    client_message = bytesAddressPair[0]
    client_address = bytesAddressPair[1]
    clientMsg = "Message from Client:{}".format(client_message)
    clientIP  = "Client IP Address:{}".format(client_address)
    print(clientMsg)
    print(clientIP)

    msgFromProxy       = client_message
    # Sending a reply to client
    UDPProxySocket.sendto(client_message, ("140.118.122.155", 5055))