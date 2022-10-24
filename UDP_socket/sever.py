import socket

localIP     = "140.118.122.155"
localPort   = 5055
bufferSize  = 1024



# Create a datagram socket
UDPServerSocket = socket.socket(family=socket.AF_INET, type=socket.SOCK_DGRAM)

# Bind to address and ip
UDPServerSocket.bind((localIP, localPort))

print("UDP server up and listening")

# Listen for incoming datagrams
n= 0
while(True):
    bytesAddressPair = UDPServerSocket.recvfrom(bufferSize)
    client_message = bytesAddressPair[0][0]
    client_address = bytesAddressPair[0][1]
    clientMsg = "Message from Client:{}".format(client_message)
    clientIP  = "Client IP Address:{}".format(client_address)
    
    print(clientMsg)
    print(clientIP)

    # Sending a reply to client
    n+= 1
    msgFromServer       = "World " + str(n)
    bytesToSend         = str.encode(msgFromServer)
    UDPServerSocket.sendto(bytesToSend, ("140.118.122.155", 5057))